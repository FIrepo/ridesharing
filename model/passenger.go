package model

import (
	"errors"
	"fmt"
	"log"
)

/*
	Login
*/

func (myPassenger *Passenger) Login(username string, password string) (bool, bool, string) {

	var anyErr bool = false
	var isAuthenticated bool = false
	var tmpLoginClaim LoginCLaim

	stmt := Db.QueryRow("SELECT id,password FROM passenger WHERE email = ?", username)
	rowScan := stmt.Scan(&tmpLoginClaim.Id, &tmpLoginClaim.Password)

	if rowScan == nil {
		byteHash := []byte(tmpLoginClaim.Password)
		bytePassString := []byte(password)
		if IsLoginCredValid(byteHash, bytePassString) {
			isAuthenticated = true
		} else {
			isAuthenticated = false
		}
	} else {
		// ignore the error cause it statement db.queryRow
		anyErr = true
	}

	return !anyErr, isAuthenticated, tmpLoginClaim.Id

}

/*
	GetId
*/

func (myPassenger *Passenger) GetId() string {
	return myPassenger.Id
}

/*
	Connect
*/
func (myPassenger *Passenger) Connect() (bool, error) {
	var anyErr bool = false
	stmt, err := Db.Prepare("UPDATE passenger SET is_connected = '1' WHERE id=?;")
	if err != nil {
		log.Fatal(err)
		anyErr = true
	}
	_, err = stmt.Exec(myPassenger.Id)
	if err != nil {
		log.Fatal(err)
		anyErr = true
	}

	if !anyErr {
		// no error
		return true, err
	} else {
		return false, err
	}

}

/*
	Send Presence
*/
func (myPassenger *Passenger) SendPresence() (bool, error) {
	var anyErr bool = false
	stmt, err := Db.Prepare("UPDATE passenger SET is_visible = '1' WHERE id=?;")
	if err != nil {
		fmt.Println(err)
		anyErr = true
	}
	_, err = stmt.Exec(myPassenger.Id)
	if err != nil {
		fmt.Println(err)
		anyErr = true
	}

	if !anyErr {
		// no error
		return true, err
	} else {
		return false, err
	}
}

/*
	Send Request
*/
func (myPassenger *Passenger) SendRequest(idDriver string, lat string, lon string) (bool, error) {
	var anyError bool = false
	stmt, err := Db.Prepare("INSERT INTO request(id_passenger,lat,lon, id_driver) VALUES(?,?,?,?);")

	if err != nil {
		log.Fatal(err)
		anyError = true
	}

	r, err := stmt.Exec(myPassenger.Id, lat, lon, idDriver)

	if err != nil {
		log.Fatal(err)
		anyError = true
	}

	affectedRows, err := r.RowsAffected()

	if err != nil {
		anyError = true
		log.Fatal(err)
	}

	if affectedRows == 0 {
		err = errors.New("request not found")
	}

	if !anyError {
		return true, err
	} else {
		return false, err
	}

}

/*
	ReceiveRequest
*/

func (myPassenger *Passenger) ReceiveRequest() (bool, []PassengerRequest, error) {
	var (
		anyErr   bool = false
		request  PassengerRequest
		requests []PassengerRequest
	)
	stmt, err := Db.Query("SELECT driver.name,driver.email,request.lat,request.lon FROM request INNER JOIN driver ON request.id_driver = driver.id WHERE request.status = 1;")

	if err != nil {
		log.Fatal(err)
		anyErr = true
	}
	defer stmt.Close()

	count := 0
	for stmt.Next() {
		stmt.Scan(&request.DriverName, &request.DriverEmail, &request.Lat, &request.Lon)
		requests = append(requests, request)
		count++
	}

	if count == 0 {
		err = errors.New("No request found")
	}

	return !anyErr, requests, err
}

/*
	Receive Location
*/

func (myPassenger *Passenger) ReceiveLocation(idRequest string) (bool, []ReceiveLocation, error) {
	var (
		anyErr    bool = false
		location  ReceiveLocation
		locations []ReceiveLocation
	)
	stmt, err := Db.Query("SELECT driver.lat,driver.lon FROM request INNER JOIN driver ON request.id_driver = driver.id WHERE request.id = ?;", idRequest)

	if err != nil {
		anyErr = true
		log.Fatal(err)
	}
	defer stmt.Close()

	count := 0
	for stmt.Next() {
		stmt.Scan(&location.Lat, &location.Lon)
		locations = append(locations, location)
	}

	if count == 0 {
		err = errors.New("No request found")
	}

	return !anyErr, locations, err
}
