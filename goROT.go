package main

import(
	"fmt"
	"bufio"
	"os"
)

func encrypt(text string)string{

	letters := [26]string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}

	var lengthText = len([]rune(text))
	//var lengthLet = len([]rune(letters))   

	var encrypted string

	for i := 0; i < lengthText; i++{
		for a := 0; a < 26; a++{
			if string(text[i]) == letters[a] {
				if a+13 > 25 {
					index := (a+13)-26
					encrypted = encrypted+letters[index]
				} else {
					encrypted = encrypted+letters[a+13]
				}
			}
		}
	} 

	return encrypted
}

func main(){

	fmt.Println("Enter word you want to encrypt: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	toEncrypt := scanner.Text()

	var encrypted string

	encrypted = encrypt(toEncrypt)

	fmt.Println("input: ", toEncrypt, " output: ", encrypted)
}