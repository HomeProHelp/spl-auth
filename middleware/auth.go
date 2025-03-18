package middleware

import (
	"github/LissaiDev/spl-auth/pkg/token"
	"github/LissaiDev/spl-auth/utils"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		hdr := ctx.GetHeader("spl_auth_token")
		if hdr == "" {
			ctx.AbortWithStatusJSON(401, &utils.Response{
				Code: utils.AuthenticationCodes["invalid_token"],
				Data: map[string]string{},
			})
			return
		}

		claims, code := token.ValidateToken(hdr)
		if code != utils.AuthenticationCodes["success"] {
			ctx.AbortWithStatusJSON(401, &utils.Response{
				Code: code,
				Data: map[string]string{},
			})
			return
		}

		ctx.Set("ID", claims.ID)
		ctx.Next()
	}
}
