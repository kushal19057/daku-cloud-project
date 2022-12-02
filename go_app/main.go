package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
)

// refer: https://github.com/zupzup/golang-http-file-upload-download/blob/main/main.go
// NOTE : can also use gin-gonic or gorilla mux to do the job.

// TODO : handle persist code (daku_mantra) before communicating with user

const UPLOAD_PATH = "./tmp/"
const UPLOAD_FOLDER = "tmp"
const BIN_PATH = "./bin/"
const BIN_FOLDER = "bin"

func main() {
	// NOTE : ensure that ./tmp directory already exists
	_ = os.MkdirAll(filepath.Join(".", UPLOAD_FOLDER), os.ModePerm)

	// NOTE : ensure that ./bin directory already exists
	_ = os.MkdirAll(filepath.Join(".", BIN_FOLDER), os.ModePerm)

	// 1. route to upload file to server
	http.HandleFunc("/upload", uploadFileHandler())
	// 2. List all files stored in the server
	http.HandleFunc("/files", listWorkFilesHandler())
	// 3. Get the size of current directory working files
	http.HandleFunc("/size", getWorkDirSizeHandler())

	// // 2. route to build file and run stuff
	// http.HandleFunc("/beast", runBeastHandler())

	// 2. List all files stored in the bin
	// http.HandleFunc("/bin", listBinFilesHandler())

	// // 6. Delete File
	// http.HandleFunc("/delete_file", deleteFileHandler())

	// // 7. Download File
	// http.HandleFunc("/download_file", downloadFileHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// 1. Handle uploading of files
func uploadFileHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if r.Method == "POST" {
			file, fileHeader, err := r.FormFile("upload_file")
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println("Error while reading file from http request")
				return
			}

			defer file.Close()
			file_bytes, err := ioutil.ReadAll(file)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Println("Cannot read file contents to a byte array")
				return
			}

			file_upload_path := filepath.Join(UPLOAD_PATH, fileHeader.Filename)
			uploaded_file, err := os.Create(file_upload_path)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Println("Cannot create new file on disk")
				return
			}

			defer uploaded_file.Close()

			if _, err := uploaded_file.Write(file_bytes); err != nil || uploaded_file.Close() != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Println("Cannot write contents of file to disk")
				return
			}

			w.WriteHeader(http.StatusAccepted)
			fmt.Println("File created and written to disk successfully")
			return
		}

		// handle non GET requests to this route
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Non POST request")
		return
	})
}

// 2. Handle listing of files
func listWorkFilesHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "GET" {
			files, err := ioutil.ReadDir(UPLOAD_PATH)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Println("Error walking thru file directory")
				return
			}

			// https://stackoverflow.com/questions/14668850/list-directory-in-go
			// create an empty list
			file_names := []string{}

			for _, f := range files {
				file_names = append(file_names, f.Name())
			}

			// sort alphabetically
			sort.Strings(file_names)

			// reply back to client with files
			w.WriteHeader(http.StatusAccepted)
			w.Header().Set("Content-Type", "application/json")

			resp := make(map[string]interface{})
			resp["files"] = file_names

			json_resp, _ := json.Marshal(resp)
			w.Write(json_resp)
			return
		}

		// handle non GET requests
		w.WriteHeader(http.StatusBadRequest)
		return
	})
}

/*
// func deleteFileHandler() http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		enableCors(&w)

// 		if r.Method == "GET" {
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Header().Set("Content-Type", "application/json")

// 			resp := make(map[string]string)
// 			resp["message"] = "The route does not accept GET requests. Try POST request."

// 			jsonResp, err := json.Marshal(resp)
// 			if err != nil {
// 				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
// 			}

// 			w.Write(jsonResp)
// 			return
// 		}

// 		if err := r.ParseForm(); err != nil {
// 			log.Fatalf("Error happened in parsing POST form. Err: %s", err)
// 		}

// 		fileName := r.FormValue("file_name")
// 		filePath := filepath.Join(uploadPath, fileName)

// 		inputFile, err := os.Open(filePath)

// 		if err != nil {
// 			log.Fatalf("Error happened in file opening at original location. Err: %s", err)
// 		}

// 		deleteFilePath := filepath.Join(deletePath, fileName)
// 		deletedFile, err := os.Create(deleteFilePath)

// 		if err != nil {
// 			inputFile.Close()
// 			log.Fatalf("Error happened in file creating at deletion location. Err: %s", err)
// 		}

// 		defer deletedFile.Close()
// 		_, err = io.Copy(deletedFile, inputFile)
// 		inputFile.Close()

// 		if err != nil {
// 			log.Fatalf("Error happened in file content copying. Err: %s", err)
// 		}

// 		err = os.Remove(filePath)
// 		if err != nil {
// 			log.Fatalf("Error in deleting original file. Err: %s", err)
// 		}

// 		// TODO change http status headers in all places
// 		w.WriteHeader(http.StatusAccepted)
// 		w.Header().Set("Content-Type", "application/json")

// 		resp := make(map[string]interface{})
// 		resp["fileToBeDeleted"] = fileName
// 		resp["status"] = "Deleted"

// 		jsonResp, err := json.Marshal(resp)
// 		if err != nil {
// 			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
// 		}

// 		w.Write(jsonResp)

// 		return
// 	})
// }

// func downloadFileHandler() http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		enableCors(&w)
// 		if r.Method == "GET" {
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Header().Set("Content-Type", "application/json")

// 			resp := make(map[string]string)
// 			resp["message"] = "The route does not accept GET requests. Try POST request."

// 			jsonResp, err := json.Marshal(resp)
// 			if err != nil {
// 				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
// 			}

// 			w.Write(jsonResp)
// 			return
// 		}

// 		if err := r.ParseForm(); err != nil {
// 			log.Fatalf("Error happened in parsing POST form. Err: %s", err)
// 		}

// 		fileName := r.FormValue("file_name")
// 		filePath := filepath.Join(uploadPath, fileName)

// 		http.ServeFile(w, r, filePath)
// 	})
// }
*/

