package db

import (
	"backend/api"
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func PrintItems(conn *pgx.Conn, c *fiber.Ctx) error {

	// 1. Leer parámetro de página desde la URL
	page := c.QueryInt("page", 1)
	if page < 1 {
		page = 1
	}

	rows, err := conn.Query(context.Background(),
		`SELECT id, ticket, company, action, brokerage, target_to, target_from, rating_from, rating_to, time, page_count, order_index, score FROM items WHERE page_count = $1 ORDER BY order_index`, page)
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
			&item.Score,
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

func PrintRecomendedItems(conn *pgx.Conn, c *fiber.Ctx) error {
	rows, err := conn.Query(context.Background(),
		`SELECT id, ticket, company, action, brokerage, target_to, target_from, rating_from, rating_to, time, page_count, order_index, score FROM items ORDER BY score DESC LIMIT 12`)
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
			&item.Score,
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
