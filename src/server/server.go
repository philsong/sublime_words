package main

import (
	"github.com/codegangsta/martini"
	//"github.com/codegangsta/martini-contrib/auth"
	"github.com/codegangsta/martini-contrib/sessions"
	"log"
	"sublime"
)

func main() {
	m := martini.Classic()

	// log before and after a request
	m.Use(func(c martini.Context, log *log.Logger) {
		log.Println("before a request")

		c.Next()

		log.Println("after a request")
	})
	sublime.Route(m)

	store := sessions.NewCookieStore([]byte("secret_words_key_xxx"))
	m.Use(sessions.Sessions("weiyan_session", store))

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
