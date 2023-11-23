package encrypt

import "fmt"

func Nimbus (str string) string{
	encryptedStr := ""
	for _,c := range str{
		asciiCode := int(c)
		character := fmt.Sprintf("%c", asciiCode+3)
		encryptedStr+=character
	}	
	return encryptedStr
}