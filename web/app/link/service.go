package link

import (
	"github.com/Seven4X/link/web/app/comment"
	"github.com/Seven4X/link/web/app/messages"
	"github.com/Seven4X/link/web/app/risk"
	"github.com/Seven4X/link/web/app/util"
	"github.com/Seven4X/link/web/app/vote"
	cuckoo "github.com/seven4x/cuckoofilter"
	"strconv"
	"sync"
)

type Service struct {
	dao        *Dao
	commentSvr *comment.Service
	voteSvr    *vote.Service
	filter     *cuckoo.ScalableCuckooFilter
}

func NewService() (s *Service) {
	s = &Service{
		dao:        NewDao(),
		commentSvr: comment.NewService(),
		voteSvr:    vote.NewService(),
		filter:     util.GetCuckooFilter(), //优雅关闭时dump filter
	}
	return
}

/*
1.黑名单检查
2.当前主题重复检查，重复是不添加返回原ID
3.保存关联评论
*/
func (s *Service) Save(link *Link) (id int, errs *util.Err) {

	if b := risk.IsAllowUrl(link.Link); !b {
		return -1, util.NewError(messages.LinkNotAllowDomain)
	}

	bytes := []byte(strconv.Itoa(link.TopicId) + "_" + link.Link)
	success := s.filter.InsertUnique(bytes)
	if !success {
		return -1, util.NewError(messages.LinkRepeatInSameTopic)
	}
	if link.FirstComment == "" {
		link.CommentCnt = 0
	}
	_, err := s.dao.Save(link)
	if err != nil {
		util.Error(err.Error())
		s.filter.Delete(bytes)
		return -1, util.NewError(messages.GlobalErrorAboutDatabase)
	}

	if link.FirstComment == "" {
		return link.Id, nil
	}

	cmt := &comment.Comment{
		LinkId:   link.Id,
		Content:  risk.SafeUserText(link.FirstComment),
		CreateBy: link.CreateBy,
	}
	_, err = s.commentSvr.Save(cmt)
	if err != nil {
		util.Error(err.Error())
	}

	return link.Id, nil
}

func (s *Service) ListLink(req *ListLinkRequest) (res []*ListLinkResponse, total int, errs *util.Err) {

	res, t, errs := s.listLinkNoJoin(req)
	return res, int(t), errs
}

//两种查询方法需要用基准测一下哪个快
func (s *Service) listLinkJoin(req *ListLinkRequest) (res []*ListLinkResponse, total int64, errs *util.Err) {
	req.Size = 10
	var links []LinkUser
	var err error

	if req.UserId == 0 {
		links, total, err = s.dao.ListLink(req)
	} else {
		links, total, err = s.dao.ListLinkJoinUserVote(req)
	}
	if err != nil {
		util.Error(err.Error())
	}
	res = make([]*ListLinkResponse, 0)
	visit(&links, func(m LinkUser) {
		link := BuildLinkResponseOfModel(&m)
		res = append(res, link)
	})
	ids := getLinkIds(&links)
	//var wg sync.WaitGroup
	wg := sync.WaitGroup{}
	//热评
	wg.Add(1)
	go func() {
		s.fetchHotComment(ids, res)
		wg.Done()
	}()

	wg.Wait()
	return res, total, nil
}

/* 需要关联查询：
创建人 头像，昵称
热评，热评头像、昵称
是否喜欢
*/
func (s *Service) listLinkNoJoin(req *ListLinkRequest) (res []*ListLinkResponse, total int64, errs *util.Err) {
	req.Size = 10
	var links []LinkUser
	var err error
	links, total, err = s.dao.ListLink(req)
	if err != nil {
		util.Error(err.Error())
	}

	res = make([]*ListLinkResponse, 0)
	visit(&links, func(m LinkUser) {
		link := BuildLinkResponseOfModel(&m)
		res = append(res, link)
	})

	ids := getLinkIds(&links)
	//var wg sync.WaitGroup
	wg := sync.WaitGroup{}
	//热评
	wg.Add(1)
	go func() {
		s.fetchHotComment(ids, res)
		wg.Done()
	}()

	//是否喜欢
	if req.UserId != 0 {
		wg.Add(1)
		go func() {
			s.fetchIsLike(ids, res, req)
			wg.Done()
		}()
	}
	wg.Wait()

	return res, total, nil
}

func (s *Service) fetchHotComment(ids []interface{}, links []*ListLinkResponse) {

	commentList, err := s.commentSvr.ListHotCommentByLinkId(ids)
	if err != nil {
		return
	}
	hash := make(map[int]comment.CommentUser)
	for _, cmt := range commentList {
		hash[cmt.LinkId] = cmt
	}
	for _, link := range links {
		c, b := hash[link.Id]
		if b {
			link.HotComment = &HotComment{UserId: c.CreateBy, Content: c.Content, Avatar: c.Creator.Avatar}
		}
	}
}

func (s *Service) fetchIsLike(ids []interface{}, links []*ListLinkResponse, req *ListLinkRequest) {
	userVotes, e := s.voteSvr.ListIsLike(ids, req.UserId, vote.VoteType_Link)
	if e == nil {
		for _, userVote := range userVotes {
			for _, link := range links {
				if userVote.Id == link.Id {
					link.IsLike = userVote.IsLike
					break
				}
			}
		}

	}
}

func visit(links *[]LinkUser, f func(link LinkUser)) {
	for _, link := range *links {
		f(link)
	}
}

func getLinkIds(links *[]LinkUser) []interface{} {
	linkIds := make([]interface{}, 0)
	for _, link := range *links {
		linkIds = append(linkIds, link.Id)
	}
	return linkIds
}
