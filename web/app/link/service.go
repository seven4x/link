package link

import (
	"github.com/Seven4X/link/web/app/comment"
	"github.com/Seven4X/link/web/app/risk"
	"github.com/Seven4X/link/web/app/vote"
	"github.com/Seven4X/link/web/library/api"
	"github.com/Seven4X/link/web/library/api/messages"
	"github.com/Seven4X/link/web/library/log"
	"github.com/Seven4X/link/web/library/util"
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
func (s *Service) Save(link *Link) (id int, errs *api.Err) {

	if b := risk.IsAllowUrl(link.Link); !b {
		return -1, api.NewError(messages.LinkNotAllowDomain)
	}

	_, err := s.dao.Save(link)
	if err != nil {
		log.Error(err.Error())
		return -1, api.NewError(messages.GlobalErrorAboutDatabase)
	}
	//后添加 cuckoo-filter 第一次如果保存失败，下次还能成功保存
	success := s.filter.InsertUnique([]byte(strconv.Itoa(link.TopicId) + "_" + link.Link))
	if !success {
		return -1, api.NewError(messages.LinkRepeatInSameTopic)
	}

	comment := &comment.Comment{
		LinkId:   link.Id,
		Context:  risk.SafeUserText(link.FirstComment),
		CreateBy: link.CreateBy,
	}
	_, err = s.commentSvr.Save(comment)
	if err != nil {
		log.Error(err.Error())
	}

	return link.Id, nil
}

func (s *Service) SaveComment(req *NewCommentRequest) (id int, errs *api.Err) {
	comment := &comment.Comment{
		LinkId:   req.LinkId,
		Context:  risk.SafeUserText(req.Content),
		CreateBy: req.CreateBy,
	}
	if _, err := s.commentSvr.Save(comment); err != nil {
		return -1, api.NewError(messages.GlobalErrorAboutDatabase)
	}
	return comment.Id, nil
}

func (s *Service) ListLink(req *ListLinkRequest) (res []*ListLinkResponse, errs *api.Err) {

	return s.listLinkNoJoin(req)
}

//两种查询方法需要用基准测一下哪个快
func (s *Service) listLinkJoin(req *ListLinkRequest) (res []*ListLinkResponse, errs *api.Err) {
	req.Size = 10
	var links []Link
	var total int64
	var err error

	if req.UserId == 0 {
		links, total, err = s.dao.ListLinkJoinHotComment(req)
	} else {
		links, total, err = s.dao.ListLinkJoinCommentAndUserVote(req)
	}
	if err != nil {
		log.Error(err.Error())
	}
	res = make([]*ListLinkResponse, 0)
	visit(&links, func(m Link) {
		link := BuildLinkResponseOfModel(&m)
		res = append(res, link)
	})

	log.DebugW("ListLink",
		"links", len(links),
		"total", total)

	return res, nil
}

func visit(links *[]Link, f func(link Link)) {
	for _, link := range *links {
		f(link)
	}
}

/* 需要关联查询：
创建人 头像，昵称
热评，热评头像、昵称
是否喜欢
*/
func (s *Service) listLinkNoJoin(req *ListLinkRequest) (res []*ListLinkResponse, errs *api.Err) {
	req.Size = 10
	var links []Link
	var total int64
	var err error
	links, total, err = s.dao.ListLink(req)
	if err != nil {
		log.Error(err.Error())
	}
	log.DebugW("ListLink",
		"links", len(links),
		"total", total)
	res = make([]*ListLinkResponse, 0)
	visit(&links, func(m Link) {
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

	return res, nil
}
func (s *Service) fetchHotComment(ids []interface{}, links []*ListLinkResponse) {

	commentList, err := s.commentSvr.ListHotCommentByLinkId(ids)
	if err != nil {
		return
	}
	hash := make(map[int]comment.Comment)
	for _, comment := range commentList {
		hash[comment.LinkId] = comment
	}
	for _, link := range links {
		c, b := hash[link.Id]
		if b {
			link.HotComment = &HotComment{UserId: c.CreateBy, Context: c.Context, Avatar: c.Creator.Avatar}
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

func getLinkIds(links *[]Link) []interface{} {
	linkIds := make([]interface{}, 0)
	for _, link := range *links {
		linkIds = append(linkIds, link.Id)
	}
	return linkIds
}