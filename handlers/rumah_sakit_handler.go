package handlers

import (
	"context"
	"net/http"
	"os"
	"strconv"
	"time"

	"rs-bandung-api/config"
	"rs-bandung-api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetRumahSakit(c *gin.Context) {
	db := config.MongoClient.Database(os.Getenv("MONGO_DB"))
	collection := db.Collection("rumah_sakit")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	var results []models.RumahSakit
	if err := cursor.All(ctx, &results); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"total":  len(results),
		"data":   results,
	})
}

func GetRumahSakitNear(c *gin.Context) {
	lat, _ := strconv.ParseFloat(c.Query("lat"), 64)
	lng, _ := strconv.ParseFloat(c.Query("lng"), 64)
	radius, _ := strconv.ParseInt(c.Query("radius"), 10, 64)

	pipeline := []bson.M{
		{
			"$geoNear": bson.M{
				"near": bson.M{
					"type":        "Point",
					"coordinates": []float64{lng, lat},
				},
				"distanceField": "distance",
				"maxDistance":   radius,
				"spherical":     true,
			},
		},
	}

	cursor, err := config.RumahSakitCollection.Aggregate(
		context.Background(),
		pipeline,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var results []bson.M
	if err := cursor.All(context.Background(), &results); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"total":  len(results),
		"data":   results,
	})
}
