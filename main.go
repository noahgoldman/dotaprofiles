package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/noahgoldman/dotaprofiles/upload"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	InitDB()
	upload.AWSInit()

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/static/:file", StaticFiles)
	router.POST("/make_images/:id", MakeImageHandler)
	router.POST("/upload", Upload)

	http.ListenAndServe(":8080", router)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles("templates/upload.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}

func StaticFiles(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	http.ServeFile(w, r, "static/"+params.ByName("file"))
}

func MakeImageHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		sendInternalError(w, err)
		return
	}

	ps, err := getPictureSet(id)
	if err != nil {
		log.Print(err)
		http.NotFound(w, r)
		return
	}

	r.ParseForm()
	rect, err := GetRect(r)
	if err != nil {
		sendInternalError(w, err)
		return
	}

	file, err := GetImageFile(ps.original, func(s string) (io.Reader, error) { return nil, nil })
	MakeImages(rect, file)

	fmt.Fprintf(w, "done")
}

func Upload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := r.ParseMultipartForm(100000000)
	if err != nil {
		panic(err)
	}

	m := r.MultipartForm

	files := m.File["picture"]
	if len(files) != 1 {
		sendHTTPError(w, "Only 1 file should ever be uploaded")
		return
	}

	filename := files[0].Filename
	file, err := files[0].Open()
	defer file.Close()
	if err != nil {
		err_string := fmt.Sprintf("Failed to open uploaded file %s", files[0].Filename)
		sendHTTPError(w, err_string)
		return
	}

	ps, err := newPictureSet(filename)
	if err != nil {
		sendInternalError(w, err)
		return
	}
	if ps == nil {
		fmt.Printf("wow err")
	}

	fmt.Printf("%s", ps.original)

	// Create the local destination file
	dest, err := CreateImageFile(ps.original)
	if err != nil {
		sendInternalError(w, err)
		return
	}
	defer dest.Close()

	// Copy the file to the local image store directory
	_, err = io.Copy(dest, file)
	if err != nil {
		sendInternalError(w, err)
		return
	}

	// Seek back to the start of the file so that we can read it again
	_, err = file.Seek(0, 0)
	if err != nil {
		sendInternalError(w, err)
		return
	}

	// Upload the file to Amazon S3
	err = upload.Upload_S3(file, ps.original)
	if err != nil {
		sendInternalError(w, err)
		return
	}

	fmt.Fprintf(w, "%s", upload.GetURL(ps.original))
}

func sendHTTPError(w http.ResponseWriter, err_string string) {
	log.Print(err_string)
	http.Error(w, err_string, 500)
}

func sendInternalError(w http.ResponseWriter, err error) {
	log.Print(err)
	http.Error(w, "Internal server error", 500)
}
