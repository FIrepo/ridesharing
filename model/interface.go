package model

/**************************
	Interface
***************************/
type DriverInterface interface {
	/*
		1.Connect(email,password)
		[+] email : Registered driver email
		[+] password : Bcrypt hash
		Login and connect to ridesharing system
	*/

	/*
		Driver Login
	*/
	Login(username string, password string) (bool, bool, string)

	/*
		Get Driver ID
	*/
	GetId() string

	/*
		1. Connect() set Passenger is connected to ridesharing system (not login)
	*/
	Connect() (bool, error)

	/*
		2.sendPresence()
		Make Driver visible on the map (Presence)
	*/
	SendPresence() (bool, error)

	/*
		3. Receive Request
	*/
	ReceiveRequest() (bool, []Request, error)

	/*
		4. Accept Request
	*/
	AcceptRequest(idRequest string) (bool, error)

	/*
		5. Send Location
	*/
	SendLocation(lat string, lon string) (bool, error)

	/*
		6. Start Trip
	*/
	StartTrip(idRequest string) (bool, error)

	/*
		7. End Trip
	*/
	EndTrip(idRequest string, distance string, time string) (bool, error)
}

type PassangerInterface interface {

	/*
		Passenger login
	*/
	Login(username string, password string) (bool, bool, string)

	/*
		Get Passenger ID
	*/
	GetId() string
	/*
		1. Connect() set Passenger is connected to ridesharing system (not login)
	*/
	Connect() (bool, error)

	/*
		2. Make passenger visible on the map (Presence)
	*/
	SendPresence() (bool, error)

	/*
		3. Send Request
	*/
	SendRequest(idDriver string, lat string, lon string) (bool, error)

	/*
		4. Receive Request
	*/
	ReceiveRequest() (bool, []PassengerRequest, error)

	/*
		5. Receive Location
	*/
	ReceiveLocation(idRequest string) (bool, []ReceiveLocation, error)
}
