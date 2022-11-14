package server.go

import (
	"flag",
	"os",
	"time"
)

var (
	addr = flag.String("addr,":"", os.Getenv("PORT"),"")
	cert = flag.String("cert", "", "")
	key = flag.String("key", "", "")
)

func Run() error {
	flag.Parse()
	
	if *addr == ":" {
		*addr = ":8080"
	}

	app.Get("/", handlers.Welcome)
	
	app.Get("/room/create", handlers.CreateRoom)
	app.Get("/room/:id", handlers.JoinRoom)
	app.Get("/room/:id/ws")
	app.Get("/room/:id/chat", handlers.ChatRoom)
	app.Get("/room/id/chat/ws", websocket.New(handlers.ChatRoomWS))
	app.Get("/room/id/view/ws", websocket.New(handlers.ViewRoomWS))
	
	app.Get("/stream/:sid", handlers.Stream)
	app.Get("/stream/:sid/ws",)
	app.Get("/stream/:sid/chat/ws",)
	app.Get("/stream/:sid/view/ws",)
}