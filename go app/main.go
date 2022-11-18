package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// refer: https://github.com/zupzup/golang-http-file-upload-download/blob/main/main.go
// NOTE : can also use gin-gonic or gorilla mux to do the job.

// TODO : handle persist code (daku_mantra) before communicating with user

const uploadPath = "./tmp/"

func main() {
	// NOTE : ensure that ./tmp directory already exists
	newpath := filepath.Join(".", "tmp")
	_ = os.MkdirAll(newpath, os.ModePerm)

	// 1. route to upload file to server
	http.HandleFunc("/upload_file", uploadFileHandler())

	// 2. route to download file from server
	// fs := http.FileServer(http.Dir(uploadPath))
	// http.Handle("/download_file/", http.StripPrefix("/files", fs))

	// 3. route to build file and run stuff
	http.HandleFunc("/run_beast", runBeastHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func uploadFileHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if r.Method == "GET" {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")

			resp := make(map[string]string)
			resp["message"] = "The route does not accept GET requests. Try POST request."

			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}

			w.Write(jsonResp)
			return
		}

		file, fileHeader, err := r.FormFile("upload_file")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")

			resp := make(map[string]string)
			resp["message"] = "Invalid File"

			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}

			w.Write(jsonResp)
			return
		}

		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")

			resp := make(map[string]string)
			resp["message"] = "File Cannot be read"

			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}

			w.Write(jsonResp)
			return
		}

		newPath := filepath.Join(uploadPath, fileHeader.Filename)

		// write file
		newFile, err := os.Create(newPath)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")

			resp := make(map[string]string)
			resp["message"] = "Cannot write file to disk"

			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}

			w.Write(jsonResp)
			return
		}

		defer newFile.Close()

		if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")

			resp := make(map[string]string)
			resp["message"] = "Cannot write file to disk (2)"

			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}

			w.Write(jsonResp)
			return
		}

		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "application/json")

		resp := make(map[string]string)
		resp["message"] = newPath

		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.Write(jsonResp)
		return
	})
}

func runBeastHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "GET" {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")

			resp := make(map[string]string)
			resp["message"] = "The route does not accept GET requests. Try POST request."

			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}

			w.Write(jsonResp)
			return
		}

		reqBody, err := ioutil.ReadAll(r.Body)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")

			resp := make(map[string]string)
			resp["message"] = "Invalid Request"

			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}

			w.Write(jsonResp)
			return
		}

		fmt.Printf("%s\n", reqBody)
		var dat map[string]interface{}
		if err := json.Unmarshal(reqBody, &dat); err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		dat["data"] = "dummy"

		fmt.Println(dat)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "application/json")

		jsonResp, err := json.Marshal(dat)

		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.Write(jsonResp)
		return

	})
}
