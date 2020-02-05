package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/suzuito/bookstore-go/entity"
)

func toJSON(v interface{}) string {
	ret, _ := json.Marshal(v)
	return string(ret)
}

func Test(t *testing.T) {
	testCases := []struct {
		desc           string
		inputMethod    string
		setupMock      func(*MockRepository, *MockApplication)
		inputPath      string
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			desc:      "Success",
			inputPath: "/books/123",
			setupMock: func(repository *MockRepository, application *MockApplication) {
				repository.
					EXPECT().
					GetBookByID(gomock.Eq("123"), gomock.Any()).
					SetArg(1, entity.Book{ID: "123", Name: "hoge", ISBN: "fuga", Price: 100.0}).
					Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   &responseBook{Name: "hoge", ISBN: "fuga", Price: 100.0},
		},
		{
			desc:      "Repository return not found error",
			inputPath: "/books/123",
			setupMock: func(repository *MockRepository, application *MockApplication) {
				repository.
					EXPECT().
					GetBookByID(gomock.Eq("123"), gomock.Any()).
					Return(NewWrappedError(NotFoundError, "dummy error"))
			},
			expectedStatus: http.StatusNotFound,
			expectedBody:   &responseError{Message: "dummy error"},
		},
		{
			desc:      "Repository return invalid parameters error",
			inputPath: "/books/123",
			setupMock: func(repository *MockRepository, application *MockApplication) {
				repository.
					EXPECT().
					GetBookByID(gomock.Eq("123"), gomock.Any()).
					Return(NewWrappedError(InvalidParameterError, "dummy error"))
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   &responseError{Message: "dummy error"},
		},
		{
			desc:      "Repository return unknown error",
			inputPath: "/books/123",
			setupMock: func(repository *MockRepository, application *MockApplication) {
				repository.
					EXPECT().
					GetBookByID(gomock.Eq("123"), gomock.Any()).
					Return(NewWrappedError(fmt.Errorf("unknown error"), "dummy error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   &responseError{Message: "dummy error"},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// New Mock interface
			ctrlRepository := gomock.NewController(t)
			defer ctrlRepository.Finish()
			mockRepository := NewMockRepository(ctrlRepository)
			ctrlApplication := gomock.NewController(t)
			defer ctrlApplication.Finish()
			mockApplication := NewMockApplication(ctrlApplication)
			mockApplication.
				EXPECT().
				NewRepository(gomock.Any()).
				Return(mockRepository, nil)
			tC.setupMock(mockRepository, mockApplication)
			// Setup router
			root := gin.Default()
			InitializeRouter(root, mockApplication)
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
