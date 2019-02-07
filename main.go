package main

// https://www.flaticon.com/free-icon/dog_1402202
import (
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	minio "github.com/minio/minio-go"
)

var store = NewStore()

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{reponame}/{path:.*}", TestHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	srv.ListenAndServe()
}

func GetArtifact(path string, writer http.ResponseWriter) {
	object, err := store.Client.GetObject("mymusic", path, minio.GetObjectOptions{})
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	defer object.Close()

	objInfo, err := object.Stat()
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.Header().Add("Content-Type", "application/octet-stream")
	writer.Header().Add("Content-Length", strconv.FormatInt(objInfo.Size, 10))
	writer.WriteHeader(http.StatusOK)

	io.Copy(writer, object)
}

func UploadArtifact(path string, r *http.Request, writer http.ResponseWriter) {
	defer r.Body.Close()
	_, err := store.Client.PutObject("mymusic", path, r.Body, r.ContentLength, minio.PutObjectOptions{})
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if r.Method == "GET" {
		GetArtifact(vars["path"], w)
	} else if r.Method == "PUT" || r.Method == "POST" {
		log.Printf("request: %v", r)
		UploadArtifact(vars["path"], r, w)
	}

	log.Printf("Method: %s", r.Method)
}
