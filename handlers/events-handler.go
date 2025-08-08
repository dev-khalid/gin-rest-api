package handlers

import (
	"net/http"
	"time"

	"github.com/dev-khalid/gin-rest-api/models"
	"github.com/gin-gonic/gin"
)


func GetEvents(context *gin.Context) {
	allEvents := models.GetAllEvents()
	context.JSON(http.StatusOK, allEvents)
}

func CreateEvents(context *gin.Context) { 
	var event models.Event
	/**
	Raw http Style
	err := json.NewDecoder(context.Request.Body).Decode(&event)
	if err != nil {
		context.Error(errors.New("something went wrong"))
	}
	*/

	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, map[string]string {"message": "Could not parse request data", "errMessage": err.Error()})
		return
	}

	event.ID = int(time.Now().Unix())
	event.UserID = int(time.Now().Unix())

	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created!", "event": event})
}