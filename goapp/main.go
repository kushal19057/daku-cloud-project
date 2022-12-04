package main

import (
	"encoding/json"
	"errors"
	"fmt"
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

const UPLOAD_PATH = "./tmp/"
const UPLOAD_FOLDER = "tmp"
const BIN_PATH = "./bin/"
const BIN_FOLDER = "bin"
const BEAST_FILE_NAME = "beast.build"

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
	// 4. route to build file and run stuff
	http.HandleFunc("/beast", runBeastHandler())
	// 5. Delete File
	http.HandleFunc("/delete", deleteFileHandler())

	// permanently Delete File
	http.HandleFunc("/permanent_delete", permanentDeleteFileHandler())

	// restore deleted File
	http.HandleFunc("/restore", restoreFileHandler())

	// 6. Download File
	http.HandleFunc("/download", downloadFileHandler())

	// 2. List all files stored in the bin
	http.HandleFunc("/bin_files", listBinFilesHandler())

	//

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

func listBinFilesHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "GET" {
			files, err := ioutil.ReadDir(BIN_PATH)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Println("Error walking thru file directory")
				return
			}

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

// https://stackoverflow.com/questions/50740902/move-a-file-to-a-different-drive-with-go
// https://golangbyexample.com/move-file-from-one-location-to-another-golang/
// 5. Handle file deletion (move from /tmp to /bin)
func deleteFileHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "POST" {
			reqBody, err := ioutil.ReadAll(r.Body)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println("Cannot reqd request body")
				fmt.Printf("%s\n", err.Error())
				return
			}

			var data map[string]interface{}
			if err := json.Unmarshal(reqBody, &data); err != nil {
				fmt.Println("Cannot unmarshal request body to json")
				fmt.Printf("%s\n", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			file_name := data["file"].(string)
			file_path := filepath.Join(UPLOAD_PATH, file_name)
			bin_file_path := filepath.Join(BIN_PATH, file_name)

			err = os.Chmod(file_path, 0777)
			if err != nil {
				fmt.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = os.Rename(file_path, bin_file_path)
			if err != nil {
				fmt.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusAccepted)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Only POST requests accepted on the /delete route")
		return
	})
}

func restoreFileHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "POST" {
			reqBody, err := ioutil.ReadAll(r.Body)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println("Cannot reqd request body")
				fmt.Printf("%s\n", err.Error())
				return
			}

			var data map[string]interface{}
			if err := json.Unmarshal(reqBody, &data); err != nil {
				fmt.Println("Cannot unmarshal request body to json")
				fmt.Printf("%s\n", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			file_name := data["file"].(string)
			file_path := filepath.Join(BIN_PATH, file_name)
			bin_file_path := filepath.Join(UPLOAD_PATH, file_name)

			err = os.Chmod(file_path, 0777)
			if err != nil {
				fmt.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			err = os.Rename(file_path, bin_file_path)
			if err != nil {
				fmt.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusAccepted)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Only POST requests accepted on the /delete route")
		return
	})
}

func permanentDeleteFileHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "POST" {
			reqBody, err := ioutil.ReadAll(r.Body)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println("Cannot reqd request body")
				fmt.Printf("%s\n", err.Error())
				return
			}

			var data map[string]interface{}
			if err := json.Unmarshal(reqBody, &data); err != nil {
				fmt.Println("Cannot unmarshal request body to json")
				fmt.Printf("%s\n", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			file_name := data["file"].(string)
			file_path := filepath.Join(BIN_PATH, file_name)
			// bin_file_path := filepath.Join(BIN_PATH, file_name)

			err = os.Remove(file_path)
			if err != nil {
				fmt.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusAccepted)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Only POST requests accepted on the /permanent_delete route")
		return
	})
}

// 6. Download file handler
func downloadFileHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if r.Method == "POST" {
			reqBody, err := ioutil.ReadAll(r.Body)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println("Cannot reqd request body")
				fmt.Printf("%s\n", err.Error())
				return
			}

			var data map[string]interface{}
			if err := json.Unmarshal(reqBody, &data); err != nil {
				fmt.Println("Cannot unmarshal request body to json")
				fmt.Printf("%s\n", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			file_name := data["file"].(string)
			file_path := filepath.Join(UPLOAD_PATH, file_name)

			// https://stackoverflow.com/a/12518877
			if _, err := os.Stat(file_path); err == nil {
				// path/to/whatever exists
				http.ServeFile(w, r, file_path)
				return

			} else if errors.Is(err, os.ErrNotExist) {
				// path/to/whatever does *not* exist
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println("File not found on server")
				return
			} else {
				// Schrodinger: file may or may not exist. See err for details.

				// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println("File not found on server | schrodinger")
				return
			}

		}

		// handle non POST requests
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("/download route only handle POST requests")
		return

	})
}

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

// 4. beast file handler
func runBeastHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		if r.Method == "POST" {
			reqBody, err := ioutil.ReadAll(r.Body)

			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Println("Cannot reqd request body")
				fmt.Printf("%s\n", err.Error())
				return
			}

			// fmt.Println(string(reqBody))
			var data map[string]interface{}
			if err := json.Unmarshal(reqBody, &data); err != nil {
				fmt.Println("Cannot unmarshal request body to json")
				fmt.Printf("%s\n", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			beast_file_contents := data["script"].(string)
			beast_file_path := filepath.Join(UPLOAD_PATH, BEAST_FILE_NAME)

			if err := ioutil.WriteFile(beast_file_path, []byte(beast_file_contents), 0644); err != nil {
				fmt.Println("Cannot write contents to beast.build file")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			cmd := exec.Command("beast")
			cmd.Dir = UPLOAD_PATH // change directory
			output, err := cmd.Output()

			response := make(map[string]interface{})
			outputList := strings.Split(string(output), "\n")
			if len(outputList) > 0 {
				outputList = outputList[:len(outputList)-1]
			}
			response["output"] = outputList

			jsonResp, err := json.Marshal(response)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Printf("%s\n", err.Error())
				return
			}

			w.WriteHeader(http.StatusAccepted)
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResp)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		fmt.Println("Only POST requests handled at /beast route")
		return
	})
}
