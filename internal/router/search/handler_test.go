package search

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	wr := httptest.NewRecorder()

	jsonStr := []byte(`{"startDate":"2016-01-21","endDate":"2016-01-30","minCount":2900,"maxCount":3000}`)

	out := `{"code":0,"msg":"Success","records":[{"createdAt":"2016-01-29T03:59:53.494+02:00","key":"bxoQiSKL","totalCount":2991},{"createdAt":"2016-01-21T15:50:50.679+02:00","key":"lGUyjflG","totalCount":2940}]}`
	req := httptest.NewRequest(http.MethodPost, "/search", bytes.NewBuffer(jsonStr))

	mongoServer := new(MongoDB)
	mongoServer.ServeMongo(wr, req)

	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}

	if !strings.Contains(wr.Body.String(), out) {
		t.Errorf(
			`response body "%s" does not contain `+out,
			wr.Body.String(),
		)
	}
}
