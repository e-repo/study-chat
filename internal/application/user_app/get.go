package user_app

//nolint:stylecheck // fit to generated code
//func (h UserEndpoints) GetUsersId(c echo.Context, Id openapi_types.UUID) error {
//	user, err := h.repo.GetUserById(c.Request().Context(), Id)
//	if err != nil {
//		msg := err.Error()
//		if errors.Is(err, user.ErrUserNotFound) {
//			return c.JSON(http.StatusNotFound, openapi.ErrorResponse{Message: &msg})
//		}
//		return c.JSON(http.StatusInternalServerError, openapi.ErrorResponse{Message: &msg})
//	}
//
//	name := user.FirstName()
//	Email := openapi_types.Email(user.Email())
//	return c.JSON(http.StatusOK, openapi.GetUserResponse{
//		Id:    &Id,
//		Name:  &name,
//		Email: &Email,
//	})
//}
//
//func (h UserEndpoints) GetUserById(ctx context.Context, req *protobuf.GetUserRequest) (*protobuf.GetUserResponse, error) {
//	uid, err := uuid.Parse(req.GetId())
//	if err != nil {
//		return nil, status.Error(codes.InvalidArgument, "invalid UUID")
//	}
//	user, err := h.repo.GetUserById(ctx, uid)
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
