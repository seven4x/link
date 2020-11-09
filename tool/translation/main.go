package main

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alimt"
	"github.com/iancoleman/orderedmap"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"sync"
)

var (
	client *alimt.Client
	from   string
	to     []string
	file   string
	output string
)

type (
	Req struct {
		from   string
		to     string
		key    string
		text   string
		result string
	}
)

func InitClient(ak, sk string) {
	c, err := alimt.NewClientWithAccessKey("cn-hangzhou", ak, sk)
	if err != nil {
		fmt.Print(err.Error())
	}
	client = c
}

/*
en英，fr法，de德，it意大利，ja日语,pt葡萄牙，ru俄语,es西班牙，ko韩语
idea -s zh -t en,fr,de,it,ja,pt,ru,ko -f zh-message.json
trans -s zh -t en,fr,de,it,ja,pt,ru,ko -f zh-cn-message.json  ak sk
json 解析成 orderedMap 保持原文案顺序
api请求并发数控制
*/
func main() {
	Execute()
}

var rootCmd = &cobra.Command{
	Use:   "trans",
	Short: "trans",
	Long:  `trans`,
	Run: func(cmd *cobra.Command, args []string) {
		InitClient(args[0], args[1])
		omap := getMessage(file)
		for _, t := range to {
			trans(from, t, omap, output)
		}
	},
}

func getMessage(file string) *orderedmap.OrderedMap {
	jsonFile, err := os.Open(file)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Successfully Opened %s \n", file)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	//var result map[string]string
	o := orderedmap.New()

	me := json.Unmarshal(byteValue, &o)
	if me != nil {
		println(me.Error())
	}
	return o
}

func trans(from, to string, o *orderedmap.OrderedMap, output string) error {
	wg := sync.WaitGroup{}
	var limit = make(chan int, 50)
	var lock sync.Mutex // map not safe
	resultMap := orderedmap.New()
	defer close(limit)

	for _, k := range o.Keys() {
		v, _ := o.Get(k)
		wg.Add(1)
		limit <- 1
		key := k
		go func() {
			req := &Req{from: from, to: to, key: key, text: v.(string)}

			fmt.Printf("%s \t %s \n", req.key, req.text)
			aliApi(req)
			fmt.Printf("result:\t %s \n\n", req.result)

			lock.Lock()
			resultMap.Set(req.key, req.result)
			lock.Unlock()

			wg.Done()
			<-limit //api请求完成再从ch取出才能实现并发控制效果
		}()
	}

	wg.Wait()

	fmt.Printf("%s done!\n", to)

	//因并发调度resultMap非有序的结果
	orderedMap := orderedmap.New()
	for _, k := range o.Keys() {
		v, _ := resultMap.Get(k)
		orderedMap.Set(k, v.(string))
	}
	bytes, err := json.Marshal(orderedMap)
	if err != nil {
		println(err.Error())
	}
	err = ioutil.WriteFile(output+"/"+to+"-message.json", bytes, 0777)
	if err != nil {
		println(err.Error())
		return err
	}
	return nil
}

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&from, "source", "s", "zh", "from")
	//https://github.com/spf13/cobra/issues/661
	rootCmd.PersistentFlags().StringSliceVarP(&to, "target", "t", []string{"en", "fr"}, "to")
	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "zh-message.json", "file")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", ".", "output")
	rootCmd.MarkFlagRequired("file")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func aliApi(req *Req) {

	request := alimt.CreateTranslateGeneralRequest()
	request.Scheme = "https"

	request.FormatType = "text"
	request.SourceLanguage = req.from
	request.TargetLanguage = req.to
	request.SourceText = req.text

	response, err := client.TranslateGeneral(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	if response == nil {
		req.result = req.text
	} else {
		req.result = response.Data.Translated
	}

}
