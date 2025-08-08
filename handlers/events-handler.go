package handlers

import (
	"net/http"
	"strconv"

	"github.com/dev-khalid/gin-rest-api/config"
	"github.com/dev-khalid/gin-rest-api/models"
	"github.com/dev-khalid/gin-rest-api/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var eventRepo repositories.EventRepository

func init() {
	eventRepo = repositories.NewEventRepository(config.DB)
}

func GetEvents(c *gin.Context) {
	events, err := eventRepo.GetEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch events", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, events)
}

func GetEvent(c *gin.Context) {
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}
	event, err := eventRepo.GetEventByID(uint(id64))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch event", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, event)
}

func CreateEvents(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "errMessage": err.Error()})
		return
	}

	if err := eventRepo.CreateEvent(&event); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create event", "error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, event)
}

func UpdateEvent(c *gin.Context) {
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}
	var payload models.Event
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data", "errMessage": err.Error()})
		return
	}
	updated, err := eventRepo.UpdateEvent(uint(id64), &payload)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update event", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event Updated!", "event": updated})
}

func DeleteEvent(c *gin.Context) {
	idParam := c.Param("id")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}
	if err := eventRepo.DeleteEvent(uint(id64)); err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete event", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Event Deleted!"})
}
