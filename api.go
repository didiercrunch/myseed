package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

/*
Api structure

/weblabs
	POST:  create a new weblab.  The payload must be
	       {
				"name": "web lab frendly name",
				"cases": ["case_name_1", "case_name_2", ..., "case_name_n"],
				"probabilities": [prob_1, prob_2, ..., prob_n],                 // the sum must be 1
				"private": false,
			}
	GET:
			returns the list of all the created weblabs

/weblabs/<id>
   GET:  returns the weblab as

/weblabs/<id>/history
	GET: get the history
*/

// the smallest value that makes two floats the same.
const epsilon float64 = 0.000000001

type Cases struct {
	Name        string  `json:"name"`
	Probability float64 `json:"probability"`
}

type WeblabDefinition struct {
	Id      bson.ObjectId `json:"id"`
	Name    string        `json:"name"`
	Cases   []*Cases      `json:"cases"`
	Private bool          `json:"private"`
}

func (w *WeblabDefinition) verifyCasesStructure() error {
	names := make(map[string]bool)
	for _, c := range w.Cases {
		if c.Probability+epsilon < 0 {
			return errors.New("probability cannot be less 0")
		}
		if c.Probability-epsilon > 1 {
			return errors.New("probability cannot be greater 0")
		}
		c.Name = strings.Trim(c.Name, " \t\n\r")
		if c.Name == "" {
			return errors.New("weblab cannot have empty names")
		}
		names[c.Name] = true
	}
	if len(names) != len(w.Cases) {
		return errors.New("multiple weblab cannot have the same name")
	}
	return nil
}

func (w *WeblabDefinition) VerifyProbabilitySum() {

}
func (w *WeblabDefinition) Verify() error {

	return nil
}

func CreateNewWeblab(w http.ResponseWriter, request *http.Request) {
	d := json.NewDecoder(request.Body)
	wab := new(WeblabDefinition)
	if err := d.Decode(d); err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, err.Error())
		return
	}
	wab.Id = bson.NewObjectId()

}
