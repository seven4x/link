package service

import (
	"errors"
	"github.com/seven4x/link/api"
	"github.com/seven4x/link/app"
	"github.com/seven4x/link/db"
	"github.com/seven4x/link/service/risk"
	"sync"
)

/*SaveLink
1.黑名单检查
2.当前主题重复检查，重复是不添加返回原ID
3.保存关联评论
*/
func (s *Service) SaveLink(link *db.Link) (id int, errs *app.Err) {

	if b := risk.IsAllowUrl(link.Link); !b {
		return -1, app.NewError(api.LinkNotAllowDomain)
	}

	if link.FirstComment == "" {
		link.CommentCnt = 0
	}
	_, err := s.Dao.SaveLink(link)
	if err != nil {
		app.Error(err.Error())
		return -1, app.NewError(api.GlobalErrorAboutDatabase)
	}

	if link.FirstComment == "" {
		return link.Id, nil
	}

	cmt := &db.Comment{
		LinkId:   link.Id,
		Content:  risk.SafeUserText(link.FirstComment),
		CreateBy: link.CreateBy,
	}
	_, err = s.SaveComment(cmt)
	if err != nil {
		app.Error(err.Error())
	}

	return link.Id, nil
}

/*ListLinkNoJoin
需要关联查询：
创建人 头像，昵称
热评，热评头像、昵称
是否喜欢
*/
func (s *Service) ListLinkNoJoin(req *api.ListLinkRequest) (res []*api.ListLinkResponse, errs error) {
	req.Size = 10
	var links []db.WithUser
	var err error
	links, err = s.Dao.ListLink(req)
	if err != nil {
		app.Error(err.Error())
		return nil, errors.New("db error")
	}

	res = make([]*api.ListLinkResponse, 0)
	visit(&links, func(m db.WithUser) {
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

func (s *Service) fetchHotComment(ids []interface{}, links []*api.ListLinkResponse) {

	commentList, err := s.ListHotCommentByLinkId(ids)
	if err != nil {
		return
	}
	hash := make(map[int]db.CommentUser)
	for _, cmt := range commentList {
		hash[cmt.LinkId] = cmt
	}
	for _, link := range links {
		c, b := hash[link.Id]
		if b {
			link.HotComment = &api.HotComment{UserId: c.CreateBy, Content: c.Content, Avatar: c.Creator.Avatar}
		}
	}
}

func (s *Service) fetchIsLike(ids []interface{}, links []*api.ListLinkResponse, req *api.ListLinkRequest) {
	userVotes, e := s.ListIsLike(ids, req.UserId, "like")
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

func (s *Service) GetRecentLinks(prev int) (links []db.Link, err error) {
	return s.Dao.GetRecentLinks(prev)
}

func visit(links *[]db.WithUser, f func(link db.WithUser)) {
	for _, link := range *links {
		f(link)
	}
}

func getLinkIds(links *[]db.WithUser) []interface{} {
	linkIds := make([]interface{}, 0)
	for _, link := range *links {
		linkIds = append(linkIds, link.Id)
	}
	return linkIds
}
