package main

import (
	"fmt"
	"learngo_muke/interface/infra"
)

func getRetriever() retriever {
	return infra.Retriever{}
}

// ?: Something that can "Get"
type retriever interface {
	Get(string) string
}

func main() {
	var r retriever = getRetriever()
	fmt.Println(r.Get("http://www.baidu.com"))
}
