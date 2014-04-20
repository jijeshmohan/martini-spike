package models

import (
	"net/http"

	"github.com/martini-contrib/binding"
)

type Post struct {
	Id int `json:"id"`

	Title string `json:"title"`

	Body string `json:"body"`
}

func (p Post) Validate(errors *binding.Errors, req *http.Request) {

	if len(p.Title) < 4 {
		errors.Fields["title"] = "Too short; minimum 5 characters"
	} else if len(p.Title) > 120 {
		errors.Fields["title"] = "Too long; maximum 120 characters"
	}
	if len(p.Body) < 5 {
		errors.Fields["body"] = "Too short; minimum 5 characters"
	}
}
