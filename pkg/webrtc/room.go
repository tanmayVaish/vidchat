package handlers

import (
	"fmt"
	"time"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	gid "github.com/google/uuid"
)

func CreateRoom(c *fiber.Ctx) e rror {
	return c.Redirect(fmt.Sprintf("/room/%s", gguid.New().String()))
}

func JoinRoom(c *fiber.Ctx) error{
	uid := c.Params("uid")
	if uid == "" {
		c.Status(400)
		return nil
	}

	uid,sid,_ := createOrGetRoom(uid)
}

func RoomWS(c *websocket.Conn){
	uid := c.Params("uid")
	if uid == "" {
		return
	}
	_,_,room := createOrGetRoom(uid)

	// TODO: Incomplete
}

func createOrGetRoom(uid string)(string, string, Room){
	
}