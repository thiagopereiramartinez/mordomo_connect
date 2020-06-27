package sensors

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/gofiber/fiber"
	"log"
	"time"
)

func Sensors(c *fiber.Ctx) {

	// Use the application default credentials
	ctx := context.Background()
	conf := &firebase.Config{ProjectID: "masterdeveloper-mordomo-hub"}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	// Ler do Firestore
	snap, _ := client.Collection("esp32").Doc("state").Get(ctx)
	doc := snap.Data()
	temp, humidity, timestamp := doc["temp"], doc["humidity"].(float64), doc["timestamp"].(time.Time)

	// Devolver dados
	c.JSON(fiber.Map{
		"temp":      temp,
		"humidity":  humidity,
		"timestamp": timestamp.Unix(),
	})

}
