package public

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"errors"
)

var ErrAesNotFound = errors.New("aes key does not exist")

func AesEncrypt(key, iv, origData []byte) (string, error) {
	if 0 == len(key) || 0 == len(iv) {
		return "", ErrAesNotFound
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding1(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	crypted := make([]byte, len(origData))

	blockMode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func AesDecrypt(key, iv []byte, crypted string) ([]byte, error) {
	if 0 == len(key) || 0 == len(iv) {
		return nil, ErrAesNotFound
	}
	decodeData, err := base64.StdEncoding.DecodeString(crypted)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	origData := make([]byte, len(decodeData))
	blockMode.CryptBlocks(origData, decodeData)
	origData = PKCS5UnPadding1(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func PKCS5Padding1(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding1(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
