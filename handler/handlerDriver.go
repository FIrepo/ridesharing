package handler

import (
	"strconv"
	"time"

	"github.com/degananda/ridesharing/model"
	jwt "github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

func DriverLogin() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var (
			jsonLogin model.Login
		)
		err := c.BindJSON(&jsonLogin)
		if err == nil {
			model.LoggedInDriver = &model.Driver{}
			errfree, isAuthenticated, userId := model.LoggedInDriver.Login(jsonLogin.Username, jsonLogin.Password)

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
						"userType": "driver",
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

func DriverConnect() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		errfree, issue := model.LoggedInDriver.Connect()
		if errfree {
			c.JSON(200, gin.H{
				"status": "1",
				"msg":    "success",
			})
		} else {
			c.JSON(400, gin.H{
				"status": "-1",
				"msg":    "failed",
				"issue":  issue.Error(),
			})
		}
	}
	return gin.HandlerFunc(fn)
}

func DriverPresence() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		errfree, issue := model.LoggedInDriver.SendPresence()
		if errfree {
			c.JSON(200, gin.H{
				"status": "1",
				"msg":    "success",
			})
		} else {
			c.JSON(400, gin.H{
				"status": "-1",
				"msg":    "success",
				"issue":  issue.Error(),
			})
		}
	}
	return gin.HandlerFunc(fn)
}

func DriverReceiveRequest() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		errfree, request, issue := model.LoggedInDriver.ReceiveRequest()
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
				"issue":  issue.Error(),
			})
		}
	}
	return gin.HandlerFunc(fn)
}

func DriverAcceptRequest() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		errfree, issue := model.LoggedInDriver.AcceptRequest(c.Param("id"))
		if errfree {
			c.JSON(200, gin.H{
				"status": "1",
				"msg":    "success",
			})
		} else {
			c.JSON(400, gin.H{
				"status": "-1",
				"msg":    "failed",
				"issue":  issue.Error(),
			})
		}
	}
	return gin.HandlerFunc(fn)

}

func DriverSendLocation() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var sendLocationInput model.SendLocation
		err := c.BindJSON(&sendLocationInput)
		// make sure lat is float64
		_, err = strconv.ParseFloat(sendLocationInput.Lat, 64)
		// make sure lon is float64
		_, err = strconv.ParseFloat(sendLocationInput.Lon, 64)
		if err == nil {
			errfree, issue := model.LoggedInDriver.SendLocation(sendLocationInput.Lat, sendLocationInput.Lon)
			if errfree {
				c.JSON(200, gin.H{
					"status": "1",
					"msg":    "success",
				})
			} else {
				c.JSON(200, gin.H{
					"status": "-1",
					"msg":    "failed",
					"issue":  issue.Error(),
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

func DriverStartTrip() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		errfree, issue := model.LoggedInDriver.StartTrip(c.Param("id"))
		if errfree {
			c.JSON(200, gin.H{
				"status": "1",
				"msg":    "success",
			})
		} else {
			c.JSON(200, gin.H{
				"status": "-1",
				"msg":    "failed",
				"issue":  issue.Error(),
			})
		}
	}
	return gin.HandlerFunc(fn)

}

func DriverEndTrip() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var endTripInput model.EndTrip
		err := c.BindJSON(&endTripInput)
		// make sure lat is float64
		_, err = strconv.ParseFloat(endTripInput.Distance, 64)
		// make sure lon is unsign integer 64 (must be positive int)
		_, err = strconv.ParseUint(endTripInput.Time, 10, 64)
		if err == nil {
			errfree, issue := model.LoggedInDriver.EndTrip(c.Param("id"), endTripInput.Distance, endTripInput.Time)
			if errfree {
				c.JSON(200, gin.H{
					"status": "1",
					"msg":    "success",
				})
			} else {
				c.JSON(200, gin.H{
					"status": "-1",
					"msg":    "failed",
					"issue":  issue.Error(),
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
