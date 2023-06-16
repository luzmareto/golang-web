package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
XSS (CROS STIE SCRIPTING) adalah metode mengamankan cookie
ALUR :
1. file post.gohtml di package templates yang isinya parsing title dan body
2. handler menggunakan global parsing myTemplates dari file template_chacing_test.go
3. unit testing
4. unit test server dan jalankan url localhost:8080 di browser

*/

// handler
func TemplateAutoEscpe(writer http.ResponseWriter, request *http.Request) {
	//execute menggunakan global parsing myTemplates
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  "<p>Ini adalah Body</p>",
	})
}

func TestTemplateAutoEscpae(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscpe(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TestTemplateAutoEscpeServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscpe),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
Cara mematikan auto render template
1. func handler
2. pada bagian body, masukan perintah render seperti HTML,CSS,JS,DLL
*/

func TemplateAutoEscpeDisbled(writer http.ResponseWriter, request *http.Request) {
	//execute menggunakan global parsing myTemplates
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title": "Template Auto Escape",
		"Body":  template.HTML("<p>Ini Adalah Body</p>"),
	})
}

func TestTemplateAutoEscpeDisbled(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscpeDisbled(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

func TestTemplateAutoEscpeDisbledServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscpeDisbled),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
