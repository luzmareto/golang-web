package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateDateMap(writer http.ResponseWriter, request *http.Request) {
	//render data ke package templates lalu ke file name.gohtml
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	//execute template dari parameter, file gohtmldan map
	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Tittle": "Template Data Map", //value tittle akan dikirim ke tittle html
		"Name":   "Luz",               ////value name akan dikirim ke body html
	})
}

func TestTemplateDateMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	//memanggil func SimpleHTML
	TemplateDateMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}

/*
TEMPLATE MENGGUNAKAN STRUCT :
1. bikin struct
2. bikin handler untuk struct tersebut
3. buat unit test

Menambahkan struct didalam struct :
1. tambahkan alamt pada html di h2
2. deklarasi struct baru
3. masukan struct tersebut ke dalam field struct utama
4. masukan struct tersebut pada handler
*/

type Address struct {
	Street string
}

// struct utama
type Page struct {
	Title   string
	Name    string
	Address Address
}

// handler
func TemplateDateStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "luz",
		Address: Address{
			Street: "Jalan Belum Ada",
		},
	})
}

func TestTemplateDateStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	//memanggil func SimpleHTML
	TemplateDateStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))
}
