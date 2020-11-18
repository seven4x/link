package service

import (
	commentModel "github.com/Seven4X/link/web/app/comment/model"
	commentService "github.com/Seven4X/link/web/app/comment/service"
	"github.com/Seven4X/link/web/app/link/dao"
	"github.com/Seven4X/link/web/app/link/model"
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
	repeat := service.filter.Lookup([]byte(strconv.Itoa(link.TopicId) + "_" + link.Link))
	if repeat {
		return -1, api.NewError(messages.LinkRepeatInSameTopic)
	}
	_, err := service.dao.Save(link)
	if err != nil {
		log.Error(err.Error())
		return -1, api.NewError(messages.GlobalErrorAboutDatabase)
	}
	comment := &commentModel.Comment{
		LinkId:  link.Id,
		Context: risk.SafeUserText(link.FirstComment),
	}
	_, err = service.mSvr.Save(comment)
	if err != nil {
		log.Error(err.Error())
	}

	return link.Id, nil
}
