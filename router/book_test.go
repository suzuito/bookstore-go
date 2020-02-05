package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func toJSON(v interface{}) string {
	ret, _ := json.Marshal(v)
	return string(ret)
}

func Test(t *testing.T) {
	testCases := []struct {
		desc           string
		inputMethod    string
		inputPath      string
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			desc:           "",
			inputPath:      "/status",
			expectedStatus: http.StatusOK,
			expectedBody:   &responseStatus{Message: "ok"},
		},
		{
			desc:           "",
			inputPath:      "/books/123",
			expectedStatus: http.StatusOK,
			expectedBody:   &responseStatus{Message: "ok"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			root := gin.Default()
			// Setup router
			InitializeRouter(root, nil)
			// Mock request
			req, _ := http.NewRequest(tC.inputMethod, tC.inputPath, nil)
			rec := httptest.NewRecorder()
			// Run
			root.ServeHTTP(rec, req)
			expectedBody, _ := json.Marshal(tC.expectedBody)
			// Assertion
			assert.Equal(t, tC.expectedStatus, rec.Result().StatusCode)
			assert.JSONEq(t, string(expectedBody), rec.Body.String())
		})
	}
}
