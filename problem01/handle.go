package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	fileName := "../files/hard.json"
	res := Handle(fileName)
	log.Printf("ผลลัพธ์สูงสุด: %d\n", res)
}

func Handle(fileName string) int {
	byteValue, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("ไม่สามารถอ่านไฟล์: %v", err)
	}

	var data [][]int
	if err := json.Unmarshal(byteValue, &data); err != nil {
		log.Fatalf("ไม่สามารถแปลง JSON: %v", err)
	}
	if len(data) == 0 {
		return 0
	}

	// คำนวณจากล่างขึ้นบน
	// เอาค่าตัวเองบวกกับค่ามากที่สุดจากสองทางเลือกด้านล่าง
	for i := len(data) - 2; i >= 0; i-- {
		for j := 0; j < len(data[i]); j++ {
			left := data[i+1][j]
			right := data[i+1][j+1]

			data[i][j] += findMax(left, right)
		}
	}

	return data[0][0] // (แถว 0, ตำแหน่ง 0)
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
