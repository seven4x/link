package link

import (
	"github.com/cockroachdb/errors"
	cuckoo "github.com/seven4x/cuckoofilter"
	"github.com/seven4x/link/web/comment"
	"github.com/seven4x/link/web/messages"
	"github.com/seven4x/link/web/risk"
	"github.com/seven4x/link/web/util"
	"github.com/seven4x/link/web/vote"
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

/*Save
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

/*ListLinkNoJoin
需要关联查询：
创建人 头像，昵称
热评，热评头像、昵称
是否喜欢
*/
func (s *Service) ListLinkNoJoin(req *ListLinkRequest) (res []*ListLinkResponse, errs error) {
	req.Size = 10
	var links []WithUser
	var err error
	links, err = s.dao.ListLink(req)
	if err != nil {
		util.Error(err.Error())
		return nil, errors.Wrap(err, "db error")
	}

	res = make([]*ListLinkResponse, 0)
	visit(&links, func(m WithUser) {
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

func (s *Service) GetRecentLinks(prev int) (links []Link, err error) {
	return s.dao.GetRecentLinks(prev)
}

func visit(links *[]WithUser, f func(link WithUser)) {
	for _, link := range *links {
		f(link)
	}
}

func getLinkIds(links *[]WithUser) []interface{} {
	linkIds := make([]interface{}, 0)
	for _, link := range *links {
		linkIds = append(linkIds, link.Id)
	}
	return linkIds
}
