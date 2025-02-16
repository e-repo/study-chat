package application

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type ServiceLocator struct {
	suite.Suite
	HTTPServer *echo.Echo
	//UsersRepo    *usersInfra.InMemoryRepo
	//ProductsRepo *productsInfra.InMemoryRepo
}

//func (s *ServiceLocator) SetupTest() {
//	s.UsersRepo = usersInfra.NewInMemoryRepo()
//	s.ProductsRepo = productsInfra.NewInMemoryRepo()
//	s.HTTPServer = SetupHTTPServer(s.UsersRepo, s.ProductsRepo)
//}
