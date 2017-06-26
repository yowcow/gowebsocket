package handler

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHtml(t *testing.T) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)
	Html(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "text/html;charset=utf-8", w.Header().Get("content-type"))

	re := regexp.MustCompile("^<!doctype html>")

	assert.True(t, re.Match(w.Body.Bytes()))
}
