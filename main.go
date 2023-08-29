package main

import (
    "fmt"
    "strings"
	"os"
	"bufio"
	"strconv" 
)

func check(text string) (int, int) {
	if number, err := strconv.Atoi(text); err == nil {
		if number <= 0 || number > 10 {
			fmt.Println("Error.")
			os.Exit(0)
		}
		return number, 1
	}
	return rimtodec(text), 2
}
	
func check_operation(text string, x, y int) int {
	switch text {
		case "+":
			return x + y
		case "-":
			return x - y
		case "*":
			return x * y
		case "/": {
			if y != 0 { 
			return x / y
			}
		}
	} 
	return 10000
}


func rimtodec(text string) int {
	switch text {
	case "I":  return 1
	case "II": return 2
	case "III": return 3
	case "IV": return 4
	case "V": return 5
	case "VI": return 6
	case "VII": return 7
	case "VIII": return 8
	case "IX": return 9
	case "X": return 10
	}
	return 0
}

var rom_hlp = [11]string {"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

func dectorim(number int) string{
    if number >= 90 {
        if number == 100 {
			return "C"
		} else {
			return "XC" + dectorim(number - 90)
    	}
	} else {
        if number <= 10 {
			return rom_hlp[number]
		}
        if number >= 40 {
            if number >=50 {
				return "L" + dectorim(number - 50)
			} else {
				return "XL" + dectorim(number - 40)
            }
		}
		return "X" + dectorim(number - 10)
   }
 }


func main(){

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	elem := strings.Fields(text)
	if len(elem) != 3 {
		fmt.Println("Error.")
		os.Exit(0)
	}
	x, sx := check(elem[0])
	y, sy := check(elem[2])

	if sx == sy && sx == 1 && check_operation(elem[1], x, y) != 10000 {		
		fmt.Println(check_operation(elem[1], x, y))
	} else if sx == sy && sx == 2 && check_operation(elem[1], x, y) > 0 &&
	    	  check_operation(elem[1], x, y) != 10000 && x != 0 && y != 0 {
		fmt.Println(dectorim(check_operation(elem[1], x, y)))
	} else {
		fmt.Println("Error.")
		os.Exit(0)
	}
}
