package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Message struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var messages = make(map[int]Message)
var nextID = 1

func GetHandler(c echo.Context) error {
	var msgSlice []Message
	for _, msg := range messages {
		msgSlice = append(msgSlice, msg)

	}
	return c.JSON(http.StatusOK, &msgSlice)
}

func PostHandler(c echo.Context) error {
	var message Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not add the message",
		})
	}

	// to increment user id
	message.ID = nextID
	nextID++

	messages[message.ID] = message
	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message was added successfully",
	})
}

func PatchHandler(c echo.Context) error {
	// to convert id from string to int
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "ID is not correct",
		})
	}

	var updatedMessage Message
	if err := c.Bind(&updatedMessage); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Could not update the message",
		})
	}

	if _, exists := messages[id]; !exists {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Message was not found",
		})
	}

	updatedMessage.ID = id
	messages[id] = updatedMessage

	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message was updated successfully",
	})
}

func DeleteHandler(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "ID is not correct",
		})
	}

	if _, exists := messages[id]; !exists {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Message was not found",
		})
	}

	delete(messages, id)
	return c.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Message was deleted successfully",
	})
}

func main() {
	e := echo.New()
	e.GET("/message", GetHandler)
	e.POST("/message", PostHandler)
	e.PATCH("/message/:id", PatchHandler)
	e.DELETE("/message/:id", DeleteHandler)

	err := e.Start(":8080")
	if err != nil {
		return
	}
}
