package controller

import (
	"encoding/json"
	"errors"
	"final-project/dto"
	"final-project/helper"
	"final-project/service"
	"net/http"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{
		userService: userService,
	}
}

// UserRegister godoc
// @Summary register a new user
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.UserRegister true "required body"
// @Success 201 {object} helper.Response[dto.UserCreateResponse]
// @Failure 400 {object} helper.Response[any]
// @Failure 409 {object} helper.Response[any]
// @Failure 500 {object} helper.Response[any]
// @Router /users/register [post]
func (u *userController) Register(w http.ResponseWriter, r *http.Request) {
	var (
		data dto.UserRequest
		resp = helper.NewResponse[dto.UserCreateResponse](helper.UserRegister)
	)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		resp.Error(err).Code(http.StatusBadRequest).Send(w)
		return
	}

	err = data.ValidateCreate()
	if err != nil {
		resp.Error(err).Code(http.StatusBadRequest).Send(w)
		return
	}

	user, err := u.userService.Create(r.Context(), data)
	if err != nil {
		respErr := new(helper.ResponseError)
		if errors.As(err, &respErr) {
			resp.Error(respErr).Code(respErr.Code()).Send(w)
			return
		}
		resp.Error(err).Code(http.StatusInternalServerError).Send(w)
		return
	}

	resp.Success(true).Data(user).Code(http.StatusCreated).Send(w)
}

// UserLogin godoc
// @Summary login user
// @Tags users
// @Accept json
// @Produce json
// @Param request body dto.UserLogin true "required body"
// @Success 200 {object} helper.Response[dto.UserLoginResponse]
// @Failure 400 {object} helper.Response[any]
// @Failure 401 {object} helper.Response[any]
// @Failure 500 {object} helper.Response[any]
// @Router /users/login [post]
func (u *userController) Login(w http.ResponseWriter, r *http.Request) {
	var (
		data dto.UserRequest
		resp = helper.NewResponse[dto.UserLoginResponse](helper.UserLogin)
	)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		resp.Error(err).Code(http.StatusBadRequest).Send(w)
		return
	}

	err = data.ValidateLogin()
	if err != nil {
		resp.Error(err).Code(http.StatusBadRequest).Send(w)
		return
	}

	token, err := u.userService.Login(r.Context(), data)
	if err != nil {
		respErr := new(helper.ResponseError)
		if errors.As(err, &respErr) {
			resp.Error(respErr).Code(respErr.Code()).Send(w)
			return
		}
		resp.Error(err).Code(http.StatusInternalServerError).Send(w)
		return
	}

	resp.Success(true).Data(token).Code(http.StatusOK).Send(w)
}

// UserUpdate godoc
// @Summary update user
// @Tags users
// @Accept json
// @Produce json
// @Security BearerToken
// @Param request body dto.UserUpdate true "required body"
// @Success 200 {object} helper.Response[dto.UserUpdateResponse]
// @Failure 400 {object} helper.Response[any]
// @Failure 404 {object} helper.Response[any]
// @Failure 409 {object} helper.Response[any]
// @Failure 500 {object} helper.Response[any]
// @Router /users [put]
func (u *userController) Update(w http.ResponseWriter, r *http.Request) {
	var (
		data dto.UserRequest
		resp = helper.NewResponse[dto.UserUpdateResponse](helper.UserUpdate)
	)

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		resp.Error(err).Code(http.StatusBadRequest).Send(w)
		return
	}

	err = data.ValidateUpdate()
	if err != nil {
		resp.Error(err).Code(http.StatusBadRequest).Send(w)
		return
	}

	user, err := u.userService.Update(r.Context(), data)
	if err != nil {
		respErr := new(helper.ResponseError)
		if errors.As(err, &respErr) {
			resp.Error(respErr).Code(respErr.Code()).Send(w)
			return
		}
		resp.Error(err).Code(http.StatusInternalServerError).Send(w)
		return
	}

	resp.Success(true).Data(user).Code(http.StatusOK).Send(w)
}

// UserDelete godoc
// @Summary delete user
// @Tags users
// @Produce json
// @Security BearerToken
// @Success 200 {object} helper.Response[any]
// @Failure 404 {object} helper.Response[any]
// @Failure 500 {object} helper.Response[any]
// @Router /users [delete]
func (u *userController) Delete(w http.ResponseWriter, r *http.Request) {
	var resp = helper.NewResponse[any](helper.UserDelete)

	err := u.userService.Delete(r.Context())
	if err != nil {
		respErr := new(helper.ResponseError)
		if errors.As(err, &respErr) {
			resp.Error(respErr).Code(respErr.Code()).Send(w)
			return
		}
		resp.Error(err).Code(http.StatusInternalServerError).Send(w)
		return
	}

	resp.Success(true).Code(http.StatusOK).Send(w)
}
