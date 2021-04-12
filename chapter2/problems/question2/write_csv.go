package main

import (
	"encoding/csv"
	"os"
)

func main() {

	records := [][]string{
		[]string{"名前", "年齢", "身長", "体重"},
		[]string{"Tanaka", "31", "190cm", "97kg"},
		[]string{"Suzuki", "46", "180cm", "79kg"},
		[]string{"Matsui", "45", "188cm", "95kg"},
	} 
	file, err := os.Create("test.csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file)
	writer.Comma = ','
	writer.UseCRLF = true

	for _, record := range records {
		if err := writer.Write(record); err != nil {
			panic(err)
		}
	}
	writer.Flush() // バッファに残っているデータを全て書き込む
	
}