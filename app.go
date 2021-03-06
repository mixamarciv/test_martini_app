package main

import (
	"log"
	"mime/multipart"

	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/sessions"

	"strconv"
)
import "html/template"

func init() {
	InitLog()
	InitDb()
}

func main() {
	var m *martini.ClassicMartini = martini.Classic()
	//m := martini.Classic() // martini.ClassicMartini

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
	type UploadForm struct {
		Title      string                `form:"title"`
		TextUpload *multipart.FileHeader `form:"txtUpload"`
	}
	m.Post("/", binding.MultipartForm(UploadForm{}), func(uf UploadForm) string {
		//file, err := uf.TextUpload.Open()
		// ... you can now read the uploaded file
		return ""
	})
	/***/
	//--- /fileupload -----------------------------------------------------

	//--- render  ---------------------------------------------------------
	m.Use(render.Renderer(render.Options{
		Directory:  "templates",                // Specify what path to load the templates from.
		Layout:     "maintemplate",             // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		Funcs:      []template.FuncMap{},       // Specify helper function maps for templates to access.
		Delims:     render.Delims{"{{", "}}"},  // Sets delimiters to the specified strings.
		Charset:    "UTF-8",                    // Sets encoding for json and html content-types. Default is "UTF-8".
		IndentJSON: true,                       // Output human readable JSON
	}))

	// This is set the Content-Type to "text/html; charset=ISO-8859-1"
	m.Get("/html", func(r render.Render) {
		r.HTML(200, "hello", map[string]interface{}{"hello": "world"})
	})

	// This is set the Content-Type to "application/json; charset=ISO-8859-1"
	m.Get("/api", func(r render.Render) {
		r.JSON(200, map[string]interface{}{"hello": "world"})
	})
	//--- /render  --------------------------------------------------------

	//--- all  ---------------------------------------------------------
	m.Get("/test", func(r render.Render, session sessions.Session) {
		v := session.Get("cnt")
		var a string
		if v == nil {
			a = "0"
		} else {
			a = v.(string)
		}
		cnt, err := strconv.Atoi(a)
		LogPrintErrAndExit("strconv.Atoi(v) error: \n"+a+"\n\n", err)
		session.Set("cnt", strconv.Itoa(cnt+1))
		r.HTML(200, "test", map[string]interface{}{"hello": "world", "cnt": a})
	})
	//--- /all  --------------------------------------------------------

	m.RunOnAddr(":8091")
}
