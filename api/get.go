package queueapi

import "github.com/gin-gonic/gin"

func Get(c *gin.Context) {
	queueName := c.Param("queue")
	qm := QueueMetrics[queueName]

	qm.RequestsReceived += 1
	qm.TotalMessages -= 1
	QueueMetrics[queueName] = qm

	FrontMessageInQueue := Messages[queueName][0:1][0]

	Messages[queueName] = Messages[queueName][1:len(Messages[queueName])]
	qm.TotalMessages -= 1

	PromTotalMessagesInQueue.WithLabelValues(queueName).Dec()
	PromRequestsReceived.WithLabelValues(queueName).Inc()

	c.JSON(201, FrontMessageInQueue)
}
