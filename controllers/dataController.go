package controllers

import (
	"Wedding.com/database"
	"Wedding.com/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"context"
	"strconv"
	"time"

	Logger "Wedding.com/log"
	"github.com/gofiber/fiber/v2"
)

func PostComment(c *fiber.Ctx) error {
	Logger.CommonLog.Print("==================== Hit Endpoint Post /api/comment/ ==========================")

	var postData fiber.Map

	if err := c.BodyParser(&postData); err != nil {
		return err
	}

	client_data := postData["user_data"].(string)

	attendace, _ := strconv.Atoi(postData["attendance"].(string))
	total_guest, _ := strconv.Atoi(postData["total_guest"].(string))
	commentPost := models.PostData{
		Name:       postData["name"].(string),
		Attendance: attendace,
		Comment:    postData["wishes"].(string),
		TotalGuest: total_guest,
		CreatedAt:  time.Now(),
	}

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	_, errdb := database.DB.Collection(client_data).InsertOne(ctx, commentPost)

	if errdb != nil {
		Logger.ErrorLog.Println(errdb)
		return errdb
	}
	Logger.CommonLog.Println("==========Success Insert to Database " + client_data + " with Data ============")
	Logger.CommonLog.Println(commentPost)

	return c.JSON(postData)
}

func GetAllData(c *fiber.Ctx) error {
	Logger.CommonLog.Print("====================== Hit Endpoint Get /api/data/:user_data =======================")

	client_data := c.Params("user_data")
	var results []models.PostData

	data, _ := database.DB.Collection(client_data).Find(context.Background(), bson.D{})

	count, _ := database.DB.Collection(client_data).CountDocuments(context.Background(), bson.D{})

	groupStage := bson.D{
		{
			"$group", bson.D{
				{"_id", ""},
				{"attendance", bson.D{{"$sum", "$attendance"}}},
				{"total_guest", bson.D{{"$sum", "$totalguest"}}},
			},
		},
	}
	opts := options.Aggregate().SetMaxTime(2 * time.Second)

	cursor, err := database.DB.Collection(client_data).Aggregate(context.TODO(), mongo.Pipeline{groupStage}, opts)
	var attendanceCount []bson.M

	if err = cursor.All(context.TODO(), &attendanceCount); err != nil {
		Logger.ErrorLog.Fatal(err)
	}
	coming_data := attendanceCount[len(attendanceCount)-1]["attendance"]
	total_guest := attendanceCount[len(attendanceCount)-1]["total_guest"]

	data.All(context.TODO(), &results)

	Logger.CommonLog.Print("Get All Data")
	Logger.CommonLog.Print(results)

	return c.JSON(fiber.Map{
		"attendace_come":      coming_data,
		"attendance_not_come": int32(count) - coming_data.(int32),
		"repondense":          count,
		"comment":             results,
		"total_guest":         total_guest,
	})
}
