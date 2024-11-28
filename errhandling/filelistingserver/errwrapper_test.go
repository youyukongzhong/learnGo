package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(w http.ResponseWriter, req *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return string(e)
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(w http.ResponseWriter, req *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(w http.ResponseWriter, req *http.Request) error {
	return os.ErrNotExist
}

func errNotPermission(w http.ResponseWriter, req *http.Request) error {
	return os.ErrPermission
}

func errUnknown(w http.ResponseWriter, req *http.Request) error {
	return errors.New("unknown error")
}

func noError(w http.ResponseWriter, req *http.Request) error {
	fmt.Fprintln(w, "ok")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "not found"},
	{errNotPermission, 403, "not permission"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "not error"},
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet,
			"https://www.imooc.com", nil)
		f(response, request)

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(server.URL)

		verifyResponse(response, tt.code, tt.message, t)
	}
}

func verifyResponse(response *http.Response, expectedCode int, expectedMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(response.Body)
	body := strings.Trim(string(b), "\n")
	if response.StatusCode != expectedCode || body != expectedMsg {
		t.Errorf("expect (%d, %s); "+
			"got (%d, %s)",
			expectedCode, expectedMsg,
			response.StatusCode, body)
	}
}
