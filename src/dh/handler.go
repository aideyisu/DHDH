package dh

import (
	"fmt"
)

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
	AliceSecretKey   float64 = 424124212412
	BobSecretKey     float64 = 351232133213
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

// MixKeys Returns the "Mix" of the Secret Key operated with the common prime and square root base.
// (( commonSquare ** privateKey ) mod commonPrime )
func MixKeys(privateKey float64) int64 {
	mixedKey := modularExponent(int64(commonSquareRoot), int64(privateKey), commonPrime)
	return int64(int64(mixedKey) % commonPrime)
}

// MixSecretKeys Returns a "Mix" of the received mixed key from the other part
// in the communication, and operates it with own's secretKey
// (( receivedMixKey ** ownSecretKey ) mod commonPrime )
func MixSecretKeys(receivedMixKey int64, ownSecretKey float64) int64 {
	mixedKey := modularExponent(int64(receivedMixKey), int64(ownSecretKey), commonPrime)
	return int64(int64(mixedKey) % commonPrime)
}

// TestDH 测试DH功能完整性
func TestDH() {
	PrintSecretKeys(AliceSecretKey, BobSecretKey)

	aliceMixedKey := MixKeys(AliceSecretKey)
	bobMixedKey := MixKeys(BobSecretKey)

	PrintMixedKeys(aliceMixedKey, bobMixedKey)

	PrintKeyExchange(aliceMixedKey, bobMixedKey)

	aliceMixedSecret := MixSecretKeys(bobMixedKey, AliceSecretKey)
	bobMixedSecret := MixSecretKeys(aliceMixedKey, BobSecretKey)

	PrintCommonSecretKey(aliceMixedSecret, bobMixedSecret)
	fmt.Println("DH模块初始化完成")
}
