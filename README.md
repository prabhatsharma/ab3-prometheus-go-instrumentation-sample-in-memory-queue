# Simple In-memory Queue implementation in Go to demonstrate prometheus metrics instrumentation



HTTP GET /queue/:queue - Receive the last message from the queue

HTTP GET /metrics - prometheus metrics

### Send message to the queue

HTTP POST /queue/:queue - send a message

{
    "body": "message for the body. Can be any type (string, integer, object). "
}


test