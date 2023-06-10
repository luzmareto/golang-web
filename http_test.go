package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
ALUR TEST :
1. bikin contract handler yang akan di test
2.buat func unit test
3. bikin var requesst yang valuenya : request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
4. bikin recorder yang isinya
5. recorder := httptest.NewRecorder()
6. panggil func yang isinya contract dan masukan 2 parameternya dari var di func unit test
*/

// make contract handler
func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "hello world")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil) //nil artinya tidak pakai body
	recorder := httptest.NewRecorder()

	//parameter func HelloHandler diisi dari var di atas
	HelloHandler(recorder, request)

	//membaca hasilnya
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodystring := string(body)

	fmt.Println(bodystring)
}
