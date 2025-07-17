package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ApiResponse struct {
	Items    []Item `json:"items"`     // Cada elemento es un "Item" que viene dentro del array "items"
	NextPage string `json:"next_page"` // Cadena que representa la siguiente página de resultados
}

type Item struct { //Los nombres en json:"..." le dicen a Go cómo debe mapear el JSON con los campos de la estructura.
	Ticker      string `json:"ticker"`
	Target_from string `json:"target_from"`
	Target_to   string `json:"target_to"`
	Company     string `json:"company"`
	Action      string `json:"action"`
	Brokerage   string `json:"brokerage"`
	Rating_from string `json:"rating_from"`
	Rating_to   string `json:"rating_to"`
	Time        string `json:"time"`
	PageCount   int    // Contador de páginas
	OrderIndex  int    // Índice de orden para mantener el orden de los items
}

func ApiGetItems() []Item {
	baseURL := "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list"
	token := os.Getenv("API_TOKEN")

	nextPage := ""
	pageCount := 1
	order := 1
	var allItems []Item // Slice para almacenar todos los items obtenidos de la API

	for {
		var url string // Variable para almacenar la URL de la API
		if nextPage != "" {
			url = baseURL + "?next_page=" + nextPage // Si hay una página siguiente, la agregamos a la URL
		} else {
			url = baseURL
		}

		//fmt.Printf("\n📄 Página %d: %s\n", pageCount, url) // Imprimir la URL de la página actual

		request, err := http.NewRequest(http.MethodGet, url, nil) // Crear una nueva solicitud HTTP GET
		// Manejo de errores al crear la solicitud
		if err != nil {
			log.Fatalf("Error creando la solicitud: %v", err)
		}

		request.Header.Add("Authorization", "Bearer "+token) // Agregar el token de autorización al encabezado de la solicitud
		// Agregar encabezados adicionales si es necesario
		request.Header.Add("Accept", "application/json")

		client := &http.Client{} // Crear un cliente HTTP para enviar la solicitud
		// Enviar la solicitud y manejar la respuesta
		response, err := client.Do(request)
		if err != nil {
			log.Fatalf("Error al hacer la solicitud: %v", err)
		}
		defer response.Body.Close() // Asegurarse de cerrar el cuerpo de la respuesta al final

		if response.StatusCode != http.StatusOK { // Verificar si la respuesta fue exitosa (código 200 OK)
			// Si no fue exitosa, imprimir un mensaje de error y salir del programa
			log.Fatalf("Error: código %d\n", response.StatusCode)
		}

		body, err := ioutil.ReadAll(response.Body) // Leer el cuerpo de la respuesta
		// Manejo de errores al leer el cuerpo de la respuesta
		if err != nil {
			log.Fatalf("Error leyendo la respuesta: %v", err)
		}

		var page ApiResponse // Crear una variable para almacenar el resultado de la respuesta
		// Parsear el JSON de la respuesta a la estructura ApiResponse
		if err := json.Unmarshal(body, &page); err != nil {
			log.Fatalf("Error parseando JSON: %v", err)
		}

		// Añadir PageCount a cada item
		for _, item := range page.Items {

			item.PageCount = pageCount
			item.OrderIndex = order // Asignar el índice de orden al item
			allItems = append(allItems, item)
			order++ // Incrementar el índice de orden para cada item
		}

		if page.NextPage == "" {
			break
		}
		nextPage = page.NextPage
		pageCount++
	}

	return allItems
}
