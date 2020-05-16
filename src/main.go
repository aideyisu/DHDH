package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	aes "aideyisu/DHDH/src/aes"
	dh "aideyisu/DHDH/src/dh"
	"hash/adler32"
	crc32 "hash/crc32"

	"github.com/gin-gonic/gin"
)

const Size = 4

const (
	CRYPT_KEY_256 = "1~$c31kjtR^@@c2#9&iy"
	CRYPT_KEY_128 = "c31kjtR^@@c2#9&"
)

func calc(t int, b []byte) {
	var ret uint32
	if ADLER32 == t {
		ret = adler32.Checksum([]byte(b))
		fmt.Printf("ADLER32 %15d  : %s...  \n", ret, string(b[:5]))
	} else if CRC32 == t {
		ret = crc32.ChecksumIEEE([]byte(b))
		fmt.Printf("CRC32   %15d  : %s...  \n", ret, string(b[:5]))
	} else {
		return
	}
}

var ADLER32 int = 0
var CRC32 int = 1

func main() {
	var HistoryMessage [5]string
	HistoryMessage[0] = "None"
	HistoryMessage[1] = "None"
	HistoryMessage[2] = "None"
	HistoryMessage[3] = "None"
	HistoryMessage[4] = "None"

	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	InputIP := "10.255.252.93"
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		// if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
		if ipnet, ok := address.(*net.IPNet); ok {
			if ipnet.IP.String() == InputIP {
				fmt.Println(ipnet.IP.String())
			}
		}
	}

	for _, v := range []string{"aaaaaaaaaa", "3333sdfsdffsdffsd", "234esrewr234324", `An Adler-32 checksum is obtained by calculating two 16-bit checksums A and B and concatenating their bits into a 32-bit integer. A is the sum of all bytes in the stream plus one, and B is the sum of the individual values of A from each step.
					At the beginning of an Adler-32 run, A is initialized to 1, B to 0. The sums are done modulo 65521 (the largest prime number smaller than 216). The bytes are stored in network order (big endian), B occupying the two most significant bytes.
					The function may be expressed as
					A = 1 + D1 + D2 + ... + Dn (mod 65521)
					 B = (1 + D1) + (1 + D1 + D2) + ... + (1 + D1 + D2 + ... + Dn) (mod 65521)
					   = n×D1 + (n−1)×D2 + (n−2)×D3 + ... + Dn + n (mod 65521)
					 Adler-32(D) = B × 65536 + A
					where D is the string of bytes for which the checksum is to be calculated, and n is the length of D.`} {
		calc(ADLER32, []byte(v))
		calc(CRC32, []byte(v))
	}
	// var IEEE  = 0xedb88320
	// var IEEETable = crc32.simpleMakeTable(IEEE)

	// check_str := "Discard medicine more than two years old."

	// ieee := crc32.NewIEEE()
	// io.WriteString(ieee, check_str)
	// s := ieee.Sum32()

	// fmt.Printf("IEEE(%s) = 0x%x", check_str, s)
	dh.TestDH()

	// 此处为AES加密解密示例,DH可以参考此处
	encrData := aes.AesEncrypt([]byte(CRYPT_KEY_128), []byte("1234567890qwertyu"), []byte("HHHHHasdasdas"))
	origData := aes.AesDecrypt([]byte(CRYPT_KEY_128), []byte("1234567890qwertyu"), encrData)

	// encrData := aes.AesEncrypt([]byte(strconv.Itoa(int(aliceMixedSecret))), []byte("1234567890qwertyu"), []byte("HHHHHasdasdas"))
	// origData := aes.AesDecrypt([]byte(strconv.Itoa(int(aliceMixedSecret))), []byte("1234567890qwertyu"), encrData)

	fmt.Println("HHHHHasdasdas", string(origData))
	if string(origData) != "HHHHHasdasdas" {
		fmt.Println("出错啦！")
	}
	aliceMixedKey := dh.MixKeys(dh.AliceSecretKey)
	bobMixedKey := dh.MixKeys(dh.BobSecretKey)
	aliceMixedSecret := dh.MixSecretKeys(bobMixedKey, dh.AliceSecretKey)
	bobMixedSecret := dh.MixSecretKeys(aliceMixedKey, dh.BobSecretKey)

	fmt.Println("初始化完成")
	fmt.Printf("\n欢迎来到红烧牛肉和醋溜砖头的爱恨情仇\n")
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/test", func(c *gin.Context) {
		c.HTML(200, "test.html", gin.H{})
	})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Main site",
		})
	})

	r.GET("/info", func(c *gin.Context) {

		c.HTML(200, "info.html", gin.H{
			"HsSKey": dh.AliceSecretKey,
			"ClSKey": dh.BobSecretKey,
			"HsMKey": aliceMixedKey,
			"ClMKey": bobMixedKey,
			"HsFKey": aliceMixedSecret,
			"ClFKey": bobMixedSecret,
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
		EncrData := aes.AesEncrypt([]byte(CRYPT_KEY_128), []byte("1234567890qwertyu"), []byte(param.Message))
		OrigData := aes.AesDecrypt([]byte(CRYPT_KEY_128), []byte("1234567890qwertyu"), EncrData)

		K := 0
		for i := 0; i < 5; i++ {
			if HistoryMessage[i] == "None" {
				HistoryMessage[i] = "红烧牛肉to醋溜砖头 : " + param.Message
				K = 1
				break
			}
		}
		if K == 0 {
			HistoryMessage[0] = HistoryMessage[1]
			HistoryMessage[1] = HistoryMessage[2]
			HistoryMessage[2] = HistoryMessage[3]
			HistoryMessage[3] = HistoryMessage[4]
			HistoryMessage[4] = "红烧牛肉to醋溜砖头 : " + param.Message
		}
		c.JSON(200, gin.H{
			"加密Message": string(EncrData),
			"红烧牛肉to醋溜砖头验证解密信息": string(OrigData),
			"源信息": param.Message,
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
		EncrData := aes.AesEncrypt([]byte(CRYPT_KEY_128), []byte("1234567890qwertyu"), []byte(param.Message))
		OrigData := aes.AesDecrypt([]byte(CRYPT_KEY_128), []byte("1234567890qwertyu"), EncrData)
		K := 0
		for i := 0; i < 5; i++ {
			if HistoryMessage[i] == "None" {
				HistoryMessage[i] = "醋溜砖头to红烧牛肉 : " + param.Message
				K = 1
				break
			}
		}
		if K == 0 {
			HistoryMessage[0] = HistoryMessage[1]
			HistoryMessage[1] = HistoryMessage[2]
			HistoryMessage[2] = HistoryMessage[3]
			HistoryMessage[3] = HistoryMessage[4]
			HistoryMessage[4] = "醋溜砖头to红烧牛肉 : " + param.Message
		}
		c.JSON(200, gin.H{
			"加密Message": string(EncrData),
			"醋溜砖头to红烧牛肉验证解密信息": string(OrigData),
			"源信息": param.Message,
		})
	})

	r.GET("/history", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"最近的一条消息": HistoryMessage[4],
			"第二条消息":   HistoryMessage[3],
			"第三条消息":   HistoryMessage[2],
			"第四条消息":   HistoryMessage[1],
			"第五条消息":   HistoryMessage[0],
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
