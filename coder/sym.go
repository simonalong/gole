package coder

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rc4"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
)

func AesDecrypt(content string, key string, iv string) string {
	b, _ := base64.StdEncoding.DecodeString(content)
	block, _ := aes.NewCipher([]byte(key))
	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	originData := make([]byte, len(b))
	mode.CryptBlocks(originData, b)
	origData := pkcs5UnPadding(originData)
	return string(origData)
}

// AesEncrypt key的长度必须是 16，24，32
func AesEncrypt(content string, key string, iv string) string {
	origData := pkcs5Padding([]byte(content), aes.BlockSize)
	block, _ := aes.NewCipher([]byte(key))
	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	crypted := make([]byte, len(origData))
	mode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted)
}

// AesDecryptECBForJava 兼容java的AES解密方式
// key的长度必须是 16，24，32
func AesDecryptECBForJava(content string, key string) string {
	b, _ := base64.StdEncoding.DecodeString(content)
	cp, _ := aes.NewCipher([]byte(key))
	d := make([]byte, len(b))
	size := 16
	for bs, be := 0, size; bs < len(b); bs, be = bs+size, be+size {
		cp.Decrypt(d[bs:be], b[bs:be])
	}
	// 去除填充的
	return strings.ReplaceAll(strings.TrimSpace(string(d)), "\x00", "")
}

// AesEncryptECBForJava 兼容java的AES加密方式
func AesEncryptECBForJava(content string, key string) string {
	b := padding([]byte(content))
	cp, _ := aes.NewCipher([]byte(key))
	d := make([]byte, len(b))
	size := 16
	for bs, be := 0, size; bs < len(b); bs, be = bs+size, be+size {
		cp.Encrypt(d[bs:be], b[bs:be])
	}
	return base64.StdEncoding.EncodeToString(d)
}

// AesDecryptECB 修复版的AES ECB解密
func AesDecryptECB(content string, key string) (string, error) {
	// 1. Base64解码
	encryptedData, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return "", err
	}

	// 2. 验证密钥长度
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", fmt.Errorf("invalid key length: %d, must be 16, 24, or 32", len(key))
	}

	// 3. 创建密码块
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// 4. 验证数据长度是块大小的倍数
	if len(encryptedData)%aes.BlockSize != 0 {
		return "", fmt.Errorf("encrypted data length is not multiple of block size")
	}

	// 5. 逐块解密
	decrypted := make([]byte, len(encryptedData))
	for i := 0; i < len(encryptedData); i += aes.BlockSize {
		block.Decrypt(decrypted[i:i+aes.BlockSize], encryptedData[i:i+aes.BlockSize])
	}

	// 6. 正确去除PKCS7填充（不是零填充）
	decrypted = pkcs7Unpad(decrypted)

	return string(decrypted), nil
}

// AesEncryptECB 修复版的AES ECB加密
func AesEncryptECB(content string, key string) (string, error) {
	// 1. 验证密钥长度
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", fmt.Errorf("invalid key length: %d, must be 16, 24, or 32", len(key))
	}

	// 2. PKCS7填充（标准填充方式）
	data := pkcs7Pad([]byte(content))

	// 3. 创建密码块
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// 4. 逐块加密
	encrypted := make([]byte, len(data))
	for i := 0; i < len(data); i += aes.BlockSize {
		block.Encrypt(encrypted[i:i+aes.BlockSize], data[i:i+aes.BlockSize])
	}

	// 5. Base64编码
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// PKCS7填充（标准填充方式）
func pkcs7Pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

// PKCS7去除填充
func pkcs7Unpad(src []byte) []byte {
	if len(src) == 0 {
		return src
	}

	padding := int(src[len(src)-1])
	if padding < 1 || padding > aes.BlockSize {
		return src
	}

	// 验证填充字节是否正确
	for i := len(src) - padding; i < len(src); i++ {
		if int(src[i]) != padding {
			return src // 填充无效，返回原始数据
		}
	}

	return src[:len(src)-padding]
}

func AesEncryptCBC(content string, key string) string {
	origData := []byte(content)
	k := []byte(key)
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(k)
	blockSize := block.BlockSize()                            // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize)              // 补全码
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize]) // 加密模式
	encrypted := make([]byte, len(origData))                  // 创建数组
	blockMode.CryptBlocks(encrypted, origData)                // 加密
	return hex.EncodeToString(encrypted)
}

func AesDecryptCBC(content string, key string) string {
	encrypted, _ := hex.DecodeString(content)
	k := []byte(key)
	block, _ := aes.NewCipher(k)                              // 分组秘钥
	blockSize := block.BlockSize()                            // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize]) // 加密模式
	decrypted := make([]byte, len(encrypted))                 // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)               // 解密
	decrypted = pkcs5UnPadding(decrypted)                     // 去除补全码
	return string(decrypted)
}

func DESEncryptCBC(content string, key string, iv string) string {
	block, _ := des.NewCipher([]byte(key))
	data := pkcs5Padding([]byte(content), block.BlockSize())
	dest := make([]byte, len(data))
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	blockMode.CryptBlocks(dest, data)
	return fmt.Sprintf("%x", dest)
}

func DESDecryptCBC(content string, key string, iv string) string {
	b, _ := hex.DecodeString(content)
	block, _ := des.NewCipher([]byte(key))
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	originData := make([]byte, len(b))
	blockMode.CryptBlocks(originData, b)
	origData := pkcs5UnPadding(originData)
	return string(origData)
}

func DESEncryptECB(content string, key string) string {
	block, _ := des.NewCipher([]byte(key))
	size := block.BlockSize()
	data := pkcs5Padding([]byte(content), size)
	if len(data)%size != 0 {
		return ""
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:size])
		data = data[size:]
		dst = dst[size:]
	}
	return fmt.Sprintf("%x", out)
}

func DESDecryptECB(content string, key string) string {
	b, _ := hex.DecodeString(content)
	block, _ := des.NewCipher([]byte(key))
	size := block.BlockSize()
	out := make([]byte, len(b))
	dst := out
	for len(b) > 0 {
		block.Decrypt(dst, b[:size])
		b = b[size:]
		dst = dst[size:]
	}
	return string(pkcs5UnPadding(out))
}

func RC4Encrypt(content string, key string) string {
	dest := make([]byte, len(content))
	cp, _ := rc4.NewCipher([]byte(key))
	cp.XORKeyStream(dest, []byte(content))
	return fmt.Sprintf("%x", dest)
}

func RC4Decrypt(content string, key string) string {
	b, _ := hex.DecodeString(content)
	dest := make([]byte, len(b))
	cp, _ := rc4.NewCipher([]byte(key))
	cp.XORKeyStream(dest, b)
	return string(dest)
}

func Base64Encrypt(content string) string {
	return base64.StdEncoding.EncodeToString([]byte(content))
}

func Base64Decrypt(content string) string {
	b, _ := base64.StdEncoding.DecodeString(content)
	return string(b)
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func padding(src []byte) []byte {
	paddingCount := aes.BlockSize - len(src)%aes.BlockSize
	if paddingCount == 0 {
		return src
	} else {
		return append(src, bytes.Repeat([]byte{byte(0)}, paddingCount)...)
	}
}
