package golang_web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

/*
ALUR :
1. buat folder resource dan bikin 3 file yaitu index.css, index.html, index.js
2. jalankan file server di web menggunakan :
	- http://localhost:8080/static/index.css
	- http://localhost:8080/static/index.js
	-http://localhost:8080/static/index.html
*/

func TestFileServer(t *testing.T) {
	//membuat alamat direktori
	directory := http.Dir("./resources")
	//membuat file server
	fileServer := http.FileServer(directory)

	//membuat mux dan handler
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) //menghapus static

	//membuat server
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	//run server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

/*
membuat embed dari folder resources ke golang embed
jadi sekarang seluruh isi dari folder resources sudah ada di var resources

cara menjalankan di browser :
	- http://localhost:8080/static/index.css
	- http://localhost:8080/static/index.js
	-http://localhost:8080/static/index.html
*/

//go:embed resources
var resources embed.FS

func TestFileServerGoEmbed(t *testing.T) {
	//memasukan folder resources ke File System
	directory, _ := fs.Sub(resources, "resources")

	//value directory mejadi FileSystems
	fileServer := http.FileServer(http.FS(directory))

	//membuat mux dan handler
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) //menghapus static

	//membuat server
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	//run server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
