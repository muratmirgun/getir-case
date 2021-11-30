package datastore

import (
	"bytes"
	"getir-case/internal/store/inmemory"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	wr := httptest.NewRecorder()

	holder := inmemory.New()

	err := holder.Set("hi", "hello")
	if err != nil {
		return
	}
	dataHandler := New(holder)

	req := httptest.NewRequest(http.MethodGet, "/holder/get", nil)

	req.Header.Set("key", "hi")

	dataHandler.GetInMemory(wr, req)

	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}

	if !strings.Contains(wr.Body.String(), `{"key":"hi","value":"hello"}`) {
		t.Errorf(
			`response body "%s" does not contain {"key":"hi","value":"hello"}`,
			wr.Body.String(),
		)
	}
}

func TestSet(t *testing.T) {
	wr := httptest.NewRecorder()

	holder := inmemory.New()

	dataHandler := New(holder)

	jsonStr := []byte(`{"key":"hi","value":"hello"}`)

	req := httptest.NewRequest(http.MethodPost, "/holder/set", bytes.NewBuffer(jsonStr))

	dataHandler.SetInMemory(wr, req)
	if wr.Code != http.StatusCreated {
		t.Errorf("got HTTP status code %d, expected 201", wr.Code)
	}

	if !strings.Contains(wr.Body.String(), `{"key":"hi","value":"hello"}`) {
		t.Errorf(
			`response body "%s" does not contain {"key":"hi","value":"hello"}`,
			wr.Body.String(),
		)
	}
}
