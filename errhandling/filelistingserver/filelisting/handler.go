package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (e userError) Error() string {
	return string(e)
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList(w http.ResponseWriter, req *http.Request) error {
	if strings.Index(req.URL.Path, prefix) != 0 {
		return userError("path must start" + "with" + prefix)
	}
	path := req.URL.Path[len("/list/"):] // 提取请求路径中的文件路径
	file, err := os.Open(path)           // 尝试打开文件
	if err != nil {
		return err // 返回打开文件时的错误
	}
	defer file.Close() // 确保函数退出时关闭文件

	all, err := ioutil.ReadAll(file) // 读取文件内容
	if err != nil {
		return err // 返回读取文件时的错误
	}

	w.Write(all) // 将文件内容写入响应
	return nil   // 成功时返回 nil
}