// https://stackoverflow.com/questions/32482673/how-to-get-directory-total-size
// 3. Handle calculating the size of all files in the /tmp folder
func getWorkDirSizeHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "GET" {
			var size int64
			err := filepath.Walk(UPLOAD_PATH, func(_ string, info os.FileInfo, err error) error {
				if err != nil {
					fmt.Printf("Error happened in while walking through file directory. Err: %s\n", err)
				}

				if !info.IsDir() {
					size += info.Size()
				}
				return err
			})

			if err != nil {
				fmt.Printf("Error happened in while calculating file directory size. Err: %s\n", err)
			}

			w.WriteHeader(http.StatusAccepted)
			w.Header().Set("Content-Type", "application/json")

			resp := make(map[string]interface{})
			resp["size"] = size
			json_resp, _ := json.Marshal(resp)
			w.Write(json_resp)

			return
		}

		// handle non GET requests
		w.WriteHeader(http.StatusBadRequest)
		return
	})
}

// func runBeastHandler() http.HandlerFunc {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		enableCors(&w)

// 		if r.Method == "GET" {
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Header().Set("Content-Type", "application/json")

// 			resp := make(map[string]string)
// 			resp["message"] = "The route does not accept GET requests. Try POST request."

// 			jsonResp, err := json.Marshal(resp)
// 			if err != nil {
// 				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
// 			}

// 			w.Write(jsonResp)
// 			return
// 		}

// 		reqBody, err := ioutil.ReadAll(r.Body)

// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Header().Set("Content-Type", "application/json")

// 			resp := make(map[string]string)
// 			resp["message"] = "Invalid Request"

// 			jsonResp, err := json.Marshal(resp)
// 			if err != nil {
// 				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
// 			}

// 			w.Write(jsonResp)
// 			return
// 		}

// 		var dat map[string]interface{}
// 		if err := json.Unmarshal(reqBody, &dat); err != nil {
// 			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
// 		}

// 		// TODO rather than using log.Fatalf (that crashes the server), just return a JSON error message back to user

// 		fmt.Println(dat)

// 		beast_file_contents := dat["script"].(string)

// 		newPath := filepath.Join(uploadPath, "beast.build")

// 		if err := ioutil.WriteFile(newPath, []byte(beast_file_contents), 0644); err != nil {
// 			log.Fatalf("Error happened while writing to `beast.build` file. Err: %s", err)
// 		}

// 		w.WriteHeader(http.StatusAccepted)
// 		w.Header().Set("Content-Type", "application/json")

// 		// output, err := exec.Command("g++", "-Wall", newPath).Output()
// 		// output, err := exec.Command("pwd").Output()
// 		cmd := exec.Command("beast")
// 		cmd.Dir = uploadPath
// 		output, err := cmd.Output()
// 		fmt.Println(string(output))

// 		// if err != nil {
// 		// 	log.Fatalf("Error happened when executing command. Err: %s", err)
// 		// }

// 		// Note : this is how to first create a map and then add elements to it
// 		// this is the correct way to do stuff
// 		response := make(map[string]interface{})
// 		outputList := strings.Split(string(output), "\n")
// 		if len(outputList) > 0 {
// 			outputList = outputList[:len(outputList)-1]
// 		}
// 		response["output"] = outputList

// 		jsonResp, err := json.Marshal(response)

// 		if err != nil {
// 			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
// 		}

// 		w.Write(jsonResp)
// 		return

// 	})
// }
