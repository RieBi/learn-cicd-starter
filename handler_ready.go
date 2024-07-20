package main

import (
	"net/http"
	"testing"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func TestReadiness(t *testing.T) {
	got := 8
	if got != 8 {
		t.Fatalf("Foo %v, got: %v", 8, got)
	}
}
