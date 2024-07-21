package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Arab2Roman(a int) string {
	ans := ""
	roms := [5]string{"C", "L", "X", "V", "I"}
	for a != 0 {
		if a == 4 {
			a -= 4
			ans += "IV"
		} else if a == 9 {
			a -= 9
			ans += "IX"
		} else if a == 45 {
			a -= 45
			ans += "XLV"
		} else if a == 40 {
			a -= 40
			ans += "XL"
		} else if a == 90 {
			a -= 90
			ans += "XC"
		}
		for i, value := range [5]int{100, 50, 10, 5, 1} {
			if a >= value {
				a -= value
				ans += roms[i]
				break
			}
		}
	}
	return ans
}

func Roman2Arab(a string) int {
	ans := 0
	var romanians = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100}
	if len(a) > 1 {
		for i := 0; i < len(a)-1; i++ {
			if romanians[string(a[i+1])] > romanians[string(a[i])] {
				ans -= romanians[string(a[i])]
			} else {
				ans += romanians[string(a[i])]
			}
		}
		ans += romanians[string(a[len(a)-1])]
	} else {
		ans += romanians[a]
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
		panic("Выдача паники, так как были использованы нецелые числа")
	}
	x, errx := strconv.Atoi(a)
	y, erry := strconv.Atoi(b)
	var IsRom bool
	if (Roman2Arab(a) > 0 && Roman2Arab(b) > 0) || (errx == nil && erry == nil) {
		if Roman2Arab(a) > 0 {
			x = Roman2Arab(a)
			y = Roman2Arab(b)
			IsRom = true
		} else {
			IsRom = false
		}
		if x > 10 || x < 1 || y > 10 || y < 1 {
			panic("Выдача паники, так как числа должны быть от 1 до 10 включительно.")
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
		panic("Выдача паники, так как не используются одновременно арабская или римская системы счисления.")
	}
}
