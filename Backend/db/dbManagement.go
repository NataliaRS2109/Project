package db

import (
	"backend/api"
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
		"CREATE TABLE items (id UUID PRIMARY KEY DEFAULT gen_random_uuid(), ticket TEXT, company TEXT, action TEXT, brokerage TEXT, target_to TEXT, target_from TEXT, rating_from TEXT, rating_to TEXT, time TEXT, page_count INT, order_index INT)"); err != nil {
		return err
	}
	return nil
}

func InsertRows(ctx context.Context, tx pgx.Tx, items []api.Item) error { // Insert rows into the "items" table.
	log.Println("Creating new rows...")
	for _, item := range items { // Iterar sobre cada elemento en los resultados
		if _, err := tx.Exec(ctx,
			`INSERT INTO items (
				id, ticket, company, action, brokerage,
				target_to, target_from, rating_from, rating_to, time, page_count, order_index
			) VALUES (
				$1, $2, $3, $4, $5,
				$6, $7, $8, $9, $10, $11, $12
			)`, // Insertar cada elemento en la tabla "items"
			uuid.New(), item.Ticker, item.Company, item.Action, item.Brokerage,
			item.Target_to, item.Target_from, item.Rating_from, item.Rating_to, item.Time, item.PageCount, item.OrderIndex, // Usar uuid.New() para generar un nuevo UUID para cada fila
		); err != nil {
			return err
		}
	}
	return nil
}

func PrintItems(conn *pgx.Conn, c *fiber.Ctx) error {

	// 1. Leer parámetro de página desde la URL
	page := c.QueryInt("page", 1)
	if page < 1 {
		page = 1
	}

	rows, err := conn.Query(context.Background(),
		`SELECT id, ticket, company, action, brokerage, target_to, target_from, rating_from, rating_to, time, page_count, order_index FROM items WHERE page_count = $1`, page)
	if err != nil {
		return fmt.Errorf("error querying items: %w", err)
	}
	defer rows.Close()

	type Result struct {
		ID       uuid.UUID `json:"id"`
		api.Item           // incrustamos los campos de api.Item
	}

	var items []Result

	for rows.Next() {
		var id uuid.UUID
		var item api.Item

		if err := rows.Scan(
			&id,
			&item.Ticker,
			&item.Company,
			&item.Action,
			&item.Brokerage,
			&item.Target_to,
			&item.Target_from,
			&item.Rating_from,
			&item.Rating_to,
			&item.Time,
			&item.PageCount,
			&item.OrderIndex,
		); err != nil {
			return fmt.Errorf("error scanning row: %w", err)
		}

		items = append(items, Result{ID: id, Item: item})
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("row iteration error: %w", err)
	}

	return c.JSON(items)
}
