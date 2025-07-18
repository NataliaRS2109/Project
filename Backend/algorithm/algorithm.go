package algorithm

import (
	"strings"
)

// ratingChangeScore calcula el puntaje según el cambio de rating
func RatingChangeScore(from, to string) float64 {
	from = strings.ToLower(from)
	to = strings.ToLower(to)

	if to == "buy" {
		switch from {
		case "sell":
			return 3.0
		case "hold":
			return 2.0
		case "neutral":
			return 1.0
		case "buy":
			return 0.5
		}
	}
	return 0.0 // No interesa si no es hacia Buy
}

// calcularScore calcula el score total
func CalcularScore(targetFrom, targetTo float64, score float64) float64 {
	deltaTarget := targetTo - targetFrom
	ratingScore := score
	return deltaTarget + ratingScore
}
