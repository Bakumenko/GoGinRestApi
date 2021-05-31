package handler

import (
	"apiserver/pkg/model"
	"apiserver/pkg/service"
	mock_service "apiserver/pkg/service/mocks"
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_createUser(t *testing.T) {
	type mockBehavior func(s *mock_service.MockUser, user model.User, roleName string)

	id := uuid.New()
	testTable := []struct {
		name                string
		inputBody           string
		inputUser           model.User
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "check data is true",
			inputBody: `{"firstname": "TestFN", "lastname": "TestLN", "email": "test@test.com", "age": 25}`,
			inputUser: model.User{
				Firstname: "TestFN",
				Lastname:  "TestLN",
				Email:     "test@test.com",
				Age:       25,
			},
			mockBehavior: func(s *mock_service.MockUser, user model.User, roleName string) {
				s.EXPECT().CreateUser(user, roleName).Return(id, nil)
			},
			expectedStatusCode:  200,
			expectedRequestBody: "{\"id\":\"" + id.String() + "\"}",
		},
		{
			name:      "check email for free",
			inputBody: `{"firstname": "TestFN", "lastname": "TestLN", "email": "test@test.com", "age": 25}`,
			inputUser: model.User{
				Firstname: "TestFN",
				Lastname:  "TestLN",
				Email:     "test@test.com",
				Age:       25,
			},
			mockBehavior: func(s *mock_service.MockUser, user model.User, roleName string) {
				s.EXPECT().CreateUser(user, roleName).Return(uuid.Nil, errors.New("email is already contains"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: "{\"message\":\"email is already contains\"}",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			//Init deps
			c := gomock.NewController(t)
			defer c.Finish()

			createdUser := mock_service.NewMockUser(c)
			testCase.mockBehavior(createdUser, testCase.inputUser, "ROLE_USER")

			services := &service.Service{User: createdUser}
			handler := NewHandler(services)

			//Test server
			r := gin.New()
			r.POST("", handler.createUser)

			//Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/", bytes.NewBufferString(testCase.inputBody))

			//Perform request
			r.ServeHTTP(w, req)

			//Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
