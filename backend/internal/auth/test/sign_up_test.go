package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"study-chat/generated/openapi"
	"study-chat/internal/auth"
)

const signUpUrl = "/sign-up"

func (s *UserSuite) TestSuccessSignUp() {
	userResponse := openapi.CreateUserResponse{}
	userRequest := openapi.CreateUserRequest{
		Email:     "test@test.ru",
		FirstName: "Иван",
		Password:  "secret",
	}
	requestBody, err := json.Marshal(userRequest)
	s.Require().NoError(err)

	s.Run("SuccessSignUp", func() {
		request := httptest.NewRequest(
			http.MethodPost,
			signUpUrl,
			bytes.NewBuffer(requestBody),
		)

		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		s.Echo.ServeHTTP(response, request)

		s.Assert().Equal(http.StatusCreated, response.Code)

		err := json.Unmarshal(response.Body.Bytes(), &userResponse)
		s.Require().NoError(err)
		s.Assert().NotEmpty(userResponse.Id)
	})
}

func (s *UserSuite) TestFailedByInvalidEmail() {
	userResponse := openapi.ErrorResponse{}
	userRequest := auth.RequestSignUp{
		Email:     "testtest.ru",
		FirstName: "Иван",
		Password:  "secret",
	}
	requestBody, err := json.Marshal(userRequest)
	s.Require().NoError(err)

	s.Run("FailedSignUpByInvalidEmail", func() {
		request := httptest.NewRequest(
			http.MethodPost,
			signUpUrl,
			bytes.NewBuffer(requestBody),
		)

		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		s.Echo.ServeHTTP(response, request)

		s.Assert().Equal(http.StatusBadRequest, response.Code)

		err := json.Unmarshal(response.Body.Bytes(), &userResponse)
		s.Require().NoError(err)
		s.Assert().Equal("Email должен быть email адресом", *userResponse.Message)
	})
}
