package main

import (
	"github.com/gin-gonic/gin"
	queueapi "github.com/prabhatsharma/ab3/api"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func main() {
	queueapi.QueueMetrics = make(map[string]queueapi.QMetric)
	queueapi.Messages = make(map[string][]queueapi.Message)

	queueapi.MtericsStartUp()

	r := gin.Default()

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	r.POST("/queue/:queue", queueapi.Post)
	r.GET("/queue/:queue", queueapi.Get)
	r.GET("/queue/:queue/stats", queueapi.Stats)
	// r.GET("/metrics", prometheusHandler()) // this is being delegated to ginprometheus
	r.GET("/queue", queueapi.GetAll)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
