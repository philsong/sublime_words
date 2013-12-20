package sublime

import (
	"github.com/codegangsta/martini"
	//"github.com/codegangsta/martini-contrib/auth"
	"github.com/codegangsta/martini-contrib/render"
)

func Route(m *martini.ClassicMartini) {
	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", nil)
	})

	//post handlers
	Signup(m)
	Signin(m)

	m.Get("/inspect", func() string {
		return "Hello inspect!"
	})
	m.Get("/add", func() string {
		return "Hello add!"
	})

	m.Post("/inspect", func() string {
		return "Hello inspect!"
	})
	m.Post("/add", func() string {
		return "Hello add!"
	})
}
