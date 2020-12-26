package main

import (
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

type Conf struct {
	ApiList   []Api   `yaml:"apiList"`
	EnvirList []Envir `yaml:"envirList"`
}

type Api struct {
	Name    string            `yaml:"name"`
	Method  string            `yaml:"method"`
	Path    string            `yaml:"path"`
	Params  string            `yaml:"params"`
	Timeout int64             `yaml:"timeout"`
	Headers map[string]string `yaml:"headers"`
}

func (a *Api) ToString() string {
	return fmt.Sprintf("%s [%s] %s params:%s", a.Name, strings.ToUpper(a.Method), a.Path, a.Params)
}

type Envir struct {
	Name    string `yaml:"name"`
	BaseUri string `yaml:"baseUri"`
	Host    string `yaml:"host"`
}

func (e *Envir) ToString() string {
	return fmt.Sprintf("%s baseUrl:%s host:%s", e.Name, e.BaseUri, e.Host)
}

func LoadConf(fileName string) *Conf {
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		ErrorAndExit(err.Error())
	}

	conf := &Conf{}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		ErrorAndExit(err.Error())
	}

	return conf
}
