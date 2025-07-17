package main

import (
	"backend/api"
	"backend/db"
	"backend/server"
	"context"
	"log"
	"os"

	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("env/data.env") // Cargar las variables de entorno desde el archivo .env
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}

	// Conectar a la base de datos CockroachDB
	// Leer la variable de conexión
	dsn := os.Getenv("DATABASE_URL")
	ctx := context.Background()        // Crear un contexto para la conexión
	conn, err := pgx.Connect(ctx, dsn) // Conectar a la base de datos usando el DSN
	// Manejo de errores al conectar a la base de datos
	defer func() {
		if err := conn.Close(context.Background()); err != nil { // Asegurarse de cerrar la conexión al final
			// Si hay un error al cerrar la conexión, imprimir un mensaje de error
			log.Fatal("failed to close connection", err)
		}
	}()
	if err != nil {
		log.Fatal("failed to connect database", err) // Si hay un error al conectar, imprimir un mensaje de error y salir del programa
	}

	// Set up table
	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return db.InitTable(context.Background(), tx)
	})
	if err != nil {
		log.Fatalf("error initializing table: %v", err)
	}

	items := api.ApiGetItems() // Obtener los items de la API
	if len(items) == 0 {
		log.Println("No items found in the response.")
		return
	}

	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return db.InsertRows(context.Background(), tx, items)
	})
	if err != nil {
		log.Fatalf("error inserting rows: %v", err) // Manejo de errores al insertar los datos en la base de datos
	}
	log.Println("New rows created.")

	// Start the server
	server.StartApiServer(conn) // Iniciar el servidor con la conexión a la base de datos
}
