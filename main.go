package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
)

func main() {
	InitDB()

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/static/:file", StaticFiles)
	router.POST("/crop/", CropImage)

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
