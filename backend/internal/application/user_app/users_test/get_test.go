package users_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"study-chat/generated/openapi"
	"study-chat/generated/protobuf"
	"study-chat/internal/domain/user_dmn"

	"github.com/google/uuid"
)

func (s *UsersSuite) TestGetUser() {
	user, _ := user_dmn.CreateUser("test", "test@test.com")
	err := s.UsersRepo.SaveUser(context.Background(), *user)
	s.Require().NoError(err)

	s.Run("HTTP", func() {
		req := httptest.NewRequest(http.MethodGet, "/user/"+user.ID().String(), nil)
		rec := httptest.NewRecorder()
		s.HTTPServer.ServeHTTP(rec, req)

		s.Require().Equal(http.StatusOK, rec.Code, rec.Body.String())
		var resp openapi.GetUserResponse
		err = json.Unmarshal(rec.Body.Bytes(), &resp)
		s.Require().NoError(err)
		s.Require().Equal(user.ID(), *resp.Id)
		s.Require().Equal(user.FirstName(), *resp.Name)
		s.Require().Equal(user.Email(), string(*resp.Email))
	})

	s.Run("GRPC", func() {
		req := &protobuf.GetUserRequest{Id: user.ID().String()}
		resp, err := s.GRPCHandlers.GetUser(context.Background(), req)

		s.Require().NoError(err)
		s.Require().Equal(user.ID().String(), resp.GetId())
		s.Require().Equal(user.FirstName(), resp.GetName())
		s.Require().Equal(user.Email(), resp.GetEmail())
	})
}

func (s *UsersSuite) TestGetUserNotFound() {
	s.Run("HTTP", func() {
		req := httptest.NewRequest(http.MethodGet, "/user/"+uuid.New().String(), nil)
		rec := httptest.NewRecorder()
		s.HTTPServer.ServeHTTP(rec, req)

		s.Require().Equal(http.StatusNotFound, rec.Code, rec.Body.String())
		s.Require().Equal(`{"message":"users not found"}`+"\n", rec.Body.String())
	})

	s.Run("GRPC", func() {
		req := &protobuf.GetUserRequest{Id: uuid.New().String()}
		_, err := s.GRPCHandlers.GetUser(context.Background(), req)

		s.Require().Error(err)
		s.Require().Equal("rpc error: code = NotFound desc = users not found", err.Error())
	})
}
