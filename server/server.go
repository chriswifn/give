package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Display the named template
func receiveDisplay(w http.ResponseWriter, page string, data interface{}) {
	serveTemplate("upload", uploadTemplate, w, data)
}

func receiveHelper(w http.ResponseWriter, r *http.Request) {
	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	// Get handler for filename, size and headers
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create file
	dst, err := os.Create(handler.Filename)
	defer dst.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func receiveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		receiveDisplay(w, "upload", nil)
		return
	}
	receiveHelper(w, r)
}

func receive() {
	var Url = `http://` + GetLocalIP() + `:8080/upload`
	fmt.Println(Url)
	http.HandleFunc("/upload", receiveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func receiveQr() {
	var Url = `http://` + GetLocalIP() + `:8080/upload`
	RenderString(Url)
	http.HandleFunc("/upload", receiveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveForm(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	jsn, _ := json.MarshalIndent(r.Form, "", "  ")
	fmt.Fprintf(w, string(jsn))
}

func serveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		serveForm(w, r)
		return
	}
	var err error
	wd, err := os.Getwd()
	if err != nil {
		log.Print(err)
		return
	}
	http.ServeFile(w, r, wd+r.URL.Path)
}

func serve() {
	var Url = `http://` + GetLocalIP() + `:8080`
	fmt.Println(Url)
	http.HandleFunc("/", serveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveQr() {
	var Url = `http://` + GetLocalIP() + `:8080`
	RenderString(Url)
	http.HandleFunc("/", serveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
