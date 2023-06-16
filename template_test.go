package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// template golang html
func SimpleHTML(writer http.ResponseWriter, request *http.Request) {
	//tanda titik akan digantikan oleh data yang di execute
	templateText := `<html><body>{{.}}</body></html>`

	// membuat nama template dengan cara klasik
	// t, err := template.New("SIMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }

	// membuat nama template dengan cara modern
	t := template.Must(template.New("SIMPLE").Parse(templateText))

	// execute template
	t.ExecuteTemplate(writer, "SIMPLE", "Hello HTML Template")
}

func TestSimpleHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	//memanggil func SimpleHTML
	SimpleHTML(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func SimpleHTMLFiles(writer http.ResponseWriter, request *http.Request) {
	//render file simple.gohtml pada package templates
	t := template.Must(template.ParseFiles("./templates/simple1.gohtml"))
	t.ExecuteTemplate(writer, "simple1.gohtml", "Hello HTML Template")
}

func TestSimpleHTMLFiles(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	//memanggil func SimpleHTML
	SimpleHTMLFiles(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TemplateDirectory(writer http.ResponseWriter, request *http.Request) {
	//render file simple.gohtml pada package templates
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(writer, "simple1.gohtml", "Hello HTML Template")
}

func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	//memanggil func SimpleHTML
	TemplateDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

/*
Mmebuat embed template
1. deklarasi var templates
2. buat func handler untuk TemplateEmbed
3. masukan var templates ke dalam template.Must
*/

func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {
	//render file simple.gohtml pada package templates
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))
	t.ExecuteTemplate(writer, "simple1.gohtml", "Hello HTML Template")
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	//memanggil func SimpleHTML
	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
