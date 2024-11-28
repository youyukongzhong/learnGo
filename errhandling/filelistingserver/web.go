package main

import (
	"learngo_muke/errhandling/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

// appHandler 它是一个专门用来处理请求的函数类型，如果发生错误，就返回一个 error。
type appHandler func(w http.ResponseWriter, req *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		// panic
		defer func() {
			if r := recover(); r != nil {
				log.Panicf("panic: %v", r)
				http.Error(w,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		err := handler(w, req) //执行传入的业务逻辑

		if err != nil { //如果有错误,处理错误
			log.Printf("Error occurred "+
				"handing request: %s",
				err.Error())

			// User error
			if userErr, ok := err.(userError); ok {
				http.Error(w, userErr.Message(), http.StatusBadRequest)
				return
			}

			// system error
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound //文件不存在 404
			case os.IsPermission(err):
				code = http.StatusForbidden // 没有权限 403
			default:
				code = http.StatusInternalServerError // 其它错误 500
			}
			http.Error(w, http.StatusText(code), code) // 返回合适的HTTP错误信息
		}
	}
}

type userError interface {
	error            // 给系统看的信息
	Message() string // 给用户看的信息
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.HandleFileList)) //注册路由

	err := http.ListenAndServe(":8888", nil) // 启动服务器
	if err != nil {
		panic(err) //启动失败时,终止程序
	}
}
