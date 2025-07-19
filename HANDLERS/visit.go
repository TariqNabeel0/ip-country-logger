package handlers

import (
	"encoding/json"
	"ip-country-logger/config"
	"ip-country-logger/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

	type GeoResponse struct {
		Country string `json:"country"`
		City string `json:"city"`
	}

	func PostVisit (c *gin.Context){
		var input models.VisitInput
		if err := c.ShouldBindJSON(&input); err != nil{
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

resp, err := http.Get("http://ip-api.com/json/" + input.IP)
if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get the geo info"})
	return
}
defer resp.Body.Close()

var geo GeoResponse
if err := json.NewDecoder(resp.Body).Decode(&geo); err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid response from geo API"})
	return
}

visit := models.Visit {
	IP: input.IP,
	Country: geo.Country,
	City: geo.City,
	WebsiteTag: input.WebsiteTag,
	Timestamp: time.Now(),
}

config.DB.Create(&visit)
c.JSON(http.StatusCreated, visit)

}

func GetVisit (c *gin.Context){
	var visits []models.Visit
	query := config.DB

	if tag := c.Query("tag"); tag != "" {
		query = query.Where("website_tag = ?", tag)
	}

if country := c.Query("country"); country != "" {
	query = query.Where("country = ?", country)
}

query.Find(&visits)
c.JSON(http.StatusOK, visits)

}

func GetSummary(c *gin.Context) {
	type Result struct {
		Value string
		Count int
	}

	var countries []Result
	config.DB.Model(&models.Visit{}).Select("country as value, count(*) as count").Group("country").Scan(&countries)

	var tags []Result
	config.DB.Model(&models.Visit{}).Select("website_tag as value, count(*) as count").Group("website_tag").Scan(&tags)

	c.JSON(http.StatusOK, gin.H{
		"visits_per_country": countries,
		"visits_per_tag":     tags,
	})
}