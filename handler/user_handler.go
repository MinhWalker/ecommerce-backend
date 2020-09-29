package handler

import (
	"ecommerce-backend/model"
	req "ecommerce-backend/model/req"
	"ecommerce-backend/repository"
	"ecommerce-backend/security"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

// UserHandler process all logic relate to user account
type UserHandler struct {
	UserRepo repository.UserRepo
}

// HandleSignUp handle user sign up
// SignUp godoc
// @Summary Create new account
// @Tags user-service
// @Accept  json
// @Produce  json
// @Param data body req.ReqSignUp true "user"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /user/sign-up [post]
func (u UserHandler) HandleSignUp(c echo.Context) error {
	req := req.SignUp{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	hash := security.HashAndSalt([]byte(req.Password))
	role := model.MEMBER.String()

	userID, err := uuid.NewUUID()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user := model.User{
		UserID:   userID.String(),
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hash,
		Address:  req.Address,
		FullName: req.FullName,
		Avatar:   req.Avatar,
		Role:     role,
	}

	user, err = u.UserRepo.SaveUser(c.Request().Context(), user)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusConflict, model.Response{
			StatusCode: http.StatusConflict,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	//gen token
	token, err := security.GenToken(user)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user.Token = token

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Success to Save User!",
		Data:       user,
	})

}

// HandleSignIn handle user sign in
// SignIn godoc
// @Summary access user login
// @Tags user-service
// @Accept  json
// @Produce  json
// @Param data body req.RepSignIn true "user"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /user/sign-in [post]
func (u UserHandler) HandleSignIn(c echo.Context) error {
	req := req.SignIn{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user, err := u.UserRepo.CheckLogin(c.Request().Context(), req.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	//check pass
	isTheSame := security.ComparePasswords(user.Password, []byte(req.Password))
	if !isTheSame {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    "Fail to Login",
			Data:       nil,
		})
	}

	//gen token
	token, err := security.GenToken(user)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	user.Token = token

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Success to login",
		Data:       user,
	})
}

// HandleProfile handle user profile
// Profile godoc
// @Summary access user login
// @Tags user-service
// @Accept  json
// @Produce  json
// @Param data body req.RepSignIn true "user"
// @Success 200 {object} model.Response
// @Failure 401 {object} model.Response
// @Router /user/profile [get]
func (u UserHandler) HandleProfile(c echo.Context) error {
	//userId := c.Param("id")

	tokenData := c.Get("user").(*jwt.Token)             // convert to jwt.Token type
	claims := tokenData.Claims.(*model.JwtCustomClaims) // convert to model.JwtCustomClaims type

	user, err := u.UserRepo.GetUserById(c.Request().Context(), claims.UserId)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Success!",
		Data:       user,
	})
}
