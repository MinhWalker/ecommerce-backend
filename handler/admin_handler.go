package handler

import (
	"ecommerce-backend/model"
	"ecommerce-backend/model/req"
	"ecommerce-backend/repository"
	"ecommerce-backend/security"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

type AdminHandler struct {
	UserRepo repository.UserRepo
}

// GenToken handle admin
// Profile godoc
// @Summary create Token for admin
// @Tags user-service
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Response
// @Router /admin/token [get]
func (a AdminHandler) GenToken(c echo.Context) error {
	userId, _ := uuid.NewUUID()
	token, _ := security.GenToken(model.User{
		UserID:    userId.String(),
		Role:      model.ADMIN.String(),
	})

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Success to gen token!",
		Data:       token,
	})
}

// HandleSignUp handle admin sign up
// SignUp godoc
// @Summary Create new account for admin
// @Tags user-service
// @Accept  json
// @Produce  json
// @Param data body req.ReqSignUp true "user"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 403 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /admin/sign-up [post]
func (a AdminHandler) HandleSignUp(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	if claims.Role != model.ADMIN.String() {
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    http.StatusText(http.StatusForbidden),
			Data:       nil,
		})
	}

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
	role := claims.Role

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

	user, err = a.UserRepo.SaveUser(c.Request().Context(), user)
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

// HandleSignIn handle admin sign in
// SignIn godoc
// @Summary access admin login
// @Tags user-service
// @Accept  json
// @Produce  json
// @Param data body req.RepSignIn true "user"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 403 {object} model.Response
// @Failure 401 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /admin/sign-in [post]
func (a AdminHandler) HandleSignIn(c echo.Context) error {
	req := req.SignIn{}
	if err := c.Bind(&req); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	user, err := a.UserRepo.CheckLogin(c.Request().Context(), req.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, model.Response{
			StatusCode: http.StatusUnauthorized,
			Message:    err.Error(),
			Data:       nil,
		})
	}
	
	//check role ADMIN
	if user.Role != model.ADMIN.String() {
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    http.StatusText(http.StatusForbidden),
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

func (a AdminHandler) HandleDeleteUser(c echo.Context) error {
	userId := c.Param("id")
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*model.JwtCustomClaims)

	if claims.Role != model.ADMIN.String() {
		return c.JSON(http.StatusForbidden, model.Response{
			StatusCode: http.StatusForbidden,
			Message:    http.StatusText(http.StatusForbidden),
			Data:       nil,
		})
	}

	err := a.UserRepo.DeleteUsers(c.Request().Context(), userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		})
	}

	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    "Delete user Success!",
		Data:       nil,
	})
}
