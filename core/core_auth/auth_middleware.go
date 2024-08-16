package core_auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const HEADER_PREFIX = "Bearer "

func AuthMiddleware(config *AuthConfig) gin.HandlerFunc {
	return func(context *gin.Context) {
		// 验证是否启用认证过滤
		if !config.Enabled {
			context.Next()
			return
		}

		url := context.Request.URL.Path

		// 验证是否处于包含内
		if config.IncludeUrls != nil && len(config.IncludeUrls) > 0 {
			isInclude := false
			for _, iurl := range config.IncludeUrls {
				if strings.HasPrefix(url, iurl) {
					isInclude = true
					break
				}
			}

			if !isInclude {
				context.Next()
				return
			}
		}

		// 验证是否处于排除内
		if config.ExcludeUrls != nil && len(config.ExcludeUrls) > 0 {
			isExclude := false
			for _, eurl := range config.ExcludeUrls {
				if strings.HasPrefix(url, eurl) {
					isExclude = true
					break
				}
			}

			if isExclude {
				context.Next()
				return
			}
		}

		// 获取令牌
		headerString := context.Request.Header.Get("Authorization")
		tokenString, preFound := strings.CutPrefix(headerString, HEADER_PREFIX)
		tokenString = strings.TrimSpace(tokenString)

		if !preFound || len(tokenString) == 0 {
			context.Status(http.StatusUnauthorized)
			context.Abort()
			return
		}

		tokenClaims := &AuthClaims{}
		err := ParseToken(config.PublicKey, tokenString, tokenClaims)

		// 解析失败、令牌过期、令牌发行人不是当前、客户端不存在
		if err != nil ||
			tokenClaims.Valid() != nil ||
			tokenClaims.Issuer != config.TokenIssuer {
			context.Status(http.StatusUnauthorized)
			context.Abort()
			return
		}

		// 认证通过，将认证令牌设置到上下文中
		SetAuthContext(context, &AuthContext{
			Claims: tokenClaims,
			Scopes: strings.Split(tokenClaims.Scope, " "),
		})

		context.Next()
	}
}
