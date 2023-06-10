package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

/*
ALUR HANDLER :
1. bikin var handler dengan type data HandlerFunc dan valuenya anonym func
2. w/writer dalam anonym func adalah respon yang kita berikan ke client
3. r/request dalam anonym func adalah permintaan dari client
4. bikin server dan masukan handler ke dalam server

Jika semua sudah jalan. buka alamat localhost:8080 di web browser
*/

func TestHandler(t *testing.T) {

	//handler with anonym func
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {

		//logic web
		fmt.Fprint(w, "hello world")
	}

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: handler,
	}

	//menjalankan server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
ServeMux adalah menggabungkan handler atau multiple ServeMux

ALUR SERVEMUX:
1. deklarasi var yang isinya http.NewServeMux()
2. bikin handlefunc dan masukan simbol / setelah itu masukan anonym func
3. masukan logics
4. bikin handlefunc lagi namun masukan endpoind atau url yang berbeda dari yg sebelumnya
5. kita bisa bikin handlefunc sebanyak mungkin
6. bikin var server yang valuenya http.Server{)
7. gabungkan semua handlefunc pada var server dengan cara masukan var mux pada field Handler
8. jalankan server
*/

func TestServeMux(t *testing.T) {

	//declaration ServeMux
	mux := http.NewServeMux()

	//handle func with anonym func and logic
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})

	//endpoind/url hi
	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi")
		// http://localhost:8080/hi/ mencetak hi

	})

	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "image")
		// http://localhost:8080/images/ mencetak image

	})

	//endpoind/url hi
	mux.HandleFunc("/images/thumbnails/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Thumbnail")
		/* http://localhost:8080/images/thumbnails/ mencetak Thumbnail
		   http://localhost:8080/images/thumbnailssasd mencetak url /images.
		*/
	})

	//gabungkan handlefunc di server
	server := http.Server{
		Addr:    "localhost:9090",
		Handler: mux,
	}

	//menjalankan server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
CONTOH REQUEST METHOD : POST, GET, PUT DELETE

*/

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
	}

	//create server
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	//run server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
