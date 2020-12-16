package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	//Lets First Parce the data from the form!
	c.Request.ParseForm()

	//Now Lets store the Info In Varibials. (Dont Worry We Will Hash Them Later!)
	var email string = c.PostForm("email")
	var passwords string = c.PostForm("pas")

	//Now Lets Make a call to the databace to make sure we have a valed request.

	//Lets Make Varibals for all the values we need to connect
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "power7"
		dbname   = "login"
	)

	//Now we can cook up the connectiong string from the
	// That we stored in varibals

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//Now we Can Put or Connect string Into the sql.open() method.
	db, _ := sql.Open("postgres", psqlInfo)

	//Check That its on
	db.Ping()

	//Lets Makes Sure to Log Some important info!
	fmt.Println("😀 Request Made!!")
	fmt.Println("🔐 Password : ", passwords)
	fmt.Println("👥 Username : ", email)
	fmt.Print("")

}
