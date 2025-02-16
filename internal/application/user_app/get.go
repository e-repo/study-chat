package user_app

//nolint:stylecheck // fit to generated code
//func (h UserServer) GetUsersId(c echo.Context, id openapi_types.UUID) error {
//	user, err := h.repo.GetUser(c.Request().Context(), id)
//	if err != nil {
//		msg := err.Error()
//		if errors.Is(err, user.ErrUserNotFound) {
//			return c.JSON(http.StatusNotFound, openapi.ErrorResponse{Message: &msg})
//		}
//		return c.JSON(http.StatusInternalServerError, openapi.ErrorResponse{Message: &msg})
//	}
//
//	name := user.FirstName()
//	email := openapi_types.Email(user.Email())
//	return c.JSON(http.StatusOK, openapi.GetUserResponse{
//		Id:    &id,
//		Name:  &name,
//		Email: &email,
//	})
//}
//
//func (h UserServer) GetUser(ctx context.Context, req *protobuf.GetUserRequest) (*protobuf.GetUserResponse, error) {
//	uid, err := uuid.Parse(req.GetId())
//	if err != nil {
//		return nil, status.Error(codes.InvalidArgument, "invalid UUID")
//	}
//	user, err := h.repo.GetUser(ctx, uid)
//	if err != nil {
//		if errors.Is(err, user.ErrUserNotFound) {
//			return nil, status.Error(codes.NotFound, "users not found")
//		}
//		return nil, err
//	}
//
//	return &protobuf.GetUserResponse{
//		Id:    user.ID().String(),
//		Name:  user.FirstName(),
//		Email: user.Email(),
//	}, nil
//}
