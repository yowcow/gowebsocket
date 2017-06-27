package handler

import (
	"net/http"
	"path/filepath"
	"runtime"
)

func Html(w http.ResponseWriter, r *http.Request) {
	_, file, _, _ := runtime.Caller(0)
	htmlfile := filepath.Join(filepath.Dir(file), "..", "public", "index.html")
	http.ServeFile(w, r, htmlfile)
}
