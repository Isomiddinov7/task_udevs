package handler

import (
	"task_udevs/api/models"
	"task_udevs/pkg/utils"

	initializers "task_udevs/initializer"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @ID login-user
// @Router /login-user [POST]
// @Summary login
// @Description login
// @Tags User
// @Accept json
// @Produce json
// @Param object body models.UserAuthRequest true "LoginUserRequestBody"
// @Success 200 {object} Response{data=string} "Token"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) AuthUser(c *gin.Context) {
	var user models.UserAuthRequest

	err := c.ShouldBindJSON(&user)
	if err != nil {
		handleResponse(c, 400, "http.BadRequest"+err.Error())
		return
	}
	resp, err := h.strg.User().Auth(
		c.Request.Context(),
		user,
	)
	if err != nil {
		handleResponse(c, 400, "BadRequest"+"login or password is wrong")
		return
	}
	config, _ := initializers.LoadConfig(".")
	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, resp.UserId, config.AccessTokenPrivateKey)
	if err != nil {
		handleResponse(c, 500, "InternalServerError: Could not generate token")
		return
	}
	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, resp.UserId, config.RefreshTokenPrivateKey)
	if err != nil {
		handleResponse(c, 500, "InternalServerError: Could not generate token")
		return
	}

	c.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	handleResponse(c, 200, "OK, token: access_token")
}
