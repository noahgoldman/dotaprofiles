package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"log"
)

func main() {
	InitDB()

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/static/:file", StaticFiles)
	router.POST("/crop/", CropImage)
	router.GET("/upload", Upload)
	router.POST("/upload", Upload)
	router.GET("/wsconnect", WebsocketHandler)

	http.ListenAndServe(":8080", router)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func StaticFiles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.ServeFile(w, r, "static/"+ps.ByName("file"))
}

func CropImage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	rect, err := GetRect(r)
	if err != nil {
		fmt.Fprintf(w, "error")
	}

	makeImages(rect, "static/cage.jpg")

	fmt.Fprintf(w, "done")
}

func Upload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method == "POST" {
		err := r.ParseMultipartForm(100000000)	
		if err != nil {
			panic(err)
		}

		m := r.MultipartForm

		files := m.File["picture"]
		if len(files) != 1 {
			log.Fatal("Only 1 file should ever be uploaded")
		}

		file, err := files[0].Open()
		defer file.Close()
		if err != nil {
			log.Printf("Failed to open uploaded file %s", files[0].Filename)
		}

	} else {
		t, err := template.ParseFiles("templates/upload.html")
		if err != nil {
			panic(err)
		}
		t.Execute(w, nil)
	}
}
