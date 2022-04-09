package controllers

import (
	"log"

	"Wedding.com/database"
	"Wedding.com/models"

	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func PostComment(c *fiber.Ctx) error {
	var postData fiber.Map

	if err := c.BodyParser(&postData); err != nil {
		return err
	}

	client_data := postData["user_data"].(string)

	attendace, _ := strconv.Atoi(postData["attendance"].(string))
	commentPost := models.PostData{
		Name:       postData["name"].(string),
		Attendance: attendace,
		Comment:    postData["wishes"].(string),
		CreatedAt:  time.Now(),
	}

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	_, errdb := database.DB.Collection(client_data).InsertOne(ctx, commentPost)

	if errdb != nil {
		return errdb
	}
	log.Println("==========Success Insert to Database " + client_data + " with Data ============")
	log.Println(commentPost)

	return c.JSON(postData)
}
