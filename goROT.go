package main

import(
	"fmt"
	"bufio"
	"os"
)

func encrypt(text string, shift int)string{

	if shift>26 {
		shift = shift-26
	}

	letters := [26]string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}

	var lengthText = len([]rune(text))

	var encrypted string

	for i := 0; i < lengthText; i++{
		for a := 0; a < 26; a++{
			if string(text[i]) == letters[a] {
				if a+shift > 25 {
					index := (a+shift)-26
					encrypted = encrypted+letters[index]
				} else {
					encrypted = encrypted+letters[a+shift]
				}
			}
		}
	} 

	return encrypted
}

func main(){

	fmt.Println("How many letters you want to shift (eg. ROT13 - input is shifted by 13 letters): ")

	var shift int

	fmt.Scanln(&shift)

	fmt.Println("Enter words you want to encrypt: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	toEncrypt := scanner.Text()

	var encrypted string

	encrypted = encrypt(toEncrypt, shift)

	fmt.Println("input: ", toEncrypt, " output: ", encrypted)
}