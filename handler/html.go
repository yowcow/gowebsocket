package handler

import (
	"bufio"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func Html(w http.ResponseWriter, r *http.Request) {
	_, file, _, _ := runtime.Caller(0)
	htmlfile := filepath.Join(filepath.Dir(file), "..", "public", "index.html")

	w.Header().Add("content-type", "text/html;charset=utf-8")
	w.WriteHeader(http.StatusOK)

	f, _ := os.Open(htmlfile)
	reader := bufio.NewReader(f)
	for {
		b, e := reader.ReadBytes('\n')
		if e != nil {
			break
		}
		w.Write(b)
	}

	f.Close()
}
