package queueapi

import "github.com/prometheus/client_golang/prometheus"

var (
	PromTotalMessagesInQueue *prometheus.GaugeVec
	PromRequestsReceived     *prometheus.CounterVec
	PromMessageID            *prometheus.CounterVec
)

func MtericsStartUp() {
	setupTotalMessages()
	setupRequestsReceived()
	setupMessageID()

	prometheus.MustRegister(*PromTotalMessagesInQueue)
	prometheus.MustRegister(PromRequestsReceived)
	prometheus.MustRegister(PromMessageID)

}

func setupTotalMessages() {
	opts := prometheus.GaugeOpts{
		Namespace: "octank",
		Subsystem: "messages",
		Name:      "in_queue",
		Help:      "How many messages are in the queue.",
	}

	labelNames := []string{"queue"}
	PromTotalMessagesInQueue = prometheus.NewGaugeVec(opts, labelNames)
}

func setupRequestsReceived() {
	PromRequestsReceived = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "octank_queue_requests_received_total",
			Help: "How many requests have been received, partitioned by queue name.",
		},
		[]string{"queue"},
	)
}

func setupMessageID() {
	PromMessageID = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "octank_message_ids",
			Help: "Highest message ID.",
		},
		[]string{"queue"},
	)
}
