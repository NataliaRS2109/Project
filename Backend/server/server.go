package server

import (
	"backend/db"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
)

func StartApiServer(conn *pgx.Conn) {
	app := fiber.New()

	app.Get("/items", func(c *fiber.Ctx) error {
		return db.PrintItems(conn, c) // <- aquí sí usas la versión que retorna JSON
	})

	log.Fatal(app.Listen(":3000")) // Iniciar el servidor en el puerto 3000
	log.Println("Server started on port 3000")
}
