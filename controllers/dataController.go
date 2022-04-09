package controllers

import (
	"log"

	"Wedding.com/database"
	"Wedding.com/models"
	"go.mongodb.org/mongo-driver/bson"

	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func PostComment(c *fiber.Ctx) error {
	log.Print("==================== Hit Endpoint Post /api/comment/ ==========================")

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

func GetAllData(c *fiber.Ctx) error {
	log.Print("====================== Hit Endpoint Get /api/data/:user_data =======================")

	client_data := c.Params("user_data")
	var results []models.PostData

	data, _ := database.DB.Collection(client_data).Find(context.Background(), bson.D{})

	data.All(context.TODO(), &results)

	log.Print("Get All Data")
	log.Print(results)

	return c.JSON(results)
}
