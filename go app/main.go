package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// refer: https://github.com/zupzup/golang-http-file-upload-download/blob/main/main.go
// NOTE : ensure that ./tmp directory already exists

const maxUploadSize = 1024 * 1024 // 1 mb
const uploadPath = "./tmp/"

func main() {

	newpath := filepath.Join(".", "tmp")
	_ = os.MkdirAll(newpath, os.ModePerm)

	http.HandleFunc("/upload", uploadFileHandler())

	http.HandleFunc("/run_beast", runBeastHandler())

	fs := http.FileServer(http.Dir(uploadPath))
	http.Handle("/files/", http.StripPrefix("/files", fs))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func runBeastHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// note that the name of the file will be 'beast.build'
		if r.Method == "GET" {
			return
		}

		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			renderError(w, "CANNOT PARSE FORM", http.StatusInternalServerError)
			return
		}

		file, fileHeader, err := r.FormFile("uploadFile")
		if err != nil {
			renderError(w, "INVALID FILE", http.StatusBadRequest)
			return
		}

		defer file.Close()

		fileSize := fileHeader.Size
		if fileSize > maxUploadSize {
			renderError(w, "FILE TOO BIG", http.StatusBadRequest)
			return
		}

		// fmt.Print(fileHeader.Filename)

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			renderError(w, "INVALID FILE", http.StatusBadRequest)
			return
		}

		// ignore file type right now

		// fileName := randToken(12)
		beastFileName := "beast.build"
		newPath := filepath.Join(uploadPath, beastFileName)

		// write file
		newFile, err := os.Create(newPath)

		if err != nil {
			fmt.Print(newPath)
			renderError(w, "CANNOT WRITE FILE", http.StatusInternalServerError)
			return
		}

		defer newFile.Close()

		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			fmt.Print(newPath)
			renderError(w, "CANNOT WRITE FILE", http.StatusInternalServerError)
			return
		}

		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		fmt.Println(exPath)

		// output, err := exec.Command("/app/tmp/beast").Output()
		// if err != nil {
		// 	fmt.Print(output, err)
		// 	renderError(w, fmt.Sprint(err), http.StatusInternalServerError)
		// 	return
		// }

		w.Write([]byte("SUCCESSFUL BUILD"))
	})
}

func uploadFileHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			return
		}

		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			renderError(w, "CANNOT PARSE FORM", http.StatusInternalServerError)
			return
		}

		file, fileHeader, err := r.FormFile("uploadFile")
		if err != nil {
			renderError(w, "INVALID FILE", http.StatusBadRequest)
			return
		}

		defer file.Close()

		fileSize := fileHeader.Size
		if fileSize > maxUploadSize {
			renderError(w, "FILE TOO BIG", http.StatusBadRequest)
			return
		}

		// fmt.Print(fileHeader.Filename)

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			renderError(w, "INVALID FILE", http.StatusBadRequest)
			return
		}

		// ignore file type right now

		// fileName := randToken(12)
		newPath := filepath.Join(uploadPath, fileHeader.Filename)
		// newPath := fileHeader.Filename
		// write file
		newFile, err := os.Create(newPath)

		if err != nil {
			fmt.Print(newPath)
			renderError(w, "CANNOT WRITE FILE", http.StatusInternalServerError)
			return
		}

		defer newFile.Close()

		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			fmt.Print(newPath)
			renderError(w, "CANNOT WRITE FILE", http.StatusInternalServerError)
			return
		}

		w.Write([]byte("SUCCESS"))
	})
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
