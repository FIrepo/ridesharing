package main

import (
	"errors"
	"strings"

	"github.com/degananda/ridesharing/model"
	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

func Auth(secret string) gin.HandlerFunc {

	type CustomClaim struct {
		Id  string
		Exp string
	}

	return func(c *gin.Context) {
		var issue string = ""
		_, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt_lib.Token) (interface{}, error) {
			b := ([]byte(secret))
			requestURL := c.Request.RequestURI
			claims := token.Claims.(jwt_lib.MapClaims)
			if claims["userType"].(string) == "driver" {
				model.LoggedInDriver = &model.Driver{claims["Id"].(string)}
			} else {
				model.LoggedInPassenger = &model.Passenger{claims["Id"].(string)}
			}

			if checkAuth(claims["userType"].(string), splitURI(requestURL)) {
				return b, nil

			} else {
				issue = "token not authorized"
				return b, errors.New("token not authorized")
			}

		})

		if err != nil {
			c.Abort()
			c.JSON(401, gin.H{
				"msg":    "Invalid Token",
				"status": "-1",
				"issue":  issue,
			})
		}
	}
}

func splitURI(uri string) string {
	splittedUri := strings.Split(uri, "/")
	var newUri string
	for i, v := range splittedUri {
		if i <= 3 {
			newUri += v + "/"
		}
	}
	return newUri
}

func checkAuth(userType string, url string) bool {
	/**
		Endpoint permission for each user
	**/
	authMap := make(map[string]map[string]bool)
	// init first layer
	authMap["driver"] = make(map[string]bool)
	authMap["passenger"] = make(map[string]bool)
	// second layer map
	authMap["driver"]["/pvt/driver/connect/"] = true
	authMap["driver"]["/pvt/driver/sendpresence/"] = true
	authMap["driver"]["/pvt/driver/receiverequest/"] = true
	authMap["driver"]["/pvt/driver/acceptrequest/"] = true
	authMap["driver"]["/pvt/driver/sendlocation/"] = true
	authMap["driver"]["/pvt/driver/starttrip/"] = true
	authMap["driver"]["/pvt/driver/endtrip/"] = true
	authMap["passenger"]["/pvt/passenger/connect/"] = true
	authMap["passenger"]["/pvt/passenger/sendpresence/"] = true
	authMap["passenger"]["/pvt/passenger/sendrequest/"] = true
	authMap["passenger"]["/pvt/passenger/receiverequest/"] = true
	authMap["passenger"]["/pvt/passenger/receivelocation/"] = true
	if authMap[userType][url] {
		return true
	} else {
		return false
	}

}
