package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/degananda/ridesharing/model"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Passenger model.Passenger
type SendRequest model.SendRequest
type PassengerRequest model.PassengerRequest
type ReceiveLocation model.ReceiveLocation

func PassengerLogin() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var (
			jsonLogin model.Login
		)
		err := c.BindJSON(&jsonLogin)
		if err == nil {
			model.LoggedInPassenger = &model.Passenger{}
			errfree, isAuthenticated, userId := model.LoggedInPassenger.Login(jsonLogin.Username, jsonLogin.Password)

			if errfree {
				if isAuthenticated {
					// Login success
					/*
						Generate Token Jwt HS256
					*/
					token := jwt.New(jwt.GetSigningMethod("HS256"))
					// Token configuration
					token.Claims = jwt.MapClaims{
						"Id":       userId,
						"exp":      time.Now().Add(time.Hour * 1).Unix(),
						"userType": "passenger",
					}

					/*
						Token Generate
					*/
					tokenString, err := token.SignedString([]byte(model.Secretkey))

					if err == nil {
						// Success generate token
						c.JSON(200, gin.H{
							"status": "1",
							"msg":    "success",
							"token":  tokenString,
						})
					} else {
						// Failed generate token
						c.JSON(500, gin.H{
							"status": "-1",
							"msg":    "failed",
							"issue":  err.Error(),
						})
					}
				} else {
					// Login failed
					c.JSON(400, gin.H{
						"status": "-1",
						"msg":    "failed",
					})
				}
			} else {
				// email not found
				c.JSON(400, gin.H{
					"status": "-1",
					"msg":    "email not registered",
				})
			}
		} else {
			// Json body not match
			c.JSON(400, gin.H{
				"status": "-1",
				"msg":    "failed",
				"issue":  err,
			})
		}
	}
	return gin.HandlerFunc(fn)

}

func PassangerConnect() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		errfree, issue := model.LoggedInPassenger.Connect()
		if errfree {
			c.JSON(200, gin.H{
				"status": "1",
				"msg":    "success",
			})
		} else {
			c.JSON(200, gin.H{
				"status": "-1",
				"msg":    "failed",
				"issue":  issue,
			})
		}
	}
	return gin.HandlerFunc(fn)
}

func PassangerPresence() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		errfree, issue := model.LoggedInPassenger.SendPresence()
		if errfree {
			c.JSON(200, gin.H{
				"status": "1",
				"msg":    "success",
			})
		} else {
			c.JSON(200, gin.H{
				"status": "-1",
				"msg":    "failed",
				"issue":  issue,
			})
		}
	}
	return gin.HandlerFunc(fn)
}

func PassangerRequest() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var (
			sendRequestInput SendRequest
		)
		err := c.BindJSON(&sendRequestInput)
		// make sure lat is float64
		_, err = strconv.ParseFloat(sendRequestInput.Lat, 64)
		// make sure lon is unsign integer 64 (must be positive int)
		_, err = strconv.ParseFloat(sendRequestInput.Lon, 64)
		if err == nil {
			// json body valid
			errfree, issue := model.LoggedInPassenger.SendRequest(c.Param("id"), sendRequestInput.Lat, sendRequestInput.Lon)
			if errfree {
				c.JSON(200, gin.H{
					"status": "1",
					"msg":    "success",
				})
			} else {
				c.JSON(200, gin.H{
					"status": "-1",
					"msg":    "failed",
					"issue":  issue,
				})
			}
		} else {
			c.JSON(200, gin.H{
				"status": "-1",
				"msg":    "failed",
				"issues": err,
			})
		}
	}
	return gin.HandlerFunc(fn)

}

func PassengerReceiveRequest() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		errfree, request, issue := model.LoggedInPassenger.ReceiveRequest()
		if errfree {
			c.JSON(200, gin.H{
				"status": "1",
				"msg":    "success",
				"data":   request,
			})
		} else {
			c.JSON(400, gin.H{
				"status": "-1",
				"msg":    "failed",
				"issue":  issue,
			})
		}
	}
	return gin.HandlerFunc(fn)
}

func PassengerReceiveLocation() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		errfree, request, issue := model.LoggedInPassenger.ReceiveLocation(c.Param("id"))
		fmt.Println(request)
		if errfree {
			c.JSON(200, gin.H{
				"status": "1",
				"msg":    "success",
				"data":   request,
			})
		} else {
			c.JSON(400, gin.H{
				"status": "-1",
				"msg":    "failed",
				"issue":  issue,
			})
		}
	}
	return gin.HandlerFunc(fn)
}
