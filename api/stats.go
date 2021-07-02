package queueapi

import "github.com/gin-gonic/gin"

func Stats(c *gin.Context) {
	queueName := c.Param("queue")
	qm := QueueMetrics[queueName]
	qm.RequestsReceived += 1
	QueueMetrics[queueName] = qm

	PromRequestsReceived.WithLabelValues(queueName).Inc()

	c.JSON(201, QueueMetrics[queueName])
}
