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

	handler := http.HandlerFunc(Html)
	handler.ServeHTTP(w, req)

	var re *regexp.Regexp

	re = regexp.MustCompile("^text/html;")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.True(t, re.MatchString(w.Header().Get("content-type")))

	re = regexp.MustCompile("^<!doctype html>")
	assert.True(t, re.Match(w.Body.Bytes()))
}
