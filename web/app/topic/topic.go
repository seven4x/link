package topic

import "github.com/Seven4X/link-link/web/app/risk"

/*
1.敏感词校验
2.重复校验
3.
*/
func NewTopic(t *Topic) (bool, string) {
	var b, s = risk.IsAllowText(t.Name)
	//todo
	return b, s
}
