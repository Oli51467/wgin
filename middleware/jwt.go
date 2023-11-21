package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"wgin/common/response"
	"wgin/global"
	"wgin/service"
)

func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.TokenFail(c)
			c.Abort()
			return
		}
		tokenStr = tokenStr[len(service.TokenType)+1:]

		// Token 解析校验
		token, err := jwt.ParseWithClaims(tokenStr, &service.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(global.App.Config.Jwt.Secret), nil
		})
		if err != nil || service.JwtService.IsInBlacklist(tokenStr) {
			response.TokenFail(c)
			c.Abort()
			return
		}

		claims := token.Claims.(*service.CustomClaims)
		// Token 发布者校验
		if claims.Issuer != GuardName {
			response.TokenFail(c)
			c.Abort()
			return
		}
		// 将token和id保存到上下文中
		c.Set("token", token)
		c.Set("id", claims.Id)

		// 在过期前的 2 小时内，如果用户访问了接口，就颁发新的 token 给客户端（设置响应头），同时把旧 token 加入黑名单
		if claims.ExpiresAt-time.Now().Unix() < global.App.Config.Jwt.RefreshPeriod {
			// 要避免并发请求导致 token 重复刷新的情况，这里需要上锁
			lock := global.MakeLock("refresh_token_key"+claims.Id, global.App.Config.Jwt.JwtBlacklistPeriod)
			// 如果能锁上
			if lock.Get() {
				err, user := service.JwtService.GetUserInfo(GuardName, claims.Id)
				if err != nil {
					global.App.Logger.Error(err.Error())
					lock.Release()
				} else {
					refreshToken, err, _ := service.JwtService.CreateToken(GuardName, user)
					if err != nil {
						response.BusinessFail(c, err.Error())
						return
					}
					c.Header("refresh-token", refreshToken.AccessToken)
					c.Header("refresh-token-ttl", strconv.Itoa(refreshToken.ExpiresIn))
					_ = service.JwtService.JoinBlackList(token)
				}
			}
		}
	}
}
