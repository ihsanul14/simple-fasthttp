// package middleware

// import (
// 	"encoding/json"
// 	"errors"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strings"

// 	"github.com/gin-contrib/cors"
// 	"github.com/gin-gonic/gin"
// 	"gopkg.in/resty.v1"
// )

// var jwt_msg = "message"
// var jwt_status = "status"
// var jwt_auth = "Authorization"
// var content_type = "Content-Type"

// // Errors related to JWT authentication service (not all of them used by now)
// var (
// 	// ErrEmptyAuthHeader can be thrown if authing with a HTTP header, the Auth header needs to be set
// 	ErrEmptyAuthHeader = errors.New("auth header is empty")

// 	// ErrInvalidAuthHeader indicates auth header is invalid, could for example have the wrong Realm name
// 	ErrInvalidAuthHeader = errors.New("auth header is invalid")
// )

// // Structs used in auth system
// type (
// 	// JWTClaims for JWT token
// 	JWTClaims struct {
// 		AuthUser
// 		StandardClaims
// 	}

// 	// StandardClaims standard claims
// 	StandardClaims struct {
// 		Audience  string `json:"aud,omitempty"`
// 		ExpiresAt int64  `json:"exp,omitempty"`
// 		Id        string `json:"jti,omitempty"`
// 		IssuedAt  int64  `json:"iat,omitempty"`
// 		Issuer    string `json:"iss,omitempty"`
// 		NotBefore int64  `json:"nbf,omitempty"`
// 		Subject   string `json:"sub,omitempty"`
// 	}

// 	// ValidateAccessRequest is used to request handler
// 	ValidateAccessRequest struct {
// 		Roles []string `json:"roles" gorm:"roles"`
// 	}

// 	// AuthUser is used as user model
// 	AuthUser struct {
// 		Pernr          string  `json:"pernr" gorm:"column:PERNR"`
// 		Sname          string  `json:"sname" gorm:"column:SNAME"`
// 		KodeDepartemen string  `json:"kode_departemen" gorm:"column:KODE_DEPARTEMEN"`
// 		Description1   string  `json:"description_1" gorm:"column:DESCRIPTION_1"`
// 		Description2   string  `json:"description_2" gorm:"column:DESCRIPTION_2"`
// 		OrgehTx        string  `json:"orgeh_tx" gorm:"column:ORGEH_TX"`
// 		StellTx        string  `json:"stell_tx" gorm:"column:STELL_TX"`
// 		Rgdesc         string  `json:"rgdesc" gorm:"column:RGDESC"`
// 		Mbdesc         string  `json:"mbdesc" gorm:"column:MBDESC"`
// 		Brdesc         string  `json:"brdesc" gorm:"column:BRDESC"`
// 		Region         string  `json:"region" gorm:"column:REGION"`
// 		Mainbranch     float64 `json:"mainbr" gorm:"column:MAINBR"`
// 		Branch         float64 `json:"branch" gorm:"column:BRANCH"`
// 	}
// )

// // Function to get JWT token from header
// // token must exists in "Authorization" header with "Bearer<space>" prefix
// func jwtFromHeader(c *fiber.Context) (string, error) {
// 	authHeader := c.Request.Header.Get(jwt_auth)

// 	if authHeader == "" {
// 		return "", ErrEmptyAuthHeader
// 	}

// 	parts := strings.SplitN(authHeader, " ", 2)
// 	if !(len(parts) == 2 && parts[0] == "Bearer") {
// 		return "", ErrInvalidAuthHeader
// 	}

// 	return parts[1], nil
// }

// // Funciton to Validate access to Auth Server
// // if roles empty then only validate user's token validity pass the request
// // otherwise, check the roles with current user roles (AccessLevel)
// func validateAccess(tokenString string, roles []string) (JWTClaims, error) {
// 	// call auth service
// 	jsonRoles, _ := json.Marshal(ValidateAccessRequest{
// 		Roles: roles,
// 	})

// 	var claims JWTClaims

// 	resp, err := resty.R().
// 		SetHeader("Content-Type", "application/json").
// 		SetHeader(jwt_auth, "Bearer "+tokenString).
// 		SetBody(jsonRoles).
// 		Post(os.Getenv("AUTH_SERVER") + `auth/validateAccess`)
// 	if err != nil {
// 		return claims, err
// 	} else {
// 		if resp.StatusCode() != http.StatusOK {
// 			res := fiber.H{
// 				jwt_status: "",
// 				jwt_msg:    "",
// 			}
// 			json.Unmarshal(resp.Body(), &res)
// 			errString := res[jwt_msg].(string)
// 			err = errors.New(errString)
// 			log.Print("JWT Response is Not OK -> ", resp.String())
// 			return claims, err
// 		} else {
// 			_ = json.Unmarshal(resp.Body(), &claims)
// 			log.Print("JWT Response  -> ", resp.String())
// 		}

// 	}

// 	return claims, err
// }

// func getDefaultCorsJWT() fiber.HandlerFunc {
// 	config := cors.DefaultConfig()
// 	config.AllowAllOrigins = true
// 	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
// 	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", jwt_auth, "X-Requested-With", "Accept",
// 		"Access-Control-Allow-Headers", "Accept-Encoding", "X-CSRF-Token"}
// 	return cors.New(config)
// }

// // JWTMiddleware to enforce JWT authorization to routes
// func JWTMiddleware(roles ...string) fiber.HandlerFunc {

// 	return func(c *fiber.Context) {
// 		tokenString, err := jwtFromHeader(c)

// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, fiber.H{
// 				jwt_status: http.StatusUnauthorized,
// 				jwt_msg:    err.Error(),
// 			})
// 			return
// 		}

// 		claims, err := validateAccess(tokenString, roles)
// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, fiber.H{
// 				jwt_status: http.StatusUnauthorized,
// 				jwt_msg:    err.Error(),
// 			})
// 			return

// 		}
// 		// set parameter to be passed to handler
// 		c.Set("JWT_TOKEN", tokenString)
// 		c.Set("JWT_CLAIMS", claims)
// 		c.Next()
// 	}
// }
