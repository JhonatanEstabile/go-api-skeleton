//package user
//
//import (
//	mocks2 "api/framework/mocks"
//	"database/sql"
//	"errors"
//	"fmt"
//	"github.com/gofiber/fiber/v2"
//	"net/http/httptest"
//	"strings"
//	"testing"
//
//	"github.com/stretchr/testify/mock"
//)
//
//type baseControllerMock struct{}
//
//func (cMock *baseControllerMock) GetData(c *fiber.Ctx, a interface{}) []string {
//	var errors []string
//
//	_ = c.BodyParser(a)
//
//	field, ok := a.(*userReqBody)
//	if ok && len(field.Name) <= 0 {
//		errors = append(errors, "New error")
//	}
//
//	return errors
//}
//
//type userMock struct{}
//
//func (uMock *userMock) CreateUser(name string, email string) (sql.Result, error) {
//	var err error = nil
//	if name == "test error" {
//		err = errors.New("test")
//	}
//	return nil, err
//}
//
//// func TestNewUserController(t *testing.T) {
//// 	userController := NewUserController()
//// 	if userController == nil {
//// 		t.Error("Request should return 500 status code")
//// 	}
//// }
//
//// func TestCreateUserErrorToInsert(t *testing.T) {
//// 	userController := UserController{
//// 		base:       &baseControllerMock{},
//// 		repository: &userMock{},
//// 	}
//
//// 	app := fiber.New()
//// 	app.Post("/user", userController.CreateUser)
//
//// 	bodyReader := strings.NewReader(`{"name":"test error", "email":"test@test.com"}`)
//// 	req := httptest.NewRequest("POST", "/user", bodyReader)
//// 	req.Header.Set("Content-Type", "application/json")
//
//// 	res, err := app.Test(req, -1)
//// 	if res.StatusCode != 500 || err != nil {
//// 		t.Error("Request should return 500 status code")
//// 	}
//// }
//
//func TestCreateUserErrorToInsertMockeryV2(t *testing.T) {
//	baseCtrlMock := mocks2.NewIBaseController(t)
//
//	baseCtrlMock.
//		EXPECT().
//		GetData(mock.Anything, mock.Anything).
//		Once().
//		Run(func(args mock.Arguments) {
//			fmt.Println(args...)
//		}).
//		Return([]string{})
//
//	repoMock := mocks2.NewIUser(t)
//	repoMock.EXPECT().CreateUser(
//		"",
//		"",
//	).
//		Run(func(name, email string) { fmt.Println(name, email) }).
//		Return(nil, errors.New("test"))
//
//	userController := UserController{
//		base:       baseCtrlMock,
//		repository: repoMock,
//	}
//
//	app := fiber.New()
//	app.Post("/user", userController.CreateUser)
//
//	bodyReader := strings.NewReader(`{"name":"test error", "email":"test@test.com"}`)
//	req := httptest.NewRequest("POST", "/user", bodyReader)
//	req.Header.Set("Content-Type", "application/json")
//
//	res, err := app.Test(req, -1)
//	if res.StatusCode != 500 || err != nil {
//		t.Error("Request should return 500 status code")
//	}
//}
//
//// func TestCreateUserErrorInvalidPayload(t *testing.T) {
//// 	userController := UserController{
//// 		base:       &baseControllerMock{},
//// 		repository: &userMock{},
//// 	}
//
//// 	app := fiber.New()
//// 	app.Post("/user", userController.CreateUser)
//
//// 	bodyReader := strings.NewReader(`{"names":"test", "emails":"test@test.com"}`)
//// 	req := httptest.NewRequest("POST", "/user", bodyReader)
//// 	req.Header.Set("Content-Type", "application/json")
//
//// 	res, err := app.Test(req, -1)
//// 	if res.StatusCode != 500 || err != nil {
//// 		t.Errorf("Request should return 500 status code returned: %d", res.StatusCode)
//// 	}
//// }
//
//// func TestCreateUser(t *testing.T) {
//// 	userController := UserController{
//// 		base:       &baseControllerMock{},
//// 		repository: &userMock{},
//// 	}
//
//// 	app := fiber.New()
//// 	app.Post("/user", userController.CreateUser)
//
//// 	bodyReader := strings.NewReader(`{"name":"test", "email":"test@test.com"}`)
//// 	req := httptest.NewRequest("POST", "/user", bodyReader)
//// 	req.Header.Set("Content-Type", "application/json")
//
//// 	res, err := app.Test(req, -1)
//// 	if res.StatusCode != 201 || err != nil {
//// 		t.Errorf("Request should return 201 status code returned: %d", res.StatusCode)
//// 	}
//// }
