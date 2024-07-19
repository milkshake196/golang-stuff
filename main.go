package main

import (
	"fmt" // горит красным потому что неиспользован. браатииииш будь оптимизированнее
	"strconv"
	"strings"
)

// функция перевода из араб в рим
func Arab2Roman(a int) string {
	ans := ""
	roms := [3]string{"X", "V", "I"}
	for a != 0 {
		if a == 4 {
			a -= 4
			ans += "IV"
		} else if a == 9 {
			a -= 9
			ans += "IX"
		}
		for i, value := range [3]int{10, 5, 1} {
			if a >= value {
				a -= value
				ans += roms[i]
				break
			}
		}
	}
	return ans
}

/*
	else {
		panic("Выдача паники, так как строка не является математической операцией.")
	}
*/
func Roman2Arab(a string) int {
	ans := 0
	romanians := "IVX"
	runes := []rune(romanians)
	if a == "IV" {
		ans += 4
		return ans
	} else if a == "IX" {
		ans += 9
		return ans
	}
	for _, value := range a {
		if value == runes[0] {
			ans += 1
		} else if value == runes[1] {
			ans += 5
		} else if value == runes[2] {
			ans += 10
		}
	}
	return ans
}

func main() {
	var a, op, b, c string
	fmt.Scanln(&a, &op, &b, &c)
	if c != "" || !strings.Contains("+-*/", op) {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	} else if b == "" {
		panic("Выдача паники, так как строка не является математической операцией.")
	} else if strings.Contains(a, ".") || strings.Contains(b, ".") || strings.Contains(a, ",") || strings.Contains(b, ",") {
		panic("Выдача паники, так как числа должны быть целые")
	}
	x, errx := strconv.Atoi(a)
	y, erry := strconv.Atoi(b)
	var IsRom bool
	if errx == nil && erry == nil && (x > 10 || x < 1 || y > 10 || y < 1) {
		panic("Выдача паники, так как числа должны быть от 1 до 10 включительно.")
	}
	if (Roman2Arab(a) > 0 && Roman2Arab(b) > 0) || (errx == nil && erry == nil) {
		if Roman2Arab(a) > 0 {
			x = Roman2Arab(a)
			y = Roman2Arab(b)
			IsRom = true
		} else {
			IsRom = false
		}
		switch op {
		case "+":
			if IsRom {
				fmt.Println(Arab2Roman(x + y))
			} else {
				fmt.Println(x + y)
			}
		case "-":
			if IsRom && (x-y < 1) {
				panic("Выдача паники, так как в римской системе счисления нет отрицательных чисел.")
			} else if IsRom && (x-y >= 1) {
				fmt.Println(Arab2Roman(x - y))
			} else {
				fmt.Println(x - y)
			}
		case "*":
			if IsRom {
				fmt.Println(Arab2Roman(x * y))
			} else {
				fmt.Println(x * y)
			}
		case "/":
			if IsRom {
				fmt.Println(Arab2Roman(x / y))
			} else {
				fmt.Println(x / y)
			}
		}
	} else if (errx != nil || erry != nil) && (Roman2Arab(a) < 1 || Roman2Arab(b) < 1) {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}
}
