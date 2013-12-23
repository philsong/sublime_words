sublime_words
=============

wei yan wei yu web2.0 site

update: I do not know how to let martini binding and sessions module work together, could someone help me?
m.Post("/signin", binding.Bind(UserSignin{}), func(usersingin UserSignin) string {
		if userpost, err := QueryUser(usersingin); err != false {
			// find user in db, how to bind a cookie/session in sessions module in here?
		} else {
			return "query DB failed"
		}
	})
