package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

/*
ALUR TEMPLATE LAYOUT :
1. buat file header.gohtml, footer.gohtml, layout.gohtml pada package templates
2. buat func handler dan panggil nama package dan file tersebut
*/

// func handle
func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	//panggil nama package dan filenya
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/footer.gohtml",
		"./templates/layout.gohtml",
	))
	//execute menggunakan map interface dan panggil nama file layout
	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Luz",
	})
}

func TestTTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	//memanggil func TemplateLayout
	TemplateLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body)) //Good
}
