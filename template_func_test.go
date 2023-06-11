package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*
ALUR PEMBUATAN TEMPLATE FUNC GO-HTML
1. buat struct
2. buat funcuntuk menampung struct tersebut
3. buat func handler template func
4. buat unit testing
*/

type MyPage struct {
	Name string
}

// func SayHello mempunyai parameter name dan return string untuk memanggil struct MyPage
func (mypage MyPage) SayHello(name string) string {

	//membuat return string
	return "Hello " + name + ", My Name Is " + mypage.Name
}

// handle func
func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "Budi"}}`))

	//execute template menggunakan parameter writter,FUCNTION dan struct MyPage
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Luz",
	})
}

func TestTTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	//memanggil func TemplateLayout
	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

// simple global func
func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name}}`))

	//execute template menggunakan parameter writter,FUCNTION dan struct MyPage
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "Luz",
	})
}

func TestTTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	//memanggil func TemplateLayout
	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

/*
ALUR PEMBUATAN GLOBAL FUNCTION :
1. buat fun handle TemplateFunctionCreateGlobal
2. gunakan anonym func pada saat register nama func
3. buat unit testing
4. maka value dari  field name akan berubah menjadi huruf kapital
*/

func TemplateFunctionCreateGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")

	//upper adalah nama func
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ upper .Name}}`))

	//execute template menggunakan parameter writter,FUCNTION dan struct MyPage
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "luz mareto",
	})
}

func TestTTemplateFunctionCreateGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	//memanggil func TemplateFunctionCreateGlobal
	TemplateFunctionCreateGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
