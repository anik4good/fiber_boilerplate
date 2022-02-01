package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

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
