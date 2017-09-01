package main

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
	connect()
	/*
		2.sendPresence()
		Make Driver visible on the map (Presence)
	*/
	sendPresence(lat string, lon string)
	/*
		3. Receive Request
	*/
	receiveRequest()
	/*
		4. Accept Request
	*/
	acceptRequest(idRequest string)
	/*
		5. Send Location
	*/
	sendLocation(lat string, lon string)
	/*
		6. Start Trip
	*/
	startTrip(idRequest string)
	/*
		7. End Trip
	*/
	endTrip(idRequest string, distance string, time string)
}

type PassangerInterface interface {
	/*
		1. Connect(email,password)
		[+] email : Registered passenger email
		[+] password : Bcrypt hash
	*/
	connect()
	/*
		2. Make passenger visible on the map (Presence)
	*/
	sendPresence(lat string, lon string)
	/*
		3. Send Request
	*/
	sendRequest(lat string, lon string)
	/*
		4. Receive Request
	*/
	receiveRequest()
	/*
		5. Receive Location
	*/
	receiveLocation(idRequest string)
}

/**
	Legend
	Request Status
	[0] = Available
	[1] = Accepted
	[2] = Canceled
	[3] = Trip Start
	[4] = Trip Finish
**/
