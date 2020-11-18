package paginator

import (
	"encoding/json"
	"math"
)

func CalculatePageCount(itemsCount, itemsPerPage int) int {
	floatResult := float64(itemsCount) / float64(itemsPerPage)
	ceil := math.Ceil(floatResult)
	result := int(ceil)
	return result
}

type PaginatorData struct {
	HasMore bool            `json:"has_more"`
	Payload json.RawMessage `json:"payload"`
}
