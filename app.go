package main

import (
	"log"
	"mime/multipart"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/sessions"
)

type UploadForm struct {
	Title      string                `form:"title"`
	TextUpload *multipart.FileHeader `form:"txtUpload"`
}

func main() {
	m := martini.Classic()

	//--- log  ---------------------------------------------------------
	m.Use(func(c martini.Context, log *log.Logger) {
		//log.Println("before a request")
		c.Next()
		//log.Println("after a request")
	})
	//--- /log ---------------------------------------------------------

	//m.Use(auth.Basic("test", "123"))

	//--- static  ---------------------------------------------------------
	m.Use(martini.Static("public"))
	//--- /static ---------------------------------------------------------

	//--- session ---------------------------------------------------------
	store := sessions.NewCookieStore([]byte("secret1234"))
	m.Use(sessions.Sessions("uzkhsess", store))

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
	//--- /session --------------------------------------------------------

	m.Get("/", func() string {
		return "Hello world!"
	})

	m.Get("/hello/:name", func(params martini.Params) string {
		return "Hello " + params["name"]
	})

	//--- fileupload  -----------------------------------------------------
	/*
				type UploadForm struct {
				    Title      string                `form:"title"`
				    TextUpload *multipart.FileHeader `form:"txtUpload"`
				}
	    **/
	m.Post("/", binding.MultipartForm(UploadForm{}), func(uf UploadForm) string {
		file, err := uf.TextUpload.Open()
		// ... you can now read the uploaded file
	})
	/***/
	//--- /fileupload -----------------------------------------------------

	m.RunOnAddr(":8091")
}
