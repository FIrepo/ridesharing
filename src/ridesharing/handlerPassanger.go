package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func passangerConnect(c *gin.Context) {
	if loggedPassenger.connect() {
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

func passangerPresence(c *gin.Context) {
	if loggedPassenger.sendPresence() {
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

func passangerRequest(c *gin.Context) {
	var (
		sendRequestInput SendRequest
	)

	if c.BindJSON(&sendRequestInput) == nil {
		// json body valid
		if loggedPassenger.sendRequest(sendRequestInput.Lat, sendRequestInput.Lon) {
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
		fmt.Println(sendRequestInput.Lat)
		c.JSON(200, gin.H{
			"status": "-1",
			"msg":    "failed",
		})
	}

}

func passengerReceiveRequest(c *gin.Context) {
	err, request := loggedPassenger.receiveRequest()
	if err {
		c.JSON(200, request)
	} else {
		c.JSON(200, gin.H{
			"status": "-1",
			"msg":    "failed",
		})
	}
}

func passengerReceiveLocation(c *gin.Context) {
	err, request := loggedPassenger.receiveLocation(c.Param("id"))
	fmt.Println(request)
	if err {
		c.JSON(200, request)
	} else {
		c.JSON(200, gin.H{
			"status": "-1",
			"msg":    "failed",
		})
	}
}

/*
	Connect
*/
func (myPassenger *Passenger) connect() bool {
	var anyErr bool = false
	stmt, err := db.Prepare("UPDATE passenger SET is_connected = '1' WHERE id=?;")
	if err != nil {
		fmt.Println(err)
		anyErr = true
	}
	_, err = stmt.Exec(myPassenger.id)
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
func (myPassenger *Passenger) sendPresence() bool {
	var anyErr bool = false
	stmt, err := db.Prepare("UPDATE passenger SET is_visible = '1' WHERE id=?;")
	if err != nil {
		fmt.Println(err)
		anyErr = true
	}
	_, err = stmt.Exec(myPassenger.id)
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
	Send Request
*/
func (myPassenger *Passenger) sendRequest(lat string, lon string) bool {
	var anyError bool = false
	stmt, err := db.Prepare("INSERT INTO request(id_passenger,lat,lon) VALUES(?,?,?);")

	if err != nil {
		anyError = true
	}

	_, err = stmt.Exec(myPassenger.id, lat, lon)

	if err != nil {
		anyError = true
	}

	if !anyError {
		return true
	} else {
		return false
	}

}

/*
	ReceiveRequest
*/

func (myPassenger *Passenger) receiveRequest() (bool, []PassengerRequest) {
	var (
		anyErr   bool = false
		request  PassengerRequest
		requests []PassengerRequest
	)
	stmt, err := db.Query("SELECT driver.name,driver.email,request.lat,request.lon FROM request INNER JOIN driver ON request.id_driver = driver.id WHERE request.status = 1;")

	if err != nil {
		anyErr = true
	}
	defer stmt.Close()

	for stmt.Next() {
		stmt.Scan(&request.DriverName, &request.DriverEmail, &request.Lat, &request.Lon)
		requests = append(requests, request)
	}
	return !anyErr, requests
}

/*
	Receive Location
*/

func (myPassenger *Passenger) receiveLocation(idRequest string) (bool, []ReceiveLocation) {
	var (
		anyErr    bool = false
		location  ReceiveLocation
		locations []ReceiveLocation
	)
	stmt, err := db.Query("SELECT driver.lat,driver.lon FROM request INNER JOIN driver ON request.id_driver = driver.id WHERE request.id = ?;", idRequest)

	if err != nil {
		anyErr = true
	}
	defer stmt.Close()

	for stmt.Next() {
		stmt.Scan(&location.Lat, &location.Lon)
		locations = append(locations, location)
	}
	return !anyErr, locations
}
