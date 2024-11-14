package infra

import (
	"io/ioutil"
	"net/http"
)

// 假设有一个做基础架构的团队
type Retriever struct {
}

func (Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	return string(bytes)
}
