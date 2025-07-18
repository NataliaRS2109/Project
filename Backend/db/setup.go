package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func InitTable(ctx context.Context, tx pgx.Tx) error {
	// Eliminar la tabla "items" si ya existe
	// Esto es útil para evitar errores si la tabla ya existe al iniciar el programa.
	log.Println("Drop existing items table if necessary.")
	if _, err := tx.Exec(ctx, "DROP TABLE IF EXISTS items"); err != nil {
		return err
	}

	// Crear la tabla "items" con los campos necesarios
	// Esta tabla almacenará los datos obtenidos de la API.
	log.Println("Creating items table.")
	if _, err := tx.Exec(ctx,
		"CREATE TABLE items (id UUID PRIMARY KEY DEFAULT gen_random_uuid(), ticket TEXT, company TEXT, action TEXT, brokerage TEXT, target_to TEXT, target_from TEXT, rating_from TEXT, rating_to TEXT, time TEXT, page_count INT, order_index INT, score FLOAT)"); err != nil {
		return err
	}
	return nil
}
