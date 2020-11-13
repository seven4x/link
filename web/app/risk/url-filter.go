package risk

import (
	"bufio"
	"github.com/Seven4X/link/web/library/log"
	"github.com/Seven4X/link/web/library/util"
	valid "github.com/asaskevich/govalidator"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var urlOnce sync.Once
var trie *util.Trie

const (
	BlackUrlListRepo = "https://gitee.com/seven4q/domain_black_list/raw/master/data.txt"
)

func initFilter() {
	trie = util.NewTrie()

	c := http.Client{
		Timeout: 5 * time.Second,
	}
	rsp, err := c.Get(BlackUrlListRepo)
	if err != nil {
		log.Error("初始化URL黑名单错误")
		log.Error(err.Error())
		return
	}
	defer rsp.Body.Close()

	load(rsp.Body)
}
func load(rd io.Reader) {
	buf := bufio.NewReader(rd)
	i := 0
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				log.Error("Read URL黑名单错误")
			}
			break
		}
		b, host := IsUrl(string(line))
		if b && host != "" {
			trie.Insert(host)
			i++
		}
	}
	log.Infow("Add_BlackUrl", "size", i)
}

func IsAllowUrl(str string) bool {
	isurl, host := IsUrl(str)
	if !isurl {
		return false
	}
	urlOnce.Do(initFilter)
	has := trie.Search(host)
	return !has
}

func isUrl2(str string) (bool, string) {
	return valid.IsURL(str), ""
}

func isUrl(str string) (bool bool, host string) {
	strTemp := str
	if !strings.Contains(str, "://") {
		// support no indicated urlscheme but with colon for port number
		// http:// is appended so u.Parse will succeed, strTemp used so it does not impact rxURL.MatchString
		strTemp = "http://" + str
	}
	u, err := url.Parse(strTemp)
	if err != nil {
		log.Info(err.Error())
		return false, ""
	}

	address := net.ParseIP(u.Host)

	if address == nil {
		return strings.Contains(u.Host, "."), u.Host
	}

	return true, u.Host
}

func IsUrl(str string) (bool bool, host string) {
	return isUrl(str)
}
