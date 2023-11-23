package main

import (
	"fmt"

	"github.com/mahdibouaziz/cryptit/decrypt"
	"github.com/mahdibouaziz/cryptit/encrypt"
)

func main(){

	encryptStr:= encrypt.Nimbus("Kodekloud")
	fmt.Println(encryptStr)

	decryptStr := decrypt.Nimbus(encryptStr)
	fmt.Println(decryptStr)
}