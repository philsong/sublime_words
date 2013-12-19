package main

import (
	"github.com/codegangsta/martini"
	//"github.com/codegangsta/martini-contrib/auth"
	"github.com/codegangsta/martini-contrib/sessions"
	"sublime"
)

func main() {
	m := martini.Classic()
	sublime.Mysqltest()
	m.Get("/", func() string {
		return "Hello weiyan!"
	})
	m.Get("/signup", func() string {
		return "Hello signup!"
	})
	m.Get("/signin", func() string {
		return "Hello signin!"
	})
	m.Get("/inspect", func() string {
		return "Hello inspect!"
	})
	m.Get("/add", func() string {
		return "Hello add!"
	})
	m.Post("/signup", func() string {
		return "Hello signup!"
	})
	m.Post("/signin", func() string {
		return "Hello signin!"
	})
	m.Post("/inspect", func() string {
		return "Hello inspect!"
	})
	m.Post("/add", func() string {
		return "Hello add!"
	})

	store := sessions.NewCookieStore([]byte("secret123"))
	m.Use(sessions.Sessions("my_session", store))

	m.Get("/set", func(session sessions.Session) string {
		session.Set("hello", "world")
		return "OK"
	})

	m.Get("/get", func(session sessions.Session) string {
		v := session.Get("hello")
		if v == nil {
			return ""
		}
		return v.(string)
	})

	//m.Use("/admin", auth.Basic("username", "secretpassword"))
	m.Use(martini.Static("assets"))
	m.Run()
}
