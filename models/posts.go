package models

import (
	"errors"
	"time"

	"github.com/coopernurse/gorp"
)

type Post struct {
	Id      int64
	Created int64
	Updated int64
	Title   string `form:"title"  binding:"required"`
	Body    string `form:"body"`
}

func (p *Post) PreInsert(s gorp.SqlExecutor) error {
	p.Created = time.Now().Unix()
	p.Updated = p.Created
	return nil
}

func CreateNewPost(post *Post) (err error) {
	return orp.Insert(post)
}

func (p *Post) GetCreatedTime() string {
	if p.Created == 0 {
		return ""
	}
	value := time.Unix(p.Created, 0)
	return value.String()
}

func GetAllPosts() (posts []Post, err error) {
	_, err = orp.Select(&posts, "select * from posts order by Id")
	return
}

func GetPost(id int64) (post *Post, err error) {
	var obj interface{}
	obj, err = orp.Get(Post{}, id)
	if err != nil || obj == nil {
		return nil, errors.New("Unable to find post")
	}

	post = obj.(*Post)
	return
}
