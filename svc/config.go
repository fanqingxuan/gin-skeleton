package svc

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func MustLoad(filename string, v interface{}) {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		panic("read file " + filename + " error:" + err.Error())
	}
	err = yaml.Unmarshal(content, v)

	if err != nil {
		panic("parse config content error:" + err.Error())
	}
	return
}
