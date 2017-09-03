package main

import (
	"github.com/degananda/ridesharing/handler"
	"github.com/degananda/ridesharing/model"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Mysql Db Connection
	model.DbConnect()
	/*
		Route
	*/
	router := gin.Default()
	/**
		Public endpoint
	**/
	public := router.Group("/api")
	public.POST("/driver/login", handler.DriverLogin())
	public.POST("/passenger/login", handler.PassengerLogin())
	/**
		Private endpoint
	**/
	private := router.Group("/pvt/")
	private.Use(Auth(model.Secretkey))
	private.PUT("/driver/connect", handler.DriverConnect())
	private.PUT("/driver/sendpresence", handler.DriverPresence())
	private.GET("/driver/receiverequest", handler.DriverReceiveRequest())
	private.PUT("/driver/acceptrequest/:id", handler.DriverAcceptRequest())
	private.PUT("/driver/sendlocation", handler.DriverSendLocation())
	private.PUT("/driver/starttrip/:id", handler.DriverStartTrip())
	private.PUT("/driver/endtrip/:id", handler.DriverEndTrip())
	private.PUT("/passenger/connect", handler.PassangerConnect())
	private.PUT("/passenger/sendpresence", handler.PassangerPresence())
	private.POST("/passenger/sendrequest/:id", handler.PassangerRequest())
	private.GET("/passenger/receiverequest", handler.PassengerReceiveRequest())
	private.GET("/passenger/receivelocation/:id", handler.PassengerReceiveLocation())
	router.Run()
}
