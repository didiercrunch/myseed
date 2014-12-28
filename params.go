package main

import (
	"io/ioutil"

	mgo "gopkg.in/mgo.v2"
	goyaml "gopkg.in/yaml.v2"
)

const PARAM_FILE string = "params.yml"

type Params struct {
	Port        int           `yaml:"port,omitempty"`
	Host        string        `yaml:"host,omitempty"`
	Mongo       *MongoParams  `yaml:"mongo,omitempty"`
	MongoClient *mgo.Session  `yaml:"-"`
	DB          *mgo.Database `yaml:"-"`
}

type MongoParams struct {
	Url          string `yaml:"url,omitempty"`
	DatabaseName string `yaml:"database_name,omitempty"`
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
	if p.Mongo == nil {
		return nil
	}
	if p.Mongo.Url != "" {
		this.Mongo.Url = p.Mongo.Url
	}
	if p.Mongo.DatabaseName != "" {
		this.Mongo.DatabaseName = p.Mongo.DatabaseName
	}
	if p.Mongo.DatabaseName != "" {
		this.Mongo.DatabaseName = p.Mongo.DatabaseName
	}

	return nil
}

func (this *Params) initMongo() error {
	var err error
	this.MongoClient, err = mgo.Dial(this.Mongo.Url)
	this.DB = this.MongoClient.DB(this.Mongo.DatabaseName)

	return err
}

func GetDefaultParams() *Params {
	mongoParams := &MongoParams{
		Url:          "localhost:27017",
		DatabaseName: "union",
	}
	return &Params{7777, "localhost", mongoParams, nil, nil}
}

var params = GetDefaultParams()

func init() {
	if err := params.LoadFromYamlFile(PARAM_FILE); err != nil {
		panic(err)
	}
	if err := params.initMongo(); err != nil {
		panic(err)
	}
}
