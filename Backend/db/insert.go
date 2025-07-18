package db

import (
	"backend/api"
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func InsertRows(ctx context.Context, tx pgx.Tx, items []api.Item) error { // Insert rows into the "items" table.
	log.Println("Creating new rows...")
	for _, item := range items { // Iterar sobre cada elemento en los resultados
		if _, err := tx.Exec(ctx,
			`INSERT INTO items (
				id, ticker, company, action, brokerage,
				target_to, target_from, rating_from, rating_to, time, page_count, order_index, score
			) VALUES (
				$1, $2, $3, $4, $5,
				$6, $7, $8, $9, $10, $11, $12, $13
			)`, // Insertar cada elemento en la tabla "items"
			uuid.New(), item.Ticker, item.Company, item.Action, item.Brokerage,
			item.Target_to, item.Target_from, item.Rating_from, item.Rating_to, item.Time, item.PageCount, item.OrderIndex, item.Score, // Usar uuid.New() para generar un nuevo UUID para cada fila
		); err != nil {
			return err
		}
	}
	return nil
}
