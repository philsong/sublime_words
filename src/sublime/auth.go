package sublime

import (
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/codegangsta/martini-contrib/sessions"
	"net/http"
	"regexp"
)

type UserSignin struct {
	UPwd       string `form:"upwd" json:"upwd" binding:"required"`
	UEmail     string `form:"uemail" json:"uemail"`
	unexported string `form:"-"` // skip binding of unexported fields
}

// This method implements binding.Validator and is executed by the binding.Validate middleware
func (up UserSignin) Validate(errors *binding.Errors, req *http.Request) {
	if len(up.UPwd) < 4 {
		errors.Fields["upwd"] = "Too short; minimum 4 characters"
	} else if len(up.UPwd) > 30 {
		errors.Fields["upwd"] = "Too long; maximum 30 characters"
	}
	if m, _ := regexp.MatchString("^[0-9a-zA-Z]+$", up.UPwd); !m {
		errors.Fields["upwd"] = "only alpha number"
	}

	if len(up.UEmail) < 4 {
		errors.Fields["uemail"] = "Too short; minimum 4 characters"
	} else if len(up.UEmail) > 30 {
		errors.Fields["uemail"] = "Too long; maximum 30 characters"
	}
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, up.UEmail); !m {
		errors.Fields["uemail"] = "It is not correct email adress"
	}
}

type UserPost struct {
	UName      string `form:"uname" json:"uname" `
	UPwd       string `form:"upwd" json:"upwd" binding:"required"`
	UEmail     string `form:"uemail" json:"uemail"`
	unexported string `form:"-"` // skip binding of unexported fields
}

// This method implements binding.Validator and is executed by the binding.Validate middleware
func (up UserPost) Validate(errors *binding.Errors, req *http.Request) {
	if len(up.UName) < 4 {
		errors.Fields["uname"] = "Too short; minimum 4 characters"
	} else if len(up.UName) > 30 {

		errors.Fields["uname"] = "Too long; maximum 30 characters"
	}
	if m, _ := regexp.MatchString("^[0-9a-zA-Z]+$", up.UName); !m {
		errors.Fields["uname"] = "only alpha number"
	}

	if len(up.UPwd) < 4 {
		errors.Fields["upwd"] = "Too short; minimum 4 characters"
	} else if len(up.UPwd) > 30 {
		errors.Fields["upwd"] = "Too long; maximum 30 characters"
	}
	if m, _ := regexp.MatchString("^[0-9a-zA-Z]+$", up.UPwd); !m {
		errors.Fields["upwd"] = "only alpha number"
	}

	if len(up.UEmail) < 4 {
		errors.Fields["uemail"] = "Too short; minimum 4 characters"
	} else if len(up.UEmail) > 30 {
		errors.Fields["uemail"] = "Too long; maximum 30 characters"
	}
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, up.UEmail); !m {
		errors.Fields["uemail"] = "It is not correct email adress"
	}
}

func Signup(m *martini.ClassicMartini) {
	m.Get("/signup", func(r render.Render) {
		r.HTML(200, "signup", nil)
	})

	m.Post("/signup", binding.Bind(UserPost{}), func(userpost UserPost) string {
		// This function won't execute if there were errors
		if InsertUser(userpost) {
			//http.Redirect(w, r, "/profile", http.StatusFound)
			//return userpost.UName
			s := "<html>"
			s += userpost.UName
			s += `,恭喜你注册成功,请 <a href="/signin">登录</a></html>`
			return s
			//return `"<html>恭喜你", userpost.UName, "通过验证,请 <a href="/signin">登录</a></html>"`
		} else {
			return "insert DB failed"
		}
	})
}

func Signin(m *martini.ClassicMartini) {
	m.Get("/signin", func(r render.Render) {
		r.HTML(200, "signin", nil)
	})
	m.Post("/signin", func(session sessions.Session) {
		fmt.Println("session set..")
		session.Set("hello", "world")
		//	return "OK"
	})
	m.Post("/signin", binding.Bind(UserSignin{}), func(usersingin UserSignin) string {
		// This function won't execute if there were errors
		if userpost, err := QueryUser(usersingin); err != false {
			//http.Redirect(w, r, "/profile", http.StatusFound)
			//return userpost.UName
			s := "<html>"
			s += userpost.UName
			s += `,恭喜你通过验证,请 <a href="/">回首页</a></html>`

			//sessions.Session.Set(userpost.UEmail, userpost.UName)
			//sessions.Session.("hello", "world")
			fmt.Println(s)
			//martini.Context.Next()
			return s
		} else {
			return "query DB failed"
		}
	})

}
