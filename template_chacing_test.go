package golang_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
ALUR CHACING :
1. deklarasi global variable yang sudah di embed global
2. deklarasi global variable untuk parsing
3. func handler
4. unit test

*/
//go:embed templates/*.gohtml
var templates embed.FS

// global parsing
var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

// Hanler
func TemplateChacing(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "hello Template Chacing")
}

func TestTemplateChacing(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	//memanggil func SimpleHTML
	TemplateChacing(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
