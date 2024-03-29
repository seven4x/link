package app

import (
	cuckoo "github.com/seven4x/cuckoofilter"
	"github.com/seven4x/link/app/log"
	"io/ioutil"
)

var (
	filter = cuckoo.NewScalableCuckooFilter()
)

const (
	fileName = "cuckoo-filter.data"
)

func GetCuckooFilter() *cuckoo.ScalableCuckooFilter {

	bytes, err := ioutil.ReadFile(getFilePath())

	if err == nil {
		decodeFilter, decodeError := cuckoo.DecodeScalableFilter(bytes)
		if decodeError == nil {
			filter = decodeFilter
		}
	}

	return filter
}

func getFilePath() string {
	if res := GetConfigString("data_path"); res != "" {
		return res + "/" + fileName
	} else {
		return fileName
	}
}

//回档，从已有数据中加载重新构建filter todo
func Correction() {

}

func DumpCuckooFilter() {
	log.Info("start DumpCuckooFilter")
	bytes := filter.Encode()
	err := ioutil.WriteFile(getFilePath(), bytes, 0755)
	if err != nil {
		log.Error(err.Error())
	}

}
