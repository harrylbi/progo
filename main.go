package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// Define error messages
	internalServerError = "Internal Server Error"
	notFound            = "Not Found"

	// Declare a histogram metric to track HTTP request durations
	httpDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Histogram of response durations for HTTP requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "status"},
	)
)

func init() {
	// Register the metric with Prometheus
	prometheus.MustRegister(httpDuration)
}

func main() {
	r := gin.Default()

	// Middleware to track HTTP request durations
	r.Use(func(c *gin.Context) {
		start := time.Now()

		// Continue processing the request
		c.Next()

		// Measure the duration of the HTTP request
		duration := time.Since(start).Seconds()
		httpDuration.WithLabelValues(c.Request.Method, fmt.Sprintf("%d", c.Writer.Status())).Observe(duration)
	})

	// Prometheus metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Define your routes
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Prometheus!")
	})

	// Route for "/get-user"
	r.GET("/get-user", func(c *gin.Context) {
		param := c.DefaultQuery("param", "") // Get the query parameter "param"

		if param == "error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": internalServerError,
			})
			return
		}

		if param == "not-found" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": notFound,
			})
			return
		}

		c.String(http.StatusOK, "Success Get Users")
	})

	// Route for "/get-role"
	r.GET("/get-role", func(c *gin.Context) {
		param := c.DefaultQuery("param", "") // Get the query parameter "param"

		if param == "error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": internalServerError,
			})
			return
		}

		if param == "not-found" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": notFound,
			})
			return
		}

		c.String(http.StatusOK, "Success Get Roles")
	})

	// Route for "/get-level"
	r.GET("/get-level", func(c *gin.Context) {
		param := c.DefaultQuery("param", "") // Get the query parameter "param"

		if param == "error" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": internalServerError,
			})
			return
		}

		if param == "not-found" {
			c.JSON(http.StatusNotFound, gin.H{
				"message": notFound,
			})
			return
		}

		c.String(http.StatusOK, "Success Get Levels")
	})

	// Start the Gin server
	r.Run("127.0.0.1:8080")
}
