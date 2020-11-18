package util

import (
	"github.com/Seven4X/link/web/library/log"
	cuckoo "github.com/seven4x/cuckoofilter"
	"io/ioutil"
)

var (
	filter = cuckoo.NewScalableCuckooFilter()
)

const (
	fileName = "cuckoo-filter.data"
)

func GetCuckooFilter() *cuckoo.ScalableCuckooFilter {

	bytes, err := ioutil.ReadFile(fileName)

	if err == nil {
		decodeFilter, decodeError := cuckoo.DecodeScalableFilter(bytes)
		if decodeError == nil {
			filter = decodeFilter
		}
	}

	return filter
}

func DumpCuckooFilter() {
	log.Info("start DumpCuckooFilter")
	bytes := filter.Encode()
	err := ioutil.WriteFile(fileName, bytes, 0755)
	if err != nil {
		log.Error(err.Error())
	}

}
