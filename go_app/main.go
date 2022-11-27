package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

// refer: https://github.com/zupzup/golang-http-file-upload-download/blob/main/main.go
// NOTE : can also use gin-gonic or gorilla mux to do the job.

// TODO : handle persist code (daku_mantra) before communicating with user

const uploadPath = "./tmp/"
const deletePath = "./bin/"

func main() {
	// NOTE : ensure that ./tmp directory already exists
	newPath := filepath.Join(".", "tmp")
	_ = os.MkdirAll(newPath, os.ModePerm)

	// NOTE : ensure that ./bin directory already exists
	newPath = filepath.Join(".", "bin")
	_ = os.MkdirAll(newPath, os.ModePerm)

	// 1. route to upload file to server
	http.HandleFunc("/upload_file", uploadFileHandler())

	// 2. route to download file from server
	// fs := http.FileServer(http.Dir(uploadPath))
	// http.Handle("/download_file/", http.StripPrefix("/files", fs))

	// 3. route to build file and run stuff
	http.HandleFunc("/run_beast", runBeastHandler())

	// 4. List all files stored in the server
	http.HandleFunc("/list_files", listFilesHandler())

	// 5. Get the size of current directory working files
	http.HandleFunc("/get_size_pwd", getSizePwdHandler())

	// 6. Delete File
	http.HandleFunc("/delete_file", deleteFileHandler())

	// 7. Download File
	http.HandleFunc("/download_file", downloadFileHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func deleteFileHandler() http.HandlerFunc {
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

		if err := r.ParseForm(); err != nil {
			log.Fatalf("Error happened in parsing POST form. Err: %s", err)
		}

		fileName := r.FormValue("file_name")
		filePath := filepath.Join(uploadPath, fileName)

		inputFile, err := os.Open(filePath)

		if err != nil {
			log.Fatalf("Error happened in file opening at original location. Err: %s", err)
		}

		deleteFilePath := filepath.Join(deletePath, fileName)
		deletedFile, err := os.Create(deleteFilePath)

		if err != nil {
			inputFile.Close()
			log.Fatalf("Error happened in file creating at deletion location. Err: %s", err)
		}

		defer deletedFile.Close()
		_, err = io.Copy(deletedFile, inputFile)
		inputFile.Close()

		if err != nil {
			log.Fatalf("Error happened in file content copying. Err: %s", err)
		}

		err = os.Remove(filePath)
		if err != nil {
			log.Fatalf("Error in deleting original file. Err: %s", err)
		}

		// TODO change http status headers in all places
		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "application/json")

		resp := make(map[string]interface{})
		resp["fileToBeDeleted"] = fileName
		resp["status"] = "Deleted"

		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.Write(jsonResp)

		return
	})
}

func downloadFileHandler() http.HandlerFunc {
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

		if err := r.ParseForm(); err != nil {
			log.Fatalf("Error happened in parsing POST form. Err: %s", err)
		}

		fileName := r.FormValue("file_name")
		filePath := filepath.Join(uploadPath, fileName)

		http.ServeFile(w, r, filePath)
	})
}

// https://stackoverflow.com/questions/32482673/how-to-get-directory-total-size
func getSizePwdHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "GET" {
			var size int64
			err := filepath.Walk(uploadPath, func(_ string, info os.FileInfo, err error) error {
				if err != nil {
					log.Fatalf("Error happened in while walking through file directory. Err: %s", err)
				}
				if !info.IsDir() {
					size += info.Size()
				}

				return err
			})

			if err != nil {
				log.Fatalf("Error happened in while calculating file directory size. Err: %s", err)
			}

			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")

			resp := make(map[string]interface{})
			resp["size"] = size

			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}

			w.Write(jsonResp)
			return
		}
	})
}

func listFilesHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "GET" {
			files, err := ioutil.ReadDir(uploadPath)
			if err != nil {
				log.Fatalf("Error happened in while walking through file directory. Err: %s", err)
			}
			// https://stackoverflow.com/questions/14668850/list-directory-in-go

			// create an empty list
			fileNames := []string{}

			for _, f := range files {
				fmt.Println(f.Name())
				fileNames = append(fileNames, f.Name())
			}

			// sort alphabetically
			sort.Strings(fileNames)

			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")

			resp := make(map[string]interface{})
			resp["file_names"] = fileNames

			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}

			w.Write(jsonResp)
			return

		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")

			resp := make(map[string]string)
			resp["message"] = "Only GET requests accepted on this route"

			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}

			w.Write(jsonResp)
			return
		}
	})
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

		var dat map[string]interface{}
		if err := json.Unmarshal(reqBody, &dat); err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		// TODO rather than using log.Fatalf (that crashes the server), just return a JSON error message back to user

		fmt.Println(dat)

		beast_file_contents := dat["script"].(string)

		newPath := filepath.Join(uploadPath, "beast.build")

		if err := ioutil.WriteFile(newPath, []byte(beast_file_contents), 0644); err != nil {
			log.Fatalf("Error happened while writing to `beast.build` file. Err: %s", err)
		}

		w.WriteHeader(http.StatusAccepted)
		w.Header().Set("Content-Type", "application/json")

		// output, err := exec.Command("g++", "-Wall", newPath).Output()
		// output, err := exec.Command("pwd").Output()
		cmd := exec.Command("beast")
		cmd.Dir = uploadPath
		output, err := cmd.Output()
		fmt.Println(string(output))

		// if err != nil {
		// 	log.Fatalf("Error happened when executing command. Err: %s", err)
		// }

		// Note : this is how to first create a map and then add elements to it
		// this is the correct way to do stuff
		response := make(map[string]interface{})
		outputList := strings.Split(string(output), "\n")
		if len(outputList) > 0 {
			outputList = outputList[:len(outputList)-1]
		}
		response["output"] = outputList

		jsonResp, err := json.Marshal(response)

		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}

		w.Write(jsonResp)
		return

	})
}
