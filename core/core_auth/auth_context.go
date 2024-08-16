package core_auth

import "github.com/gin-gonic/gin"

type AuthContext struct {
	Claims *AuthClaims

	Scopes []string
}

func SetAuthContext(context *gin.Context, authContext *AuthContext) {
	context.Set("auth-context", authContext)
}

func GetAuthContext(context *gin.Context) *AuthContext {
	value, exists := context.Get("auth-context")

	if exists {
		return value.(*AuthContext)
	}

	return nil
}

func (context *AuthContext) HasScope(scope string) bool {
	for _, v := range context.Scopes {
		if v == scope {
			return true
		}
	}

	return false
}

func (context *AuthContext) HasAnyScope(scopes ...string) bool {
	if len(scopes) < 1 {
		return false
	}

	for _, s := range scopes {
		if context.HasScope(s) {
			return true
		}
	}

	return false
}

func (context *AuthContext) HasAndScope(scopes ...string) bool {
	if len(scopes) < 1 {
		return false
	}

	for _, s := range scopes {
		if !context.HasScope(s) {
			return false
		}
	}

	return true
}
