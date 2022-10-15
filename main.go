package main

import (
	"fmt"
	"strings"
)

func Reverse(str string) string {
	var backward string

	for i := len(str) - 1; i >= 0; i-- {
		backward += string(str[i])
	}
	return backward
}

func Generate(str string) string {
	str = Reverse(str)
	//contain number or symbol should just reverse
	for _, number := range str {
		switch number {
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			return str
		case '!', '@', '#', '$', '%', '^', '&', '*', '_', '-':
			return str
		}
	}
	//HOW TO CHANGE VOWEL *basic way
	var replace string
	for _, vowels := range str {
		if string(vowels) == "a" {
			replace += "E"
		} else if string(vowels) == "e" {
			replace += "I"
		} else if string(vowels) == "i" {
			replace += "O"
		} else if string(vowels) == "o" {
			replace += "U"
		} else if string(vowels) == "u" {
			replace += "A"
		} else if string(vowels) == "A" {
			replace += "e"
		} else if string(vowels) == "E" {
			replace += "i"
		} else if string(vowels) == "I" {
			replace += "o"
		} else if string(vowels) == "O" {
			replace += "u"
		} else if string(vowels) == "U" {
			replace += "a"
		} else {
			replace += string(vowels)
		}
	}
	var change string
	for _, Uplow := range replace {
		// change space and capital to lower vice versa
		if string(Uplow) == " " { //gak kedetect pas ditest but outputnya bener
			continue
		} else if string(Uplow) == strings.ToUpper(string(Uplow)) { // change uppercase to lowercase
			change += strings.ToLower(string(Uplow))
		} else if string(Uplow) == strings.ToLower(string(Uplow)) {
			change += strings.ToUpper(string(Uplow))
		}

	}
	return change
}
func CheckPassword(str string) string {
	str = Generate(str)
	//check password criteria sangat lemah,lemah,sedang,kuat
	var password string
	//numeric := "1234567890"
	//alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//Alphasym := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_-"
	//numsym := "1234567890!@#$%^&*()_-"
	Alphanumeric := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ12345678910"
	//combinasi := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ12345678910!@#$%^&*()_-"

	// gabisa ngedetect ketentuan alphanumeric dll
	if len(str) >= 14 {
		password += "kuat"
	} else if len(str) >= 7 {
		password += "sedang"
	} else if str == Alphanumeric {
		password += "lemah" //gak kedetect
	} else {
		password += "sangat lemah"
	}
	return password
}

func PasswordGenerator(base string) (string, string) { //error should generate password
	var generate = Generate(base)
	var check = CheckPassword(generate)

	return generate, check
}

func main() {
	data := "Semangat Pagi" // bisa digunakan untuk melakukan debug

	res, check := PasswordGenerator(data)

	fmt.Println(res)
	fmt.Println(check)
}
