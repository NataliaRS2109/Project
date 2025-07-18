package main

import (
	"backend/algorithm"
	"backend/api"
	"backend/db"
	"backend/server"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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
	ctx := context.Background() // Crear un contexto para la conexión
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal("failed to connect to database pool", err)
	}
	defer pool.Close()

	// conn, err := pgx.Connect(ctx, dsn) // Conectar a la base de datos usando el DSN
	// // Manejo de errores al conectar a la base de datos
	// defer func() {
	// 	if err := conn.Close(context.Background()); err != nil { // Asegurarse de cerrar la conexión al final
	// 		// Si hay un error al cerrar la conexión, imprimir un mensaje de error
	// 		log.Fatal("failed to close connection", err)
	// 	}
	// }()
	// if err != nil {
	// 	log.Fatal("failed to connect database", err) // Si hay un error al conectar, imprimir un mensaje de error y salir del programa
	// }

	// Set up table
	// Iniciar una transacción desde el pool
	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		log.Fatalf("Failed to begin transaction: %v", err)
	}
	defer tx.Rollback(ctx) // Por seguridad, hacer rollback si falla algo

	// Llamar InitTable usando la transacción
	if err := db.InitTable(ctx, tx); err != nil {
		log.Fatalf("Failed to initialize table: %v", err)
	}

	// Confirmar la transacción si todo salió bien
	if err := tx.Commit(ctx); err != nil {
		log.Fatalf("Failed to commit transaction: %v", err)
	}

	items := api.ApiGetItems() // Obtener los items de la API
	if len(items) == 0 {
		log.Println("No items found in the response.")
		return
	}

	// Calcular el score para cada item
	for i := range items {
		target_from := items[i].Target_from

		// 1. Eliminar el símbolo $
		target_from = strings.Replace(strings.Replace(target_from, ",", "", -1), "$", "", -1)

		// 2. Convertir a float64
		target_fromFloat, err := strconv.ParseFloat(target_from, 64)
		if err != nil {
			fmt.Println("Error al convertir:", err)
		}

		target_to := items[i].Target_to

		// 1. Eliminar el símbolo $
		target_to = strings.Replace(strings.Replace(target_to, ",", "", -1), "$", "", -1)

		// 2. Convertir a float64
		target_toFloat, err := strconv.ParseFloat(target_to, 64)
		if err != nil {
			fmt.Println("Error al convertir:", err)
		}

		score := algorithm.RatingChangeScore(items[i].Rating_from, items[i].Rating_to) // Calcular el puntaje del cambio de rating
		items[i].Score = algorithm.CalcularScore(target_fromFloat, target_toFloat, score)
	}

	// err = crdbpgx.ExecuteTx(context.Background(), pool, pgx.TxOptions{}, func(tx pgx.Tx) error {
	// 	return db.InsertRows(context.Background(), tx, items)
	// })
	// if err != nil {
	// 	log.Fatalf("error inserting rows: %v", err) // Manejo de errores al insertar los datos en la base de datos
	// }
	conn, err := pool.Acquire(ctx)
	if err != nil {
		log.Fatalf("Failed to acquire connection: %v", err)
	}
	defer conn.Release()

	err = crdbpgx.ExecuteTx(ctx, conn.Conn(), pgx.TxOptions{}, func(tx pgx.Tx) error {
		return db.InsertRows(ctx, tx, items)
	})
	if err != nil {
		log.Fatalf("error inserting rows: %v", err)
	}
	log.Println("New rows created.")

	// Start the server
	server.StartApiServer(pool) // Iniciar el servidor con la conexión a la base de datos
}
