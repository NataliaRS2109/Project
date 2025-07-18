package algorithm

import (
	"strings"
)

/*
¿Cómo se calcula el puntaje de recomendación?
A cada recomendación se le asigna un puntaje basado en dos componentes, pero solo se tienen en cuenta los cambios cuya recomendación final sea "Buy", ya que el objetivo del algoritmo es identificar señales positivas de inversión.

deltaTarget: diferencia entre el nuevo y el anterior precio objetivo.
deltaTarget = target_to - target_from

ratingChangeScore: valor que representa la fuerza del cambio:

Sell → Buy = +3

Hold → Buy = +2

Neutral → Buy = +1

Buy → Buy = +0.5

El puntaje final se calcula así:
score = deltaTarget + ratingChangeScore

Las recomendaciones se ordenan de mayor a menor puntaje, priorizando los cambios más optimistas.
Se descartan los cambios que no terminan en "Buy" porque este algoritmo se enfoca en destacar mejoras que representen una posible oportunidad de inversión.
*/

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
