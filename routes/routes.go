package routes

import (
	"blogs/models"
	"fmt"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

func InitRoutes(m *martini.ClassicMartini) {
	m.Get("/", ListBlogs)
	m.Get("/new", NewBlog)
	m.Post("/new", binding.Form(models.Post{}), CreateBlog)
	m.Get("/:id", ShowBlog)

	m.NotFound(func(r render.Render) {
		r.HTML(404, "staus/404", "")
	})
}

func HandleError(err int, r render.Render) {
	template := fmt.Sprintf("status/%d", err)
	r.HTML(err, template, "")
}
