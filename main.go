package main

import (
	"log"
	"test-back-end-goneolink/service"
)

func main() {
	AlienCollection := map[string]int{"A": 1, "B": 5, "Z": 10, "L": 50, "C": 100, "D": 500, "R": 1000}
	alienSvc := service.NewConvertAlienToNumberService(AlienCollection)
	testCase := [...]string{"AAA", "LBAAA", "RCRZCAB"}
	for _, tc := range testCase {

		numberValue := alienSvc.ConvertAlienToNumber(tc)
		log.Println("input alien:",tc)
		log.Println("number value:", numberValue)
	}
}
