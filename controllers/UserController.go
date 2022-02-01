package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

	//"strings"
	"time"
	gonanoid "github.com/matoous/go-nanoid"
	Configuration "github.com/anik4good/fiber_boilerplate/config"
	"github.com/anik4good/fiber_boilerplate/models"
	"github.com/gofiber/fiber/v2"
)

var database *sql.DB

//Hello
func Hello(c *fiber.Ctx) error {
	return c.SendString("fiber")
}

func CreateUser(c *fiber.Ctx) error {

	//	data, _ := json.MarshalIndent(app.Stack(), "", "  ")
	//	fmt.Println(string(data))
	requestBody := c.Body()
	var email models.User
	json.Unmarshal(requestBody, &email)
	_, err := database.Exec(`INSERT INTO users(name, email,status) VALUES (?,?,?)`, email.Name, email.Email, email.Status)
	if err != nil {

		//	panic(err)

		fmt.Println("error creating user:", email.Name)
		json.NewEncoder(c).Encode("error creating user:")
		return nil
		//	json.NewEncoder(c).Encode("error creating user:")
	}

	json.NewEncoder(c).Encode("received Email: " + email.Email)
	return nil
}

//AddBook
func AddUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	Configuration.GormDBConn.Create(&user)
	log.Println("User Created successfully")
	return c.Status(200).JSON(user)
}

func SmsApi(c *fiber.Ctx) error   {

	data := new(models.Api_body)
	time.Sleep(1*time.Second)

	if err := c.BodyParser(data); err != nil {
		return c.Status(400).JSON(err.Error())
	}



	//Validation
	if len(data.Apikey) ==0  || len(data.MessageType) == 0 || len(data.Contacts) ==0 || len(data.Message) ==0 || len(data.Senderid) ==0{
		return c.Status(400).JSON("data cannot be empty")
	}


	fmt.Println(len(data.Contacts))
	//if len(data.Contacts) !=11 {
	//	return c.Status(400).JSON("Mobile No should be 11 digit")
	//}

	valid := GetValidPhoneNumber(data.Contacts)
	if valid == false {
		return c.Status(400).JSON("Invalid Phone Number")
	}

	id, _ := gonanoid.Nanoid(13)

	return c.Status(200).JSON("Your SMS is Submitted. ID: "+data.Senderid+"_"+id)
}






func GetValidPhoneNumber(number string) bool  {
	reg := regexp.MustCompilePOSIX("/[^0-9]/")
	msisdn := reg.ReplaceAllString(number,"")


	if len(msisdn) ==11 {
		if strings.HasPrefix(msisdn,"01") {
			log.Println("01")

		return  true
		}

		return false

	}

	if len(msisdn) ==13 {
		if strings.HasPrefix(msisdn,"8801") {

			return  true
		}

	}

	if len(msisdn) ==14 {
		if strings.HasPrefix(msisdn,"+8801") {

			return  true
		}

	}


	return false

}
