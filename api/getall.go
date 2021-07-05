package queueapi

import "github.com/gin-gonic/gin"

func GetAll(c *gin.Context) {

	PromRequestsReceived.WithLabelValues("all_queues").Inc()

	c.JSON(200, gin.H{
		"messages":     Messages,
		"QueueMetrics": QueueMetrics,
	})
}
