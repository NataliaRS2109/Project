package db

import (
	"backend/api"
	"context"
	"fmt"
	"log"

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
		"CREATE TABLE items (id UUID PRIMARY KEY DEFAULT gen_random_uuid(), ticket TEXT, company TEXT, action TEXT, brokerage TEXT, target_to TEXT, target_from TEXT, rating_from TEXT, rating_to TEXT, time TEXT, page_count INT)"); err != nil {
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
				target_to, target_from, rating_from, rating_to, time, page_count
			) VALUES (
				$1, $2, $3, $4, $5,
				$6, $7, $8, $9, $10, $11
			)`, // Insertar cada elemento en la tabla "items"
			uuid.New(), item.Ticker, item.Company, item.Action, item.Brokerage,
			item.Target_to, item.Target_from, item.Rating_from, item.Rating_to, item.Time, item.PageCount, // Usar uuid.New() para generar un nuevo UUID para cada fila
		); err != nil {
			return err
		}
	}
	return nil
}

func PrintItems(conn *pgx.Conn) error { // Imprimir los elementos insertados en la base de datos.
	rows, err := conn.Query(context.Background(), // Obtener todas las filas de la tabla "items"
		`SELECT id, ticket, company, action, brokerage, target_to, target_from, rating_from, rating_to, time, page_count FROM items WHERE page_count = 1`) // Obtener todas las filas de la tabla "items"
	if err != nil {
		return fmt.Errorf("error querying items: %w", err) // Manejo de errores al consultar la base de datos
	}
	defer rows.Close() // Asegurarse de cerrar las filas al final

	for rows.Next() { // Iterar sobre cada fila obtenida
		// Crear variables para almacenar los datos de cada fila
		var id uuid.UUID
		var item api.Item

		// Leer los campos de la fila
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
			&item.PageCount, // Agregar PageCount para mostrar el número de página
		); err != nil {
			return fmt.Errorf("error scanning row: %w", err)
		}

		// Mostrar datos detallados
		fmt.Println("─────────────────────────────")
		fmt.Printf("ID:            %s\n", id)
		fmt.Printf("Ticker:        %s\n", item.Ticker)
		fmt.Printf("Company:       %s\n", item.Company)
		fmt.Printf("Action:        %s\n", item.Action)
		fmt.Printf("Brokerage:     %s\n", item.Brokerage)
		fmt.Printf("Target From:   %s\n", item.Target_from)
		fmt.Printf("Target To:     %s\n", item.Target_to)
		fmt.Printf("Rating From:   %s\n", item.Rating_from)
		fmt.Printf("Rating To:     %s\n", item.Rating_to)
		fmt.Printf("Time:          %s\n", item.Time)
		fmt.Printf("Page:         %d\n", item.PageCount)
	}

	if err := rows.Err(); err != nil {
		return fmt.Errorf("row iteration error: %w", err)
	}

	return nil
}
