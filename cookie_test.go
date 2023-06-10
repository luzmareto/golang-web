package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// membuat func handler untuk set cookie
func SetCookie(writer http.ResponseWriter, request *http.Request) {
	// membuat cookie
	cookie := new(http.Cookie)
	//set nama cookie
	cookie.Name = "luz-mareto-name"
	//mengambil value dari request
	cookie.Value = request.URL.Query().Get("name")
	//set path
	cookie.Path = "/"

	//memasukan data cookie ke writer
	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "succsess create cookie")
}

// handler untuk mendapatkan/membaca cookie
func GetCookie(writer http.ResponseWriter, request *http.Request) {
	//cookie artinya mengambil data berdasarkan nama
	cookie, err := request.Cookie("luz-mareto-name") //note! cookies artinya mengambil semua data
	if err != nil {
		fmt.Fprintf(writer, "no cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(writer, "hello %s", name)
	}
}

/*
Data cookies akan di simpan di browser/appliaction client bukan di server
ALUR MENJALANKAN COOKIE
1. jalankan di terminal : go test -v -run=TestCookie
2. jalankan localhost:8080 di web browser
3. klik kanan > inspect > application
4. pada pojok kanan, gulir kebawah > Cookies > http://http://localhost:8080/
5. masukan cookie. contoh : ganti url menjadi localhost:8080/set-cookie?name=luz mareto

Untuk mengetahui riwayat cookie, pada inspect masuk ke menu source
cek riwayat response header, masuk ke menu network > all > set-cookie?name... >gulir ke bawah > response header

cara menjalankan Get Cookie :
ganti url menjadi http://localhost:8080/get-cookie

cara cek
*/

func TestCookie(t *testing.T) {
	//membuat handler mux
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie) //func SetCookie
	mux.HandleFunc("/get-cookie", GetCookie) //func SetCookie

	//membuat server
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	// menjalankan server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	//unit test request cookie
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=luz", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("cookie %s:%s \n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	// unit test untuk mengirim/membaca cookie
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "luz-mareto-name"
	cookie.Value = "luz"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body)) //hello luz
}
