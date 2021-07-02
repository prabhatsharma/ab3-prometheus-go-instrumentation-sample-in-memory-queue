package queueapi

var (
	Messages     map[string][]Message
	QueueMetrics map[string]QMetric
)

type Message struct {
	ID   uint64      `json:"id,omitempty"`
	Body interface{} `json:"body"`
}

type QMetric struct {
	TotalMessages    uint64 `json:"message_count"`
	MessageID        uint64 `json:"message_id"`
	RequestsReceived uint64 `json:"request_count"`
}
