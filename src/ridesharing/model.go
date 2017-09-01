package main

/*
	Struct Model & Class
*/
type Driver struct {
	id string
}

type Passenger struct {
	id string
}

var (
	loggedDriver    Driver
	loggedPassenger Passenger
)

/*
	JSON Body Model
*/
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SendRequest struct {
	Lat string `json:"lat" binding:"required"`
	Lon string `json:"lon" binding:"required"`
}

type AcceptRequest struct {
	IdRequest string `json:"IdRequest" binding:"required"`
}

type SendLocation struct {
	Lat string `json:"lat" binding:"required"`
	Lon string `json:"lon" binding:"required"`
}

type ReceiveLocation struct {
	Lat string `json:"lat" binding:"required"`
	Lon string `json:"lon" binding:"required"`
}

type EndTrip struct {
	Distance string `json:"distance" binding:"required"`
	Time     string `json:"time" binding:"required"`
}

/* end of json body model */

/**
	JSON Response Model
**/
type Request struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Lat   string `json:"lat"`
	Lon   string `json:"lon"`
}

type PassengerRequest struct {
	DriverName  string `json:"name"`
	DriverEmail string `json:"email"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
}

/*
	Auth user
*/

type PwdTmp struct {
	password string
}

type UserSingleton struct {
	Id   string
	Type string
}

var (
	loggedUser UserSingleton
)
