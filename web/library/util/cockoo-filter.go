package util

import (
	"github.com/Seven4X/link/web/library/log"
	"github.com/seiflotfy/cuckoofilter"
	"io/ioutil"
)

var (
	//todo resize
	filter = cuckoo.NewFilter(1000)
)

const (
	fileName = "cuckoo-filter.data"
)

func GetFilter() *cuckoo.Filter {

	bytes, err := ioutil.ReadFile(fileName)

	if err == nil {
		decodeFilter, decodeError := cuckoo.Decode(bytes)
		if decodeError == nil {
			filter = decodeFilter
		}
	}

	return filter
}

func DumpFilter() {

	bytes := filter.Encode()
	err := ioutil.WriteFile(fileName, bytes, 0755)
	if err != nil {
		log.Error(err.Error())
	}

}
