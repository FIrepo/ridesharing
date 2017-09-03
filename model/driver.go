package model

import (
	"errors"
	"fmt"
	"log"
)

/*
	Driver Model Implement DriverInterface
*/

/*
	Login
*/

func (myDriver *Driver) Login(username string, password string) (bool, bool, string) {

	var anyErr bool = false
	var isAuthenticated bool = false
	var tmpLoginClaim LoginCLaim

	stmt := Db.QueryRow("SELECT id,password FROM driver WHERE email = ?", username)
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

func (myDriver *Driver) GetId() string {
	return myDriver.Id
}

/*
	Connect
*/
func (myDriver *Driver) Connect() (bool, error) {
	var anyErr bool = false
	stmt, err := Db.Prepare("UPDATE driver SET is_connected = '1' WHERE id=?;")
	if err != nil {
		log.Fatal(err)
		anyErr = true
	}
	_, err = stmt.Exec(myDriver.Id)
	fmt.Println("oblo" + myDriver.Id)
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
func (myDriver *Driver) SendPresence() (bool, error) {
	var anyErr bool = false
	stmt, err := Db.Prepare("UPDATE driver SET is_visible = '1' WHERE id=?;")
	if err != nil {
		log.Fatal(err)
		anyErr = true
	}
	_, err = stmt.Exec(myDriver.Id)
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
	ReceiveRequest
*/

func (myDriver *Driver) ReceiveRequest() (bool, []Request, error) {
	var (
		anyErr   bool = false
		request  Request
		requests []Request
	)
	stmt, err := Db.Query("SELECT passenger.name,passenger.email,request.lat,request.lon FROM request INNER JOIN passenger ON request.id_passenger = passenger.id WHERE request.status = 0 AND request.id_driver = ?;", myDriver.Id)

	if err != nil {
		log.Fatal(err)
		anyErr = true
	}
	defer stmt.Close()
	count := 0
	for stmt.Next() {
		stmt.Scan(&request.Name, &request.Email, &request.Lat, &request.Lon)
		requests = append(requests, request)
		count++
	}

	if count == 0 {
		anyErr = true
		err = errors.New("No Request found")
	}

	return !anyErr, requests, err
}

/*
	AcceptRequest
*/
func (myDriver *Driver) AcceptRequest(idRequest string) (bool, error) {
	var anyErr bool = false
	stmt, err := Db.Prepare("UPDATE request SET id_driver = ?, status = ? WHERE id = ? AND status = 0;")

	if err != nil {
		log.Fatal(err)
		anyErr = true
	}
	defer stmt.Close()

	r, err := stmt.Exec(myDriver.Id, 1, idRequest)

	if err != nil {
		log.Fatal(err)
		anyErr = true
	}
	defer stmt.Close()

	affectedRows, err := r.RowsAffected()

	if err != nil {
		anyErr = true
	}

	if affectedRows == 0 {
		anyErr = true
		err = errors.New("request not found")
	}
	return !anyErr, err

}

/*
	Send Location
*/
func (myDriver *Driver) SendLocation(lat string, lon string) (bool, error) {

	var anyErr bool = false
	stmt, err := Db.Prepare("UPDATE driver SET lat = ?, lon = ? WHERE id = ?;")

	if err != nil {
		anyErr = true
		log.Fatal(err)
	}

	_, err = stmt.Exec(lat, lon, myDriver.Id)

	if err != nil {
		anyErr = true
		log.Fatal(err)
	}

	return !anyErr, err

}

/*
	Start trip
*/

func (myDriver *Driver) StartTrip(idRequest string) (bool, error) {

	var anyError bool = false
	stmt, err := Db.Prepare("UPDATE request SET status = ? WHERE id = ? AND status = 1;")
	if err != nil {
		anyError = true
		log.Fatal(err)
	}
	defer stmt.Close()

	r, err := stmt.Exec(3, idRequest)

	affectedRows, err := r.RowsAffected()

	if err != nil {
		anyError = true
		log.Fatal(err)
	}

	if affectedRows == 0 {
		err = errors.New("request not found")
	}

	return !anyError, err

}

func (myDriver *Driver) EndTrip(idRequest string, distance string, time string) (bool, error) {

	var anyError bool = false
	stmt, err := Db.Prepare("UPDATE request SET status = ?, distance = ?, time = ? WHERE id = ? AND status = 3;")
	if err != nil {
		anyError = true
	}
	defer stmt.Close()

	r, err := stmt.Exec(4, distance, time, idRequest)

	affectedRows, err := r.RowsAffected()

	if err != nil {
		anyError = true
		log.Fatal(err)
	}

	if affectedRows == 0 {
		err = errors.New("request not found")
	}

	return !anyError, err

}
