package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request) {
	//parsing
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	request.PostFormValue("first_name")

	//mengambil data
	firstName := request.PostForm.Get("first_name")
	lastName := request.PostForm.Get("last_name")

	fmt.Fprintf(writer, "hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requesBody := strings.NewReader("first_name=luz&last_name=mareto")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requesBody)
	request.Header.Add("content-type", "application/x-www-form-urlencode")

	recorder := httptest.NewRecorder()

	//memberi value pada parameter func FormPost
	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
