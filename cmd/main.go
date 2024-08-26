package main

import (
	// "stratus-core/database"
	"fmt"
	"stratus-core/utils"
	"time"
)

func main() {
	// err := database.InitMongoDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	successChan := make(chan bool)

	go utils.SendMail(
		successChan,
		"siddharthsinghrana11@gmail.com",
		"test project",
		"test site",
		"https://www.site.com",
		"501",
	)

	for {
		fmt.Println("Hello")
		time.Sleep(time.Second*5)
	}

	success:=<-successChan
	if success==true {
		fmt.Println("sent successfully")
	} else {
		fmt.Println("not sent") 
	}

}