package main

import (
	"testing"
)

const yaml_params = `
---

port: 9999

host: 0.0.0.0  #localhost

# all the params relative to mongodb
mongo:
    # the url where to get mongodb
    url: localhost:1234
    
    # the name of the database to use
    database_name: geography
    
    # the name of the collection where all the street can be found
    street_collection: quebec_street_map
      
    
    # the collection where all the heat map point will be stored
    heat_map_collection: quebec_heat_map  # quebec_heat_map
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
	if p.Mongo.Url != "localhost:1234" {
		t.Error("mongo.url")
	}
	if p.Mongo.Url != "localhost:1234" {
		t.Error("mongo.url")
	}
	if p.Mongo.DatabaseName != "geography" {
		t.Error("mongo.DatabaseName")
	}

}
