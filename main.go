package main

import (
    "fmt"
    "strings"
	"os"
	"bufio"
	"strconv" 
	"errors"
)

func check_number(text string) (int, error, int) {

	if number, err := strconv.Atoi(text); err == nil {
		if number <= 0 || number > 10 {
			return 0, errors.New("Операнды должны быть от 1 до 10."), 0
		}
		return number, nil, 1
	}
	return rimtodec(text)
}
	
func check_operation(text string, x, y int) (int, error) {
	switch text {
		case "+":
			return x + y, nil
		case "-":
			return x - y, nil
		case "*":
			return x * y, nil
		case "/": {
			if y != 0 { 
			return x / y, nil
			}
		}
	} 
	return 0, errors.New("Cтрока не является математической операцией.")
}


func rimtodec(text string) (int, error, int) {

	switch text {
	case "I":  return 1, nil, 2
	case "II": return 2,  nil, 2
	case "III": return 3,  nil, 2
	case "IV": return 4,  nil, 2
	case "V": return 5,  nil, 2
	case "VI": return 6,  nil, 2
	case "VII": return 7,  nil, 2
	case "VIII": return 8,  nil, 2
	case "IX": return 9,  nil, 2
	case "X": return 10,  nil, 2
	}
	return 0, errors.New("Операнды должны быть от 1 до 10."), 0
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

 func global_check(elem []string) error {
	var result int

	if len(elem) > 3 {
		return errors.New("Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	if len(elem) == 1 {
		return errors.New("Cтрока не является математической операцией.")
	}
	x, err, sx := check_number(elem[0])
	if err != nil {
		return err
	}
	y, err, sy := check_number(elem[2])
	if err != nil {
		return err
	}
	if sx == sy && sx != 0 {
		result, err = check_operation(elem[1], x, y)
		if err != nil {
			return err
		}
		if sx == 2 {
			if result > 0 {
				fmt.Println(dectorim(result))
				return nil
			} else {
				return errors.New("В римской системе нет отрицательных чисел.")
			}
		}
		fmt.Println(result)
		return nil
	}
	return errors.New("Используются одновременно разные системы счисления.")
 }

func main(){

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	elem := strings.Fields(text)
	if err:= global_check(elem); err != nil{
		fmt.Println(err)
	}
}
