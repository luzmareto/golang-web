package golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

/*
ALUR :
1. bikin file ok.html dan notfound.html pada package resources
2. bikin func handle untuk Serve File
3. bikin unit test untuk func tersebut
4. jalankan unit testing
5. jalankan browser dengan url http://localhost:8080/?name=masukan_nama_kita
*/

func ServeFile(writer http.ResponseWriter, request *http.Request) {

	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resources/ok.html") //name file
	} else {
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

func TestServeFile(t *testing.T) {
	//membuat server dengan handler dari func ServeFile
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	//run server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
cara membuat embed :
1. deklarasi resourceOk dan resourceNotFound beserta format embednya
2. bikin func untuk implement embed
3. membuat func unit test untuk func ServeFileEmbed
4. runing unit test tersebut
5. jalankan browser dengan url localhost:8080/?name=masukan_nama_kita
*/

//go:embed resources/ok.html
var resourceOk string

//go:embed resources/notfound.html
var resourceNotFound string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request) {
	//masukan deklarasi embed
	if request.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, resourceOk)
	} else {
		fmt.Fprint(writer, resourceNotFound)
	}
}

func TestServeFileEmbed(t *testing.T) {
	//membuat server dengan handler dari func Embed
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	//run server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
