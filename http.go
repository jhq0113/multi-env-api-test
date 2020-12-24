package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	DefaultMethod  = "GET"
	DefaultTimeout = 30
)

type Request struct {
	Method     string
	Url        string
	Params     interface{}
	Timeout    int64
	Headers    map[string]string
	HasGetBody bool
}

func NewRequest(api *Api, envir *Envir) *Request {
	req := &Request{
		Method:  DefaultMethod,
		Timeout: DefaultTimeout,
		Params:  api.Params,
		Url:     envir.BaseUri + api.Path,
		Headers: api.Headers,
	}

	if envir.Host != "" {
		if req.Headers == nil {
			req.Headers = make(map[string]string, 1)
		}
		req.Headers["Host"] = envir.Host
	}

	return req
}

func (r *Request) Send() (response *http.Response, body []byte) {
	r.Method = strings.ToUpper(r.Method)
	var reqBody io.Reader
	if r.Params != nil {
		args := bytes.NewBufferString("")
		switch r.Params.(type) {
		case map[string]string:
			for key, value := range r.Params.(map[string]string) {
				args.WriteString(key)
				args.WriteString("=")
				args.WriteString(value)
				args.WriteString("&")
			}
		case map[string]interface{}:
			for key, value := range r.Params.(map[string]interface{}) {
				args.WriteString(key)
				args.WriteString("=")
				args.WriteString(fmt.Sprint(value))
				args.WriteString("&")
			}
		case string:
			args.WriteString(r.Params.(string))
		case []byte:
			args.Write(r.Params.([]byte))
		default:
			ErrorAndExit("不支持的参数类型")
		}

		if r.Method == "GET" && !r.HasGetBody {
			paramPrefix := "?"
			if strings.Index(r.Url, "?") > 0 {
				paramPrefix = "&"
			}
			r.Url += paramPrefix + args.String()
		} else {
			reqBody = args
		}
	}

	req, err := http.NewRequest(r.Method, r.Url, reqBody)
	if err != nil {
		Error(err.Error())
		os.Exit(0)
	}

	if r.Headers != nil && len(r.Headers) > 0 {
		for key, value := range r.Headers {
			//设置请求Host
			if key == "Host" {
				req.Host = value
				continue
			}
			req.Header.Set(key, value)
		}
	}

	client := http.DefaultClient
	client.Timeout = time.Second * time.Duration(r.Timeout)
	response, err = client.Do(req)
	if err != nil {
		Error(err.Error())
		return nil, nil
	}

	defer response.Body.Close()

	body, _ = ioutil.ReadAll(response.Body)
	return response, body
}
