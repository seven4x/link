package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/stacks/arraystack"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

var (
	fileName string
)

var rootCmd = &cobra.Command{
	Use:   "topic",
	Short: "topic",
	Long:  `topic`,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open(fileName)
		if err != nil {
			println(err.Error())
		}
		defer file.Close()
		importTopic(file)
	},
}

type Topic struct {
	name     string
	code     string
	id       int
	refId    int
	position int
	subs     *arraylist.List
}

var (
	root = &Topic{name: "", id: 0, subs: arraylist.New()}
)

func importTopic(file *os.File) {
	scanner := bufio.NewScanner(file)
	stack := arraystack.New() // empty

	for scanner.Scan() {
		code, name := parse(scanner.Text())
		newTopic := Topic{name: name, code: code, subs: arraylist.New()}
		parent := getParent(stack, newTopic)
		parent.subs.Add(&newTopic)
		stack.Push(&newTopic)
	}
	log.Print(json.Marshal(root))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getParent(stack *arraystack.Stack, t Topic) (topic *Topic) {
	top, has := stack.Peek()
	if !has {
		return root
	}
	topTopic := top.(*Topic)
	if strings.Contains(t.code, topTopic.code) {
		return topTopic
	} else {
		stack.Pop()
		if stack.Size() == 0 {
			return root
		}
		return getParent(stack, t)
	}
}

func parse(text string) (code, name string) {
	res := strings.Split(text, ":")
	code = res[0]
	name = res[1]
	return code, name
}

//topic -f topic-data.txt
func main() {
	Execute()
}
func Execute() {
	rootCmd.PersistentFlags().StringVarP(&fileName, "fileName", "f", "topic-data.txt", "fileName")
	rootCmd.MarkFlagRequired("fileName")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
