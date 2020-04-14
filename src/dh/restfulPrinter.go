package dh

import "fmt"

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
