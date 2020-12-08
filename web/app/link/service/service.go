package service

import (
	commentModel "github.com/Seven4X/link/web/app/comment/model"
	commentService "github.com/Seven4X/link/web/app/comment/service"
	"github.com/Seven4X/link/web/app/link/dao"
	"github.com/Seven4X/link/web/app/link/model"
	"github.com/Seven4X/link/web/app/link/server/request"
	"github.com/Seven4X/link/web/app/link/server/response"
	"github.com/Seven4X/link/web/app/risk"
	"github.com/Seven4X/link/web/library/api"
	"github.com/Seven4X/link/web/library/api/messages"
	"github.com/Seven4X/link/web/library/log"
	"github.com/Seven4X/link/web/library/util"
	cuckoo "github.com/seven4x/cuckoofilter"
	"strconv"
)

type Service struct {
	dao    *dao.Dao
	mSvr   *commentService.Service
	filter *cuckoo.ScalableCuckooFilter
}

func NewService() (s *Service) {
	s = &Service{
		dao:    dao.New(),
		mSvr:   commentService.NewService(),
		filter: util.GetCuckooFilter(), //优雅关闭时dump filter
	}
	return
}

/*
1.黑名单检查
2.当前主题重复检查，重复是不添加返回原ID
3.保存关联评论
*/
func (service *Service) Save(link *model.Link) (id int, errs *api.Err) {

	if b := risk.IsAllowUrl(link.Link); !b {
		return -1, api.NewError(messages.LinkNotAllowDomain)
	}

	_, err := service.dao.Save(link)
	if err != nil {
		log.Error(err.Error())
		return -1, api.NewError(messages.GlobalErrorAboutDatabase)
	}
	//后添加 cuckoo-filter 第一次如果保存失败，下次还能成功保存
	success := service.filter.InsertUnique([]byte(strconv.Itoa(link.TopicId) + "_" + link.Link))
	if !success {
		return -1, api.NewError(messages.LinkRepeatInSameTopic)
	}

	comment := &commentModel.Comment{
		LinkId:   link.Id,
		Context:  risk.SafeUserText(link.FirstComment),
		CreateBy: link.CreateBy,
	}
	_, err = service.mSvr.Save(comment)
	if err != nil {
		log.Error(err.Error())
	}

	return link.Id, nil
}

func (service *Service) SaveComment(req *request.NewCommentRequest) (id int, errs *api.Err) {
	comment := &commentModel.Comment{
		LinkId:   req.LinkId,
		Context:  risk.SafeUserText(req.Content),
		CreateBy: req.CreateBy,
	}
	if _, err := service.mSvr.Save(comment); err != nil {
		return -1, api.NewError(messages.GlobalErrorAboutDatabase)
	}
	return comment.Id, nil
}

/*
1 转换对象
2 关联查询
*/
func (service *Service) ListLink(req *request.ListLinkRequest) (res []response.ListLinkResponse, errs *api.Err) {
	req.Size = 10
	var links []model.Link
	var total int64
	var err error

	if req.UserId == 0 {
		links, total, err = service.dao.ListLinkWithoutUser(req)
	} else {
		links, total, err = service.dao.ListLinkWithUser(req)
	}
	if err != nil {
		log.Error(err.Error())
	}
	log.DebugW("ListLink",
		"links", len(links),
		"total", total)

	return nil, nil
}
