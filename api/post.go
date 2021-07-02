package queueapi

import (
	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {
	queueName := c.Param("queue")

	qm := QueueMetrics[queueName]

	qm.MessageID += 1
	qm.RequestsReceived += 1
	qm.TotalMessages += 1

	QueueMetrics[queueName] = qm

	var data Message

	data.ID = qm.MessageID

	c.Bind(&data)

	Messages[queueName] = append(Messages[queueName], data)

	PromTotalMessagesInQueue.WithLabelValues(queueName).Add(1)
	PromRequestsReceived.WithLabelValues(queueName).Inc()
	PromMessageID.WithLabelValues(queueName).Inc()

	c.JSON(201, gin.H{
		"result":     "Message queued",
		"message_id": data.ID,
	})
}
