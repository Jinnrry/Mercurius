package encryption

import (
	"crypto/aes"
	"crypto/md5"
)
import "crypto/cipher"
import "bytes"
import "log"

//var key = []byte("KHGSI69YBWGS0TWX")
var iv = []byte("a01020c73554d64!")

func Encrypt(text []byte, key []byte) ([]byte, error) {

	md5Ctx := md5.New()
	md5Ctx.Write(key)
	cipherStr := md5Ctx.Sum(nil)

	//生成cipher.Block 数据块
	block, err := aes.NewCipher(cipherStr)
	if err != nil {
		log.Println("错误 -" + err.Error())
		return []byte{}, err
	}
	//填充内容，如果不足16位字符
	blockSize := block.BlockSize()
	originData := pad(text, blockSize)
	//加密方式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	//加密，输出到[]byte数组
	crypted := make([]byte, len(originData))
	blockMode.CryptBlocks(crypted, originData)
	return crypted, nil
}

func pad(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func Decrypt(data []byte, key []byte) ([]byte, error) {

	md5Ctx := md5.New()
	md5Ctx.Write(key)
	cipherStr := md5Ctx.Sum(nil)

	//生成密码数据块cipher.Block
	block, _ := aes.NewCipher(cipherStr)
	//解密模式
	blockMode := cipher.NewCBCDecrypter(block, iv)
	//输出到[]byte数组
	origin_data := make([]byte, len(data))
	blockMode.CryptBlocks(origin_data, data)
	//去除填充,并返回
	return unpad(origin_data), nil
}

func unpad(ciphertext []byte) []byte {
	length := len(ciphertext)
	//去掉最后一次的padding
	unpadding := int(ciphertext[length-1])
	return ciphertext[:(length - unpadding)]
}
