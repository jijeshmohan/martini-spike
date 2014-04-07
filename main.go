package main

import (
	"blogs/models"
	"blogs/modules/middleware"
	"blogs/routes"
	"log"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

var (
	m *martini.ClassicMartini
)

func init() {
	m = martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))
	m.Use(middleware.InitContext())

	if _, err := models.InitDb(); err != nil {
		log.Fatalln("Unable to initialize database")
	}

	routes.InitRoutes(m)
}

func main() {
	m.Run()
}
