package model

/**
	Legend
	Request Status
	[0] = Available
	[1] = Accepted
	[2] = Canceled
	[3] = Trip Start
	[4] = Trip Finish
**/

/*
	Struct Model & Class
*/
type Driver struct {
	Id string
}

type Passenger struct {
	Id string
}

var (
	LoggedInDriver    DriverInterface
	LoggedInPassenger PassangerInterface
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

type LoginCLaim struct {
	Id       string
	Password string
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
	Password string
}
