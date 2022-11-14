package server.go

import (
	"flag"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/gofiber/websocket/v2"
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

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})
	
	app.use(logger.New())
	app.user(cors.New())

	app.Get("/", handlers.Welcome)
	
	app.Get("/room/create", handlers.CreateRoom)
	app.Get("/room/:id", handlers.JoinRoom)
	app.Get("/room/:id/ws", websocket.New(handlers.RoomWS, websocket.Config{
		HandshakeTimeout: 10 * time.Second,
	}))
	app.Get("/room/:id/chat", handlers.ChatRoom)
	app.Get("/room/id/chat/ws", websocket.New(handlers.ChatRoomWS))
	app.Get("/room/id/view/ws", websocket.New(handlers.ViewRoomWS))
	
	app.Get("/stream/:sid", handlers.Stream)
	app.Get("/stream/:sid/ws",)
	app.Get("/stream/:sid/chat/ws",)
	app.Get("/stream/:sid/view/ws",)
}