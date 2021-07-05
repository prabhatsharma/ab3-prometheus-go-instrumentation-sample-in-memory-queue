package queueapi

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	queueName := c.Param("queue")
	qm := QueueMetrics[queueName]

	if qm.TotalMessages == 0 {
		log.Printf("Queue %s is empty", queueName)
		c.JSON(500, gin.H{
			"total_messages": 0,
			"total_size":     0,
		})
		return
	}

	qm.RequestsReceived += 1
	qm.TotalMessages -= 1
	QueueMetrics[queueName] = qm

	FrontMessageInQueue := Messages[queueName][0:1][0]

	Messages[queueName] = Messages[queueName][1:len(Messages[queueName])]
	qm.TotalMessages -= 1

	PromTotalMessagesInQueue.WithLabelValues(queueName).Dec()
	PromRequestsReceived.WithLabelValues(queueName).Inc()

	c.JSON(200, FrontMessageInQueue)
}
