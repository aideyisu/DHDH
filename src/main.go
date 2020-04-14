package main

import (
	"fmt"
	"net/http"

	aes "aideyisu/DHDH/src/aes"
	dh "aideyisu/DHDH/src/dh"

	"github.com/gin-gonic/gin"
)

const (
	CRYPT_KEY_256 = "1~$c31kjtR^@@c2#9&iy"
	CRYPT_KEY_128 = "c31kjtR^@@c2#9&"
)

func main() {
	dh.TestDH()

	encrData := aes.AesEncrypt([]byte(CRYPT_KEY_128), []byte("1234567890qwertyu"), []byte("HHHHHasdasdas"))
	origData := aes.AesDecrypt([]byte(CRYPT_KEY_128), []byte("1234567890qwertyu"), encrData)

	// encrData := aes.AesEncrypt([]byte(strconv.Itoa(int(aliceMixedSecret))), []byte("1234567890qwertyu"), []byte("HHHHHasdasdas"))
	// origData := aes.AesDecrypt([]byte(strconv.Itoa(int(aliceMixedSecret))), []byte("1234567890qwertyu"), encrData)

	fmt.Println(encrData, origData)
	if string(origData) != "HHHHHasdasdas" {
		fmt.Println("出错啦！")
	}
	fmt.Printf("\n欢迎来到红烧牛肉和醋溜砖头的爱恨情仇\n")
	r := gin.Default()
	r.GET("/pre", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"红烧牛肉-初始秘钥": "aliceSecretKey",
			"醋溜砖头-初始秘钥": "bobSecretKey",
		})
	})

	r.GET("/start", func(c *gin.Context) {
		// aliceMixedKey := mixKeys(aliceSecretKey)
		// bobMixedKey := mixKeys(bobSecretKey)
		c.JSON(200, gin.H{
			"红烧牛肉-混合秘钥": "aliceMixedKey",
			"醋溜砖头-混合秘钥": "bobMixedKey",
		})
	})

	r.GET("/finnal", func(c *gin.Context) {
		// aliceMixedSecret := mixSecretKeys(bobMixedKey, aliceSecretKey)
		// bobMixedSecret := mixSecretKeys(aliceMixedKey, bobSecretKey)

		c.JSON(200, gin.H{
			"红烧牛肉-最终秘钥": "aliceMixedSecret",
			"醋溜砖头-最终秘钥": "bobMixedSecret",
		})
	})

	r.GET("/BtoC", func(c *gin.Context) {
		type Param struct {
			Message string
		}
		var param Param
		if err := c.ShouldBind(&param); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status_code": 400, "reason": fmt.Sprint(err.Error())})
			return // exit on first error
		}
		c.JSON(200, gin.H{
			"加密Message":  "加密",
			"红烧牛肉to醋溜砖头": param.Message,
		})
	})

	r.GET("/CtoB", func(c *gin.Context) {
		type Param struct {
			Message string
		}
		var param Param
		if err := c.ShouldBind(&param); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status_code": 400, "reason": fmt.Sprint(err.Error())})
			return // exit on first error
		}
		c.JSON(200, gin.H{
			"加密Message":  "加密",
			"醋溜砖头to红烧牛肉": param.Message,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
