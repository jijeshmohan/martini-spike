package routes

import (
	"blogs/models"
	"blogs/modules/middleware"
	"log"
	"strconv"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func ListBlogs(r render.Render, data middleware.Data) {
	posts, err := models.GetAllPosts()
	if err != nil {
		log.Println(err)
	}
	data["posts"] = posts
	r.HTML(200, "blogs/index", data)
}

func NewBlog(r render.Render, data middleware.Data) {
	r.HTML(200, "blogs/new", data)
}

func CreateBlog(r render.Render, err binding.Errors, post models.Post, data middleware.Data) {
	if err.Count() > 0 {
		data["errors"] = err
		r.HTML(200, "blogs/new", data)
		return
	}
	if e := models.CreateNewPost(&post); e != nil {
		log.Println(err)
	}
	r.Redirect("/")
}

func ShowBlog(params martini.Params, r render.Render, c martini.Context) {
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		c.Next()
		return
	}

	post, e := models.GetPost(id)
	if e != nil {
		HandleError(404, r)
		return
	}
	r.HTML(200, "blogs/show", post)
}
