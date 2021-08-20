package main

import (
	"net/http"
    "time"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-contrib/prometheus"
    customProme "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	appVersion string
	version = promauto.NewGauge(customProme.GaugeOpts{
		Name: "version",
		Help: "Versi aplikasi",
		ConstLabels: map[string]string{
			"version": "V.1.0.0-beta",
		},
	})

)


func recordMetrics() {
        go func() {
                for {
                        opsProcessed.Inc()
                        time.Sleep(2 * time.Second)
                }
        }()
 }

var (
        opsProcessed = promauto.NewCounter(customProme.CounterOpts{
                Name: "myapp_processed_ops_total",
                Help: "The total number of processed events",
        })
)





func main() {
	e := echo.New()
	recordMetrics()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})	
	e.GET("/users/:id", getUser)
	e.GET("/testing/:id", testing)




	  // Enable metrics middleware
	p := prometheus.NewPrometheus("" ,nil)
	p.Use(e)
	e.Logger.Fatal(e.Start(":1323"))
}


//testing function
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
  return c.String(http.StatusOK, id)
}

func testing(c echo.Context) error {
	id := c.Param("id")
  return c.String(http.StatusOK, id)
}

