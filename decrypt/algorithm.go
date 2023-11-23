package decrypt

import "fmt"

func Nimbus(str string) string{
	decryptedStr := ""
	for _,c := range str{
		asciiCode := int(c)
		character := fmt.Sprintf("%c", asciiCode-3)
		decryptedStr+=character
	}	
	return decryptedStr
}