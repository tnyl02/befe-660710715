package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Ticket struct {
	ID     int    `json:"id"`
	Movie  string `json:"movie"`
	Seat   string `json:"seat"`
	Booked bool   `json:"booked"`
}

var tickets = []Ticket{
	{ID: 1, Movie: "Spider-Man", Seat: "A1", Booked: false},
	{ID: 2, Movie: "Spider-Man", Seat: "A2", Booked: false},
	{ID: 3, Movie: "Avatar 2", Seat: "B1", Booked: false},
	{ID: 4, Movie: "Avatar 2", Seat: "B2", Booked: true},
}

func getTickets(c *gin.Context) {
	movieQuery := c.Query("movie")

	filter := []Ticket{}
	for _, t := range tickets {
		if movieQuery == "" || strings.Contains(strings.ToLower(t.Movie), strings.ToLower(movieQuery)) {
			filter = append(filter, t)
		}
	}

	c.JSON(http.StatusOK, filter)
}

func bookTicket(c *gin.Context) {
	var req struct {
		ID int `json:"id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, t := range tickets {
		if t.ID == req.ID {
			if t.Booked {
				c.JSON(http.StatusConflict, gin.H{"message": "Ticket already booked"})
				return
			}
			tickets[i].Booked = true
			c.JSON(http.StatusOK, tickets[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Ticket not found"})
}


func main(){
	r := gin.Default()

	r.GET("/health", func(c*gin.Context)  {
		c.JSON(200, gin.H{"message" : "healthy"})
	})

	api := r.Group("/api/v1") 
	{
		api.GET("/tickets", getTickets)
	}

	r.Run(":8080")
}

