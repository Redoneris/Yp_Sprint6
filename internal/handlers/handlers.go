package handlers

import (
	"Yp_Sprint6/internal/service"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func HandleIndexHTMLFile(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusInternalServerError)
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	file, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	w.Write(file)

}

func HandleUpload(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusInternalServerError)
		return
	}

	err := r.ParseMultipartForm(10 << 20) //10MB
	if err != nil {
		http.Error(w, "Parse error", http.StatusInternalServerError)
	}

	file, handler, err := r.FormFile("index.html")
	if err != nil {
		http.Error(w, "Cannot get file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Cannot read file", http.StatusInternalServerError)
		return
	}

	fileContent := string(data)

	convertedText, err := service.AutoDetect(fileContent)
	if err != nil {
		http.Error(w, "Conversion error", http.StatusInternalServerError)
		return
	}

	timestamp := time.Now().UTC().Format("2006-01-02_15-04-05")
	originalExt := filepath.Ext(handler.Filename)

	resultFilename := fmt.Sprintf("result_%s%s", timestamp, originalExt)

	resultFile, err := os.Create(resultFilename)
	if err != nil {
		http.Error(w, "Cannot create a result file", http.StatusInternalServerError)
		return
	}
	defer resultFile.Close()

	_, err = resultFile.WriteString(convertedText)
	if err != nil {
		http.Error(w, "Не могу записать в файл: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}
