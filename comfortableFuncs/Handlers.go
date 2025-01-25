package comfortablefuncs

import (
	"gomodule/databases"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)



func GetHandler(c echo.Context) error {
	var messages []Message
	if err := databases.Db.Find(&messages).Error; err != nil {
		return BadRequest(c,"Could not find the messages")
	}
	return c.JSON(http.StatusOK, &messages)
}

func PostHandler(c echo.Context) error {
	var message Message
	if err := c.Bind(&message); err != nil {
		return BadRequest(c,"Could not add the message")
	}

	if err := databases.Db.Create(&message).Error; err != nil {
		return BadRequest(c,"Could not create the message")
	}
	return OK(c,"Message was successfully created")
}

func PatchHandler(c echo.Context) error {
	idParam := c.Param("ID")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return BadRequest(c,"Bad ID")
	}
	var Updatemessage Message
	if err := c.Bind(&Updatemessage); err != nil {
		return BadRequest(c,"Invalid input")
	}
	if err := databases.Db.Model(&Message{}).Where("id = ?", id).Update("text", Updatemessage.Text).Error; err != nil {
		return BadRequest(c,"Could not update the message")
	}
	return OK(c,"Message was successfully updated")
}

func DeleteHandler(c echo.Context) error {
	idParam := c.Param("ID")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return BadRequest(c,"Bad ID")
	}
	if err := databases.Db.Delete(&Message{}, id).Error; err != nil {
		return BadRequest(c,"Could not delete the message")
	}
	return OK(c,"Message was successfully deleted")
}