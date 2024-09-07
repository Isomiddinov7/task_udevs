package handler

import (
	"fmt"
	"net/http"
	"strings"
	"task_udevs/api/models"
	initializers "task_udevs/initializer"
	"task_udevs/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func (h *Handler) DeserializeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		fmt.Println("hi")

		var access_token string
		cookie, err := ctx.Cookie("access_token")

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		if access_token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config, _ := initializers.LoadConfig(".")
		sub, err := utils.ValidateToken(access_token, config.AccessTokenPublicKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		err = h.strg.User().DeserializeUser(ctx.Request.Context(), models.GetUserById{Id: cast.ToString(sub)})
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
			return
		}

		ctx.Set("currentUser", err)
		ctx.Next()
	}
}
