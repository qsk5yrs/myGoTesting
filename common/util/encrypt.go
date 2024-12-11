package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

// AseEncrypt AES加密 key长度为16字节才能加密成功
func AseEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	encrypted := make([]byte, len(origData))
	blockMode.CryptBlocks(encrypted, origData)
	return encrypted, nil
}

// PKCS5Padding 在数据末尾添加填充字节，使得数据长度成为块大小的倍数
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	// 使用分组密码（如 AES）进行加密时，输入数据的长度必须是块大小的整数倍。AES 的块大小是 16 字节（128 位），所以数据长度必须是 16 的倍数。
	// 计算需要补充的字计数
	padding := blockSize - len(ciphertext)%blockSize
	// 生成补充字节的内容：添加几个字节，就生成几个以padding为字节的内容
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

// AesDecrypt AES解密
func AesDecrypt(encrypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	originData := make([]byte, len(encrypted))
	blockMode.CryptBlocks(originData, encrypted)
	originData = PKCS5UnPadding(originData)
	return originData, nil
}

// PKCS5UnPadding 去除填充
func PKCS5UnPadding(originData []byte) []byte {
	length := len(originData)
	// 去掉最后一个字节 unPadding次
	unPadding := int(originData[length-1])
	if unPadding < 1 || unPadding > 32 {
		unPadding = 0
	}
	return originData[:(length - unPadding)]
}
