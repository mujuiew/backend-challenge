package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func main() {
	// รับ input จาก keyboard
	fmt.Print("ใส่รหัส (ตัวอักษร L, R, =): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	encoded := scanner.Text()

	if !validate(encoded) {
		fmt.Println("รหัสไม่ถูกต้อง กรุณาใส่ใหม่")
		return
	}

	// แสดงผล
	fmt.Printf("ชุดตัวเลขที่ถอดรหัสได้และผลรวมน้อยที่สุด: %s \n", process(encoded))

}
func validate(encoded string) bool {
	// ตรวจสอบความถูกต้องของรหัส
	for _, char := range encoded {
		if char != 'L' && char != 'R' && char != '=' {
			return false
		}
	}
	return true
}

// สัญลักษณ์ “L” หมายความว่า ตัวเลขด้านซ้าย มีค่ามากกว่า ตัวเลขด้านขวา
// สัญลักษณ์ “R” หมายความว่า ตัวเลขด้านขวา มีค่ามากกว่า ตัวเลขด้านซ้าย
// สัญลักษณ์ “=“ หมายความว่า ตัวเลขด้านซ้าย มีค่าเท่ากับ ตัวเลขด้านขวา
func process(encoded string) string {
	n := len(encoded)
	result := make([]int, n+1)

	result[0] = 0

	for i := 0; i < n; i++ {
		switch encoded[i] {
		case 'L':
			result[i] = result[i+1] + 1
		case 'R':
			result[i+1] = result[i] + 1
		case '=':
			result[i+1] = result[i]
		}
	}

	for {
		before := result
		for i := 0; i < len(result)-1; i++ {
			switch encoded[i] {
			case 'L':
				if result[i] <= result[i+1] {
					result[i] = result[i+1] + 1
				}
				if (i - 1) >= 0 {
					if encoded[i-1] == '=' {
						result[i-1] = result[i]
					}
				}
			case 'R':
				if result[i] >= result[i+1] {
					result[i+1] = result[i] + 1
				}
				if (i - 1) >= 0 {
					if encoded[i-1] == '=' {
						result[i-1] = result[i]
					}
				}
			case '=':
				if result[i] != result[i+1] {
					result[i+1] = result[i]
				}
			}
		}
		if reflect.DeepEqual(before, result) {
			break
		}
	}

	var response string
	for _, num := range result {
		response += fmt.Sprintf("%d", num)
	}
	return response
}
