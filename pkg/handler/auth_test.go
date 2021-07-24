package handler

import (
	"bytes"
	"github.com/KonstantinP85/api-go"
	"github.com/KonstantinP85/api-go/pkg/service"
	mock_service "github.com/KonstantinP85/api-go/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_signUp(t *testing.T) {
	type mockBehavior func(s *mock_service.MockAuthorization, user api_go.User)

	testTable := []struct {
		name 				string
		inputBody 			string
		inputUser 			api_go.User
		mockBehavior 		mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	} {
		{
			name: "OK",
			inputBody: `{"email":"test@test.ru", "password":"test"}`,
			inputUser: api_go.User{
				Email: "test@test.ru",
				Password: "test",
			},
			mockBehavior: func(s *mock_service.MockAuthorization, user api_go.User) {
				s.EXPECT().CreateUser(user).Return(1, "test@test.ru", nil)
			},
			expectedStatusCode: 200,
			expectedRequestBody: `{"email":"test@test.ru","id":1}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T){
			c := gomock.NewController(t)
			defer c.Finish()

			auth := mock_service.NewMockAuthorization(c)
			testCase.mockBehavior(auth, testCase.inputUser)

			services := &service.Service{Authorization: auth}
			handler := NewHandler(services)

			r := gin.New()
			r.POST("/sign-up", handler.signUp)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up", bytes.NewBufferString(testCase.inputBody))

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
