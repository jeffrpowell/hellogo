package web

import (
	"embed"
	"html/template"
	"io"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jeffrpowell/hellogo/internal/constants"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

//go:embed dist/*
var staticFiles embed.FS
var (
	helloWorld = parseTemplate("dist/helloworld.html", template.FuncMap{"join": strings.Join})
)

func init() {
	constants.ROUTER.HandleFunc("/static/{pathname...}", staticHandler).Methods("GET")
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	filePath := mux.Vars(r)["pathname..."]
	// Open the file from the embedded file system
	file, err := staticFiles.Open("dist/" + filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Get the file extension
	ext := filepath.Ext(filePath)
	// Set the content type based on the file extension
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		// If the content type is unknown, default to "application/octet-stream"
		contentType = "application/octet-stream"
	}

	// Set the content type header
	w.Header().Set("Content-Type", contentType)

	// Copy the file content to the response writer
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Error serving file", http.StatusInternalServerError)
		return
	}
}

func minifyTemplates(funcs template.FuncMap, filenames ...string) (*template.Template, error) {
	m := minify.New()
	m.AddFunc("text/html", html.Minify)

	var tmpl *template.Template
	for _, filename := range filenames {
		name := filepath.Base(filename)
		if tmpl == nil {
			tmpl = template.New(name).Funcs(funcs)
		}

		b, err := staticFiles.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		mb, err := m.Bytes("text/html", b) //BUG: lower-cases go interpolation tags
		if err != nil {
			return nil, err
		}

		tmpl, err = tmpl.Parse(string(mb))
		if err != nil {
			return nil, err
		}
	}
	return tmpl, nil
}

func parseTemplate(file string, funcs template.FuncMap) *template.Template {
	return template.Must(minifyTemplates(funcs, "dist/root.html", file))
}

type globalWebParams struct {
	JsFile string
}

// Hello World page

type helloWorldParams struct {
	globalWebParams
	Name   string
	Colors []string
}

func HelloWorldParams(name string, gradient constants.ColorGradient) helloWorldParams {
	return helloWorldParams{
		globalWebParams{
			JsFile: "helloworld",
		},
		name,
		gradient.Colors,
	}
}

func HelloWorldPage(w io.Writer, params helloWorldParams) {
	if err := helloWorld.Execute(w, params); err != nil {
		log.Print(err)
	}
}

// the following code is stubbed to allow for the example handler code to build

func LoginPage(w io.Writer) {

}
func ListsPageParams(lists []string, isAdmin bool) string {
	return ""
}
func ListsPage(w io.Writer, params string) {

}
func EditListParams(list string, isAdmin bool) string {
	return ""
}
func EditListPage(w io.Writer, params string) {

}
