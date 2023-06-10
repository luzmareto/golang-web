package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name") //get untuk mendapatkan 1 data

	//jika tidak ada data, outputnya string kosong
	if name == "" {
		fmt.Fprintf(writer, "hello")
	} else {
		fmt.Fprintf(writer, "hello %s", name)
		//gunakan format Fprintf jika mencetak hasil query/request dan respons

	}
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=luz", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	respon := recorder.Result()
	body, _ := io.ReadAll(respon.Body)

	fmt.Println(string(body))
}

func MultipleQueryParameter(writer http.ResponseWriter, request *http.Request) {
	//gunakan get untuk mendapatkan 1 data
	firstName := request.URL.Query().Get("first_name")
	lastName := request.URL.Query().Get("last")

	//gunakan format Fprintf jika mencetak hasil query/request dan respons
	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestMultipleQueryParameter(t *testing.T) {
	//membuat multiple parameter di url
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?first_name=luz&last_name=Mareto", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParameter(recorder, request)

	respon := recorder.Result()
	body, _ := io.ReadAll(respon.Body)

	fmt.Println(string(body))
}

func MultipleParameterValues(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	names := query["name"] //gunakan slice untuk mendapatkan banyak data dari 1 parameter
	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestMultipleParameterValues(t *testing.T) {
	//membuat multiple parameter di url dengan url yang sama
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=luz&name=mareto&name=luzmareto", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	respon := recorder.Result()
	body, _ := io.ReadAll(respon.Body)

	fmt.Println(string(body))
}
