package middleware

import (
	"eniqloStore/usecase"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// import (
// 	"7Zero4/usecase"
// 	"net/http"

// 	"strings"

// 	"github.com/gin-gonic/gin"
// )

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

type AuthTokenMiddleware interface {
	RequiredToken() gin.HandlerFunc
}

type authTokenMiddleware struct {
	tokenUsecase usecase.TokenUsecase
}

func checkBypassAPI(c *gin.Context) bool {
	bypassAPI := []string{ // store API that dont need auth bearer
		"/v1/staff/register",
		"/v1/staff/login",
	}

	for _, v := range bypassAPI {
		if c.Request.URL.Path == v {
			return true
		}
	}
	return false
}

func (a *authTokenMiddleware) RequiredToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if checkBypassAPI(c) {
			c.Next()
		} else {
			h := authHeader{}
			if err := c.ShouldBindHeader(&h); err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthorized",
				})
				c.Abort()
				return
			}

			tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)
			if tokenString == "" {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthorized",
				})
				c.Abort()
				return
			}

			pass, err := a.tokenUsecase.VerifyAccessToken(tokenString)
			fmt.Println("ini hasilnya token --> ", pass)
			if err != nil || !pass {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthorized",
				})
				c.Abort()
				return
			}
			c.Next()
		}
	}
}

func NewAuthTokenMiddleware(tokenUsecase usecase.TokenUsecase) AuthTokenMiddleware {
	middleware := new(authTokenMiddleware)
	middleware.tokenUsecase = tokenUsecase
	return middleware
}
