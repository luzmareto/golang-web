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
ALUR IF EXPRESSION WITH HTML:
1. buat file file if.gohtml pada package templates
2. pada file tersebut, buat If statement
3. buat struct di luar func handle
4. bikin handle func
*/

type Address1 struct {
	Street string
}

// struct utama
type Page1 struct {
	Title   string
	Name    string
	Address Address1
}

// handle func
func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	//panggil nama package dan filenya
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(writer, "if.gohtml", Page1{
		Title: "Template Data Struct",
		Name:  "luz",
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	//memanggil func SimpleHTML
	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body)) //Hello Luz
}

/*
ALUR OPERATOR PERBANDINGAN MENGGUNAKAN HTML:
1. Buat file comparator.gohtml di package templat
2. pada file tersebut, harus ada operator perbandingan
3. bikin handle func
4. biki func unit testing
*/

// handle func
func TemplateActionOperator(writer http.ResponseWriter, request *http.Request) {
	//panggil nama package dan filenya
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	//execute menggunakan map agar lebih mudah
	t.ExecuteTemplate(writer, "comparator.gohtml", map[string]interface{}{
		"Tittle":     "Template Actio Operator",
		"FinalValue": 90,
	})
}

func TestTemplateActionOperator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	//memanggil func TemplateActionOperator
	TemplateActionOperator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body)) //Good
}

/*
ALUR RANGE MENGGUNAKAN HTML:
1. Buat file range.gohtml di package templat
2. pada file tersebut, harus ada syntax range
3. bikin handle func
4. bikin func unit testing dan run
*/

// handle func
func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	//panggil nama package dan filenya
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	//execute range menggunakan map inteface dan slice agar lebih mudah
	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Tittle": "Template Actio Range",
		"Hobbies": []string{
			"Game", "Read", "coding",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	//memanggil func TemplateActionRange
	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
	/*
		OUTPUT :
		<h1>0 - Game</h1>
		<h1>1 - Read</h1>
		<h1>2 - coding</h1>
	*/
}

/*
ALUR WITH MENGGUNAKAN HTML:
1. Buat file address.gohtml di package templat
2. pada file tersebut, harus ada syntax with (else adalah optional)
3. bikin handle func
4. biki func unit testing
*/

// handle func
func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	//panggil nama package dan filenya
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))
	//execute menggunakan map interface agar lebih mudah penggunakan with
	t.ExecuteTemplate(writer, "address.gohtml", map[string]interface{}{
		"Tittle": "Template Action With",
		"Name":   "Luz",
		"Address": map[string]interface{}{
			"Street": "Jalan Belum Ada",
			"City":   "Jakarta",
		},
	})
}

func TestTTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	//memanggil func TemplateActionWith
	TemplateActionWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body)) //Good
}
