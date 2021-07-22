package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET},
	}))

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", getLocation)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Structs
type location struct {
	Country    string  `json:"country"`
	Region     string  `json:"region"`
	City       string  `json:"city"`
	Lat        float64 `json:"lat"`
	Long       float64 `json:"lng"`
	PostalCode string  `json:"postalCode"`
	Timezone   string  `json:"timezone"`
	GeonameID  int     `json:"geonameId"`
}

type as struct {
	ASN    int    `json:"asn"`
	Name   string `json:"name"`
	Route  string `json:"route"`
	Domain string `json:"domain"`
	Type   string `json:"type"`
}

type proxy struct {
	Proxy bool `json:"proxy"`
	VPN   bool `json:"vpn"`
	Tor   bool `json:"tor"`
}

type geo struct {
	IP       string   `json:"ip"`
	Location location `json:"location"`
	Domains  []string `json:"domains"`
	As       as       `json:"as"`
	ISP      string   `json:"isp"`
	Proxy    proxy    `json:"proxy"`
}

// Handler
func getLocation(c echo.Context) error {
	apiKey := os.Getenv("API_KEY")
	url := fmt.Sprintf("https://geo.ipify.org/api/v1?apiKey=%s", apiKey)

	// Request
	res, getErr := http.Get(url)
	if getErr != nil {
		log.Fatal(getErr)
	}

	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// Fill the data with the body from the JSON
	var g geo

	jsonErr := json.Unmarshal(body, &g)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return c.JSON(http.StatusOK, g)
}
