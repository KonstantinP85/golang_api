package handler

import (
	"fmt"
	"github.com/KonstantinP85/api-go"
	"github.com/KonstantinP85/api-go/pkg/service"
	"github.com/KonstantinP85/api-go/pkg/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http/httptest"
	"testing"
)

func TestHandler_getUser(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUser, id int)
	var userReturn api_go.User

	testTable := []struct {
		name 				string
		id 			        int
		inputUser 			api_go.User
		mockBehavior 		mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	} {
		{
			name: "OK",
			id: 1,
			inputUser: api_go.User{
				Email: "test@test.ru",
				Password: "test",
			},
			mockBehavior: func (s *mock_service.MockUser, id int){
				s.EXPECT().GetUser(id).Return(userReturn, nil)
			},
			expectedStatusCode: 200,
			expectedRequestBody: `{"email":"test@test.ru","id":1}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T){
			c := gomock.NewController(t)
			defer c.Finish()

			user := mock_service.NewMockUser(c)
			testCase.mockBehavior(user, testCase.id)

			services := &service.Service{User: user}
			handler := NewHandler(services)

			path := fmt.Sprintf("/admin/api/user/%d", testCase.id)

			r := gin.New()
			r.GET("/admin/api/user/:id", handler.getUser)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", path, nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}