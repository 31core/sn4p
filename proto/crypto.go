package proto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
)

/* 生成RSA密钥 */
func GenerateRSAKey() (*rsa.PrivateKey, *rsa.PublicKey) {
	privatekey, _ := rsa.GenerateKey(rand.Reader, 2048)
	publickey := &privatekey.PublicKey
	return privatekey, publickey
}

/* 公钥dump到[]byte类型 */
func DumpPublicKey(key *rsa.PublicKey) []byte {
	var keybytes []byte = x509.MarshalPKCS1PublicKey(key)
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: keybytes,
	}
	return pem.EncodeToMemory(block)
}

/* []byte类型load公钥 */
func LoadPublicKey(key []byte) *rsa.PublicKey {
	block, _ := pem.Decode(key)
	privatekey, _ := x509.ParsePKCS1PublicKey(block.Bytes)
	return privatekey
}

/* RSA加密 */
func EncryptRSA(src []byte, publickey *rsa.PublicKey) []byte {
	label := []byte("")
	sha256hash := sha256.New()
	cipher, _ := rsa.EncryptOAEP(sha256hash, rand.Reader, publickey, src, label)
	return cipher
}

/* RSA解密 */
func DecryptRSA(src []byte, privatekey *rsa.PrivateKey) []byte {
	sha256hash := sha256.New()
	decrypt, _ := rsa.DecryptOAEP(sha256hash, rand.Reader, privatekey, src, nil)
	return decrypt
}

/* 填充数据 */
func padding(src []byte, block int) []byte {
	num := block - len(src)%block
	pad := bytes.Repeat([]byte{byte(num)}, num)
	return append(src, pad...)
}

/* 去掉填充数据 */
func unpadding(src []byte) []byte {
	num := int(src[len(src)-1])
	return src[:len(src)-num]
}

/* AES加密 */
func EncryptAES(src []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	src = padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)
	return src
}

/* AES解密 */
func DecryptAES(src []byte, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(src, src)
	src = unpadding(src)
	return src
}
