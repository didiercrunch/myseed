package main

import (
	"testing"
)

const yaml_params = `
---

port: 9999

host: 0.0.0.0  #localhost


`

func TestLoadFromYamlData(t *testing.T) {
	p := GetDefaultParams()
	if err := p.LoadFromYamlData([]byte(yaml_params)); err != nil {
		t.Error(err)
		return
	}
	if p.Port != 9999 {
		t.Error("port")
	}
	if p.Host != "0.0.0.0" {
		t.Error("host")
	}

}
