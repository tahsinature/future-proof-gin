package controllers

import (
	"fmt"
	"math"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tahsinature/future-proof-gin/pkg/db"
)

type Ping struct{}

var startTime = time.Now()

func (Ping) Response(c *gin.Context) {
	dbStatus := "OK"
	if err := db.GetDB().Exec("SELECT 1").Error; err != nil {
		dbStatus = "ERROR"
	}

	redisStatus := "OK"
	if err := db.RedisClient.Ping().Err(); err != nil {
		redisStatus = "ERROR"
	}

	uptime := time.Since(startTime)

	c.JSON(200, map[string]interface{}{
		"app":    "OK",
		"db":     dbStatus,
		"redis":  redisStatus,
		"uptime": fmt.Sprintf("%d hour(s) %d min(s) %d sec(s)", int(math.Trunc(uptime.Hours())), int(math.Trunc(uptime.Minutes())), int(math.Trunc(uptime.Seconds()))),
	})
}
