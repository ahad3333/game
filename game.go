package main

import (
	"errors"
	"fmt"
	"log"
)

func game() (string, error) {
	var a, b, c, d, e, f, j, k, l interface{}
	a, b, c, d, e, f, j, k, l = 1, 2, 3, 4, 5, 6, 7, 8, 9
	var num int
	var num1 int
	for i := 0; i < 10; i++ {
		fmt.Print("Plear1:")
		fmt.Scan(&num)
		if num == a && a != "O" {
			a = "X"
		} else if num == b && b != "O" {
			b = "X"
		} else if num == c && c != "O" {
			c = "X"
		} else if num == d && d != "O" {
			d = "X"
		} else if num == e && e != "O" {
			e = "X"
		} else if num == f && f != "O" {
			f = "X"
		} else if num == j && j != "O" {
			j = "X"
		} else if num == k && k != "O" {
			k = "X"
		} else if num == l && l != "O" {
			l = "X"
		} else if num > 10 {
			return "sorry", errors.New("Plear1: Siz o'yin qoydasni buzdingiz! ")

		} else if num == 10 {
			return "sorry", errors.New("Plear1: Siz O'yni toxtingiz! ")
		}
		fmt.Printf("%v |%v| %v", a, b, c)
		fmt.Println()
		fmt.Printf("%v |%v| %v", d, e, f)
		fmt.Println()
		fmt.Printf("%v |%v| %v", j, k, l)
		fmt.Println()
		if a == "X" && b == "X" && c == "X" || a == "X" && d == "X" && j == "X" {
			return "Plear1: Win", nil
		} else if b == "X" && e == "X" && k == "X" || c == "X" && f == "X" && l == "X" {
			return "Plear1: Win", nil
		} else if a == "X" && e == "X" && l == "X" || j == "X" && e == "X" && c == "X" {
			return "Plear1: Win", nil
		} else if l == "X" && k == "X" && j == "X" || d == "X" && e == "X" && f == "X" {
			return "Plear1: Win", nil
		}
		for s := 0; s < 1; s++ {
			fmt.Print("Plear2:")
			fmt.Scan(&num1)
			if num1 == a && a != "X" {
				a = "O"
			} else if num1 == b && b != "X" {
				b = "O"
			} else if num1 == c && c != "X" {
				c = "O"
			} else if num1 == d && d != "X" {
				d = "O"
			} else if num1 == e && e != "X" {
				e = "O"
			} else if num1 == f && f != "X" {
				f = "O"
			} else if num1 == j && j != "X" {
				j = "O"
			} else if num1 == k && k != "X" {
				k = "O"
			} else if num1 == l && l != "X" {
				l = "O"
			} else if num1 > 10 {
				return "sorry", errors.New("Plear2: Siz o'yin qoydasni buzdingiz! ")

			} else if num1 == 10 {
				return "sorry", errors.New("Plear2: Siz O'yni toxtingiz! ")
			}
			fmt.Printf("%v |%v| %v", a, b, c)
			fmt.Println()
			fmt.Printf("%v |%v| %v", d, e, f)
			fmt.Println()
			fmt.Printf("%v |%v| %v", j, k, l)
			fmt.Println()
			if a == "O" && b == "O" && c == "O" || a == "O" && d == "O" && j == "O" {
				return "Plear2: Win", nil
			} else if b == "O" && e == "O" && k == "O" || c == "O" && f == "O" && l == "O" {
				return "Plear2: Win", nil
			} else if a == "O" && e == "O" && l == "O" || j == "O" && e == "O" && c == "O" {
				return "Plear2: Win", nil
			} else if l == "O" && k == "O" && j == "O" || d == "O" && e == "O" && f == "O" {
				return "Plear2: Win", nil
			}
		}
	}
	return "End", nil
}
func main() {
	fmt.Printf("%d |%d| %d", 1, 2, 3)
	fmt.Println()
	fmt.Printf("%d |%d| %d", 4, 5, 6)
	fmt.Println()
	fmt.Printf("%d |%d| %d", 7, 8, 9)
	fmt.Println()
	fmt.Println("O'yni to'xtatish uchun 10 bosing :-)")

	win, err := game()
	if err != nil && win == "sorry" {
		log.Println(err)
		return
	}
	fmt.Println(win)
}
