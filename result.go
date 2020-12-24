package main

import (
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

type Response struct {
	Envir    string `yaml:"envir"`
	HttpCode int    `yaml:"httpCode"`
	Body     string `yaml:"body"`
}

func NewResponse(envir *Envir, response *http.Response, body []byte) Response {
	return Response{
		Envir:    envir.ToString(),
		HttpCode: response.StatusCode,
		Body:     string(body),
	}
}

type ResultList struct {
	List []Result `yaml:"list"`
}

type Result struct {
	Api       string     `yaml:"api"`
	IsSame    bool       `yaml:"isSame"`
	EnvirList []Response `yaml:"envirList"`
}

func NewResult(api *Api) Result {
	return Result{
		Api:    api.ToString(),
		IsSame: true,
	}
}

func (r *Result) Append(response Response) {
	r.EnvirList = append(r.EnvirList, response)
}

func Out2File(resultList *ResultList, outFileName string) {
	file, err := os.Create(outFileName)
	if err != nil {
		ErrorAndExit(err.Error())
	}

	defer file.Close()

	err = yaml.NewEncoder(file).Encode(resultList)
	if err != nil {
		ErrorAndExit(err.Error())
	}
}
