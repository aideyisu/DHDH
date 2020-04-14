package main

import (
	"fmt"
	"strconv"

	aes "aideyisu/DHDH/src/aes"

	"github.com/gin-gonic/gin"
)

// PrintMixedKeys prints the result of the secret key from both participants in the communication
func PrintMixedKeys(aliceMixedKey int64, bobMixedKey int64) {
	fmt.Println("\n" + "--------------------------------------------------" + "\n")
	fmt.Println("\n\t" + "OBTENDO CHAVES MISTURADAS" + "\n")
	fmt.Print("\tMixed key da Alice = ")
	fmt.Println(aliceMixedKey)
	fmt.Print("\tMixed key do Bob = ")
	fmt.Println(bobMixedKey)
}

// PrintKeyExchange shows on terminal window the exchange of the mixedKey from both participants in the communication
func PrintKeyExchange(aliceMixedKey int64, bobMixedKey int64) {
	fmt.Println("\n" + "--------------------------------------------------" + "\n")
	fmt.Println("\n\t" + "TROCA DE CHAVES" + "\n")
	fmt.Print("\tAlice envia ")
	fmt.Print(aliceMixedKey)
	fmt.Println(" para Bob")

	fmt.Print("\tBob envia ")
	fmt.Print(bobMixedKey)
	fmt.Println(" para Alice")
	fmt.Println("\n" + "--------------------------------------------------" + "\n")
}

// PrintCommonSecretKey shows on terminal window the resultant secret common keys from both participants in the communication
func PrintCommonSecretKey(aliceMixedSecret int64, bobMixedSecret int64) {
	fmt.Println("\n\t" + "CHAVES COMPARTILHADAS OBTIDAS" + "\n")
	fmt.Print("\tChave de Alice obtida: \t")
	fmt.Println(aliceMixedSecret)
	fmt.Print("\tChave de Bob obtida: \t")
	fmt.Println(bobMixedSecret)
	fmt.Println("\n" + "--------------------------------------------------" + "\n")
}

// PrintSecretKeys shows on terminal window the secret key from bob and alice used as base for the mix with the common prime and square root
func PrintSecretKeys(aliceSecretKey float64, bobSecretKey float64) {
	fmt.Println("\n" + "--------------------------------------------------" + "\n")
	fmt.Println("\n\t" + "CHAVES SECRETAS" + "\n")
	fmt.Print("\tChave secreta de Alice ")
	fmt.Print(aliceSecretKey)

	fmt.Print("\n\tChave secreta de Bob ")
	fmt.Println(bobSecretKey)
}

var (
	// commonPrime      int64   = 1091
	// commonPrime      int64   = 105929
	// commonPrime      int64   = 1301077
	// commonPrime      int64   = 15487457
	// commonPrime      int64   = 86033551
	// commonPrime      int64   = 122955661
	// commonPrime      int64   = 160487039
	// commonPrime      int64   = 236893021
	commonPrime      int64   = 548609707
	commonSquareRoot float64 = 5
	aliceSecretKey   float64 = 424124212412
	bobSecretKey     float64 = 351232133213
	// commonSquareRoot float64 = 9
)

func modularExponent(x int64, y int64, modulos int64) int64 {
	if y == 0 {
		return 1
	}
	if y%2 == 1 {
		return (x * modularExponent(x, y-1, modulos)) % modulos
	}
	t := modularExponent(x, y/2, modulos)
	return (t * t) % modulos
}

// mixKeys Returns the "Mix" of the Secret Key operated with the common prime and square root base.
// (( commonSquare ** privateKey ) mod commonPrime )
func mixKeys(privateKey float64) int64 {
	mixedKey := modularExponent(int64(commonSquareRoot), int64(privateKey), commonPrime)
	return int64(int64(mixedKey) % commonPrime)
}

// mixSecretKeys Returns a "Mix" of the received mixed key from the other part
// in the communication, and operates it with own's secretKey
// (( receivedMixKey ** ownSecretKey ) mod commonPrime )
func mixSecretKeys(receivedMixKey int64, ownSecretKey float64) int64 {
	mixedKey := modularExponent(int64(receivedMixKey), int64(ownSecretKey), commonPrime)
	return int64(int64(mixedKey) % commonPrime)
}

func main() {

	PrintSecretKeys(aliceSecretKey, bobSecretKey)

	aliceMixedKey := mixKeys(aliceSecretKey)
	bobMixedKey := mixKeys(bobSecretKey)

	PrintMixedKeys(aliceMixedKey, bobMixedKey)

	PrintKeyExchange(aliceMixedKey, bobMixedKey)

	aliceMixedSecret := mixSecretKeys(bobMixedKey, aliceSecretKey)
	bobMixedSecret := mixSecretKeys(aliceMixedKey, bobSecretKey)

	PrintCommonSecretKey(aliceMixedSecret, bobMixedSecret)

	encrData := aes.AesEncrypt([]byte(strconv.Itoa(int(aliceMixedSecret))), []byte("1234567890qwertyu"), []byte("HHHHHasdasdas"))
	origData := aes.AesDecrypt([]byte(strconv.Itoa(int(aliceMixedSecret))), []byte("1234567890qwertyu"), encrData)
	fmt.Println(encrData, origData)
	if string(origData) != "HHHHHasdasdas" {
		fmt.Println("出错啦！")
	}

	r := gin.Default()
	r.GET("/pre", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"红烧牛肉-初始秘钥": aliceSecretKey,
			"醋溜砖头-初始秘钥": bobSecretKey,
		})
	})

	r.GET("/start", func(c *gin.Context) {
		aliceMixedKey := mixKeys(aliceSecretKey)
		bobMixedKey := mixKeys(bobSecretKey)
		c.JSON(200, gin.H{
			"红烧牛肉-混合秘钥": aliceMixedKey,
			"醋溜砖头-混合秘钥": bobMixedKey,
		})
	})

	r.GET("/finnal", func(c *gin.Context) {
		aliceMixedSecret := mixSecretKeys(bobMixedKey, aliceSecretKey)
		bobMixedSecret := mixSecretKeys(aliceMixedKey, bobSecretKey)

		c.JSON(200, gin.H{
			"红烧牛肉-最终秘钥": aliceMixedSecret,
			"醋溜砖头-最终秘钥": bobMixedSecret,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
