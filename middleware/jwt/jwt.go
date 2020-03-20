package jwt

import (
	log   "gvp/pkg/logging"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"gvp/pkg/util"
	"gvp/pkg/e"
	"gvp/models"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}else{
				name:=   claims.Username
				pwd:=    claims.Password
				log.Info("参与校验的用户名:",name,"密码：",pwd)
				ret:=  models.CheckAuth(name,pwd)
				if  ret  != true{
					log.Info("用户名和密码不匹配")
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code" : code,
				"msg" : e.GetMsg(code),
				"data" : data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}