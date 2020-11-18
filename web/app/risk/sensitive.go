package risk

import (
	"github.com/Seven4X/link/web/library/log"
	"github.com/importcjj/sensitive"
	"sync"
)

var filter *sensitive.Filter
var once sync.Once

func InitFilter() {
	log.Infow("敏感词tire开始构建")
	filter = sensitive.New()
	filter.LoadNetWordDict("https://raw.githubusercontent.com/importcjj/sensitive/master/dict/dict.txt")
	filter.LoadNetWordDict("https://gitee.com/seven4q/sensitive-words/raw/master/words.txt")
	log.Infow("敏感词tire构建完成")
}

/*
  合法 返回 true
*/
func IsAllowText(context string) (bool, string) {
	once.Do(InitFilter)
	b, s := filter.FindIn(context)
	return !b, s
}

/*
和谐用户文本
*/
func SafeUserText(context string) string {
	once.Do(InitFilter)
	return filter.Replace(context, '*')
}
