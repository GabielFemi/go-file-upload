package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		_, _ = fmt.Fprintf(w, "Uploading file\n")
		// 1. parse input, type multipart/form-data
		_ = r.ParseMultipartForm(10 << 20)

		// 2. retrieve file from parsed form-data
		file, handler, err := r.FormFile("myFile")
		if err != nil {
			fmt.Println("Error retrieving file from form-data")
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Printf("Uploaded file: %+v\n", handler.Filename)
		fmt.Printf("File size: %+v\n", handler.Size)
		fmt.Printf("MIME header: %+v\n", handler.Header)
		// 3. write temporary file

		tempFile, err := ioutil.TempFile("temp-images", "wallpaper-*.*")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		tempFile.Write(fileBytes)

		_, _ = fmt.Fprintf(w, "Successfully uploaded file!")
	} else {
		render (w, "index.html", r)
	}


}

func setUpRoutes() {
	http.HandleFunc("/upload", uploadFile)
	fmt.Println("Listening on 127.0.0.1:8000")
	http.ListenAndServe(":8000", nil)

}

func main() {
	setUpRoutes()

}

func render(w http.ResponseWriter, tmpl string, r *http.Request) {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		log.Fatalln(err)
	}
	err = t.Execute(w, nil)
}