package service

import "strings"

type AlienService interface {
	ConvertAlienToNumber(alienStr string) int
}
type alienService struct {
	AlienCollection map[string]int
}

// ConvertAlienToNumber implements AlienService.
func (a alienService) ConvertAlienToNumber(alienStr string) int {

	var total int
	var prevValue int

	totalStr := strings.Split(alienStr, "")

	for i, symbol := range totalStr {
		value := a.AlienCollection[symbol]

		if i == 0 {
			total = value
			prevValue = value
			continue
		}

		if value <= prevValue {

			total += value

		} else {

			total = (total - prevValue) + (value - prevValue)
		}

		prevValue = value

	}
	return total
}

func NewConvertAlienToNumberService(AlienCollection map[string]int) AlienService {
	return alienService{AlienCollection: AlienCollection}
}
