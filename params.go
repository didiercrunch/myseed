package main

import (
	"io/ioutil"

	goyaml "gopkg.in/yaml.v2"
)

const PARAM_FILE string = "params.yml"

type Params struct {
	Port        int           `yaml:"port,omitempty"`
	Host        string        `yaml:"host,omitempty"`
}



func (this *Params) LoadFromYamlFile(fileName string) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	return this.LoadFromYamlData(data)
}

func (this *Params) LoadFromYamlData(data []byte) error {
	var err error
	p := new(Params)
	if err = goyaml.Unmarshal(data, p); err != nil {
		return err
	}

	if p.Host != "" {
		this.Host = p.Host
	}
	if p.Port > 0 {
		this.Port = p.Port
	}

	return nil
}



func GetDefaultParams() *Params {
	return &Params{7777, "localhost"}
}

var params = GetDefaultParams()

func init() {
	if err := params.LoadFromYamlFile(PARAM_FILE); err != nil {
		panic(err)
	}
}
