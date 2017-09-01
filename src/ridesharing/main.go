package main

import (
	"database/sql"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

/**
	Mysql db pool connection global
**/
var db *sql.DB
var secretkey = "indonesiatanahairbeta"

func main() {
	// Mysql Db Connection
	var err error
	db, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/ridesharing")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	/*
		Route
	*/
	router := gin.Default()
	/**
		Public endpoint
	**/
	public := router.Group("/api")
	public.POST("/driver/login", driverLogin)
	public.POST("/passenger/login", passengerLogin)

	/**
		Private endpoint
	**/
	private := router.Group("/pvt/")
	private.Use(Auth(secretkey))
	private.PUT("/driver/connect", driverConnect)
	private.PUT("/driver/sendpresence", driverPresence)
	private.GET("/driver/receiverequest", driverReceiveRequest)
	private.PUT("/driver/acceptrequest", driverAcceptRequest)
	private.PUT("/driver/sendlocation", driverSendLocation)
	private.PUT("/driver/starttrip/:id", driverStartTrip)
	private.PUT("/driver/endtrip/:id", driverEndTrip)
	private.PUT("/passenger/connect", passangerConnect)
	private.PUT("/passenger/sendpresence", passangerPresence)
	private.POST("/passenger/sendrequest", passangerRequest)
	private.GET("/passenger/receiverequest", passengerReceiveRequest)
	private.GET("/passenger/receivelocation/:id", passengerReceiveLocation)

	router.Run()

}

func driverLogin(c *gin.Context) {
	var (
		jsonLogin   Login
		tmpPassword PwdTmp
	)
	errJsonBinding := c.BindJSON(&jsonLogin)
	if errJsonBinding == nil {
		var toBeAuth Driver
		stmt := db.QueryRow("SELECT id,password FROM driver WHERE email = ?", jsonLogin.Username)
		rowScan := stmt.Scan(&toBeAuth.id, &tmpPassword.password)
		if rowScan == nil {
			byteHash := []byte(tmpPassword.password)
			bytePassString := []byte(jsonLogin.Password)
			if isLoginCredValid(byteHash, bytePassString) {

				/*
					Generate Token Jwt HS256
				*/
				token := jwt.New(jwt.GetSigningMethod("HS256"))
				// Token configuration
				token.Claims = jwt.MapClaims{
					"Id":       toBeAuth.id,
					"exp":      time.Now().Add(time.Hour * 1).Unix(),
					"userType": "driver",
				}

				/*
					Token Generate
				*/
				tokenString, err := token.SignedString([]byte(secretkey))
				if err != nil {
					// Failed generate token
					c.JSON(500, gin.H{"message": "Could not generate token"})
				} else {
					// Success generate token
					c.JSON(200, gin.H{
						"status": "1",
						"msg":    "Login Success",
						"token":  tokenString,
					})
				}
			} else {
				c.JSON(400, gin.H{
					"status": "-1",
					"msg":    "Login failed",
				})
			}

		} else {
			c.JSON(400, gin.H{
				"status": "-1",
				"msg":    "Login failed",
			})
		}

	}

}

func passengerLogin(c *gin.Context) {
	var (
		jsonLogin   Login
		tmpPassword PwdTmp
	)
	errJsonBinding := c.BindJSON(&jsonLogin)
	if errJsonBinding == nil {
		var toBeAuth Passenger
		stmt := db.QueryRow("SELECT id,password FROM passenger WHERE email = ?", jsonLogin.Username)
		rowScan := stmt.Scan(&toBeAuth.id, &tmpPassword.password)
		if rowScan == nil {
			byteHash := []byte(tmpPassword.password)
			bytePassString := []byte(jsonLogin.Password)
			if isLoginCredValid(byteHash, bytePassString) {

				/*
					Generate Token Jwt HS256
				*/
				token := jwt.New(jwt.GetSigningMethod("HS256"))
				// Token configuration
				token.Claims = jwt.MapClaims{
					"Id":       toBeAuth.id,
					"exp":      time.Now().Add(time.Hour * 1).Unix(),
					"userType": "passenger",
				}

				/*
					Token Generate
				*/
				tokenString, err := token.SignedString([]byte(secretkey))
				if err != nil {
					// Failed generate token
					c.JSON(500, gin.H{"message": "Could not generate token"})
				} else {
					// Success generate token
					c.JSON(200, gin.H{
						"status": "1",
						"msg":    "Login Success",
						"token":  tokenString,
					})
				}
			} else {
				c.JSON(400, gin.H{
					"status": "-1",
					"msg":    "Login failed",
				})
			}

		} else {
			c.JSON(400, gin.H{
				"status": "-1",
				"msg":    "Login failexd",
			})
		}

	}

}

func isLoginCredValid(passwordHash []byte, passwordString []byte) bool {
	err := bcrypt.CompareHashAndPassword(passwordHash, passwordString)
	if err == nil {
		return true
	} else {
		return false
	}
}
