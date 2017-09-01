package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func driverConnect(c *gin.Context) {
	loggedDriver.sendPresence()
	c.JSON(200, gin.H{
		"status": "1",
		"msg":    "success",
	})
}

func driverPresence(c *gin.Context) {
	loggedDriver.connect()
	c.JSON(200, gin.H{
		"status": "1",
		"msg":    "success",
	})
}

func driverReceiveRequest(c *gin.Context) {
	err, request := loggedDriver.receiveRequest()
	if err {
		c.JSON(200, request)
	} else {
		c.JSON(200, gin.H{
			"status": "-1",
			"msg":    "failed",
		})
	}
}

func driverAcceptRequest(c *gin.Context) {

	var (
		acceptRequestInput AcceptRequest
	)
	err := c.BindJSON(&acceptRequestInput)
	if err == nil {

		if loggedDriver.AcceptRequest(acceptRequestInput.IdRequest) {
			c.JSON(200, gin.H{
				"status": "1",
				"msg":    "success",
			})
		} else {
			c.JSON(200, gin.H{
				"status": "-1",
				"msg":    "failed",
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

func driverSendLocation(c *gin.Context) {

	var sendLocationInput SendLocation
	err := c.BindJSON(&sendLocationInput)
	if err == nil {
		if loggedDriver.sendLocation(sendLocationInput.Lat, sendLocationInput.Lon) {
			c.JSON(200, gin.H{
				"status": "1",
				"msg":    "success",
			})
		} else {
			c.JSON(200, gin.H{
				"status": "-1",
				"msg":    "failed",
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

func driverStartTrip(c *gin.Context) {

	if loggedDriver.startTrip(c.Param("id")) {
		c.JSON(200, gin.H{
			"status": "1",
			"msg":    "success",
		})
	} else {
		c.JSON(200, gin.H{
			"status": "-1",
			"msg":    "failed",
		})
	}

}

func driverEndTrip(c *gin.Context) {

	var endTripInput EndTrip
	err := c.BindJSON(&endTripInput)
	if err == nil {
		if loggedDriver.endTrip(c.Param("id"), endTripInput.Distance, endTripInput.Time) {
			c.JSON(200, gin.H{
				"status": "1",
				"msg":    "success",
			})
		} else {
			c.JSON(200, gin.H{
				"status": "-1",
				"msg":    "failed",
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

/*
	Connect
*/
func (myDriver *Driver) connect() bool {
	var anyErr bool = false
	stmt, err := db.Prepare("UPDATE driver SET is_connected = '1' WHERE id=?;")
	if err != nil {
		fmt.Println(err)
		anyErr = true
	}
	_, err = stmt.Exec(myDriver.id)
	if err != nil {
		fmt.Println(err)
		anyErr = true
	}

	if !anyErr {
		// no error
		return true
	} else {
		return false
	}

}

/*
	Send Presence
*/
func (myDriver *Driver) sendPresence() bool {
	var anyErr bool = false
	stmt, err := db.Prepare("UPDATE driver SET is_visible = '1' WHERE id=?;")
	if err != nil {
		fmt.Println(err)
		anyErr = true
	}
	_, err = stmt.Exec(myDriver.id)
	if err != nil {
		fmt.Println(err)
		anyErr = true
	}

	if !anyErr {
		// no error
		return true
	} else {
		return false
	}
}

/*
	ReceiveRequest
*/

func (myDriver *Driver) receiveRequest() (bool, []Request) {
	var (
		anyErr   bool = false
		request  Request
		requests []Request
	)
	stmt, err := db.Query("SELECT passenger.name,passenger.email,request.lat,request.lon FROM request INNER JOIN passenger ON request.id_passenger = passenger.id WHERE request.status = 0;")

	if err != nil {
		anyErr = true
	}
	defer stmt.Close()

	for stmt.Next() {
		stmt.Scan(&request.Name, &request.Email, &request.Lat, &request.Lon)
		requests = append(requests, request)
	}
	return !anyErr, requests
}

/*
	AcceptRequest
*/
func (myDriver *Driver) AcceptRequest(idRequest string) bool {
	var anyErr bool = false
	stmt, err := db.Prepare("UPDATE request SET id_driver = ?, status = ? WHERE id = ? AND status = 0;")

	if err != nil {
		anyErr = true
	}
	defer stmt.Close()

	_, err = stmt.Exec(myDriver.id, 1, idRequest)

	if err != nil {
		anyErr = true
	}

	return !anyErr

}

/*
	Send Location
*/
func (myDriver *Driver) sendLocation(lat string, lon string) bool {

	var anyErr bool = false
	stmt, err := db.Prepare("UPDATE driver SET lat = ?, lon = ? WHERE id = ?;")

	if err != nil {
		anyErr = true
	}

	_, err = stmt.Exec(lat, lon, myDriver.id)

	if err != nil {
		anyErr = true
	}

	return !anyErr

}

/*
	Start trip
*/

func (myDriver *Driver) startTrip(idRequest string) bool {

	var anyError bool = false
	stmt, err := db.Prepare("UPDATE request SET status = ? WHERE id = ? AND status = 1;")
	if err != nil {
		anyError = true
	}
	defer stmt.Close()

	_, err = stmt.Exec(3, idRequest)

	return !anyError

}

func (myDriver *Driver) endTrip(idRequest string, distance string, time string) bool {

	var anyError bool = false
	stmt, err := db.Prepare("UPDATE request SET status = ?, distance = ?, time = ? WHERE id = ? AND status = 3;")
	if err != nil {
		anyError = true
	}
	defer stmt.Close()

	_, err = stmt.Exec(4, distance, time, idRequest)

	return !anyError

}
