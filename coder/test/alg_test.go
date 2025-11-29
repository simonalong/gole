package test

import (
	"crypto/dsa"
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
	"github.com/simonalong/gole/coder"
	"github.com/simonalong/gole/file"
)

func TestMd5(t *testing.T) {
	str := "abcdefg"
	s1 := coder.MD5String(str)
	t.Logf("%s md5 is %s", str, s1)

	s2, err := coder.MD5File(f)
	t.Logf("%s md5 is %s (%v)", f, s2, err)
}

func TestHMac(t *testing.T) {
	key := "simonalong"
	str := "abcdefg"
	s1 := coder.HMacMD5String(str, key)
	t.Logf("%s hmac is %s", str, s1)
}

func TestRC4(t *testing.T) {
	key := "simonalong"
	str := "abcdefg"
	s1 := coder.RC4Encrypt(str, key)
	t.Logf("%s rc4 is %s", str, s1)

	s2 := coder.RC4Decrypt(s1, key)
	t.Logf("%s rc4 is %s", s1, s2)
}

//func TestDES(t *testing.T) {
//	// CBC
//	key := "simonalong"
//	iv := "12345678"
//	str := "abcdefg"
//	s1 := coder.DESEncryptCBC(str, key, iv)
//	t.Logf("%s des is %s", str, s1)
//	s2 := coder.DESDecryptCBC(s1, key, iv)
//	t.Logf("%s des is %s", s1, s2)
//
//	// ECB
//	ss1 := coder.DESEncryptECB(str, key)
//	t.Logf("%s des is %s", str, ss1)
//	ss2 := coder.DESDecryptECB(ss1, key)
//	t.Logf("%s des is %s", ss1, ss2)
//}

func TestRSA(t *testing.T) {
	privKeyPath := "/Users/zhouzhenyong/project/seatak/gole/coder/test/rsa/private.pem"
	pubKeyPath := "/Users/zhouzhenyong/project/seatak/gole/coder/test/rsa/public.pem"
	if !file.FileExists(privKeyPath) {
		err := coder.RSAGenerateKeyPair(coder.RSA_KEY_SIZE_1024, privKeyPath, pubKeyPath)
		if err != nil {
			t.Logf("generate key pair error: %v", err)
			return
		}
	}

	str := "abcdefg"
	s1, err := coder.RSAEncryptByPath(str, pubKeyPath)
	if err != nil {
		t.Logf("encrypt error: %v", err)
		return
	}
	t.Logf("%s rsa is %s", str, s1)

	s2, err := coder.RSADecryptByPath(s1, privKeyPath)
	t.Logf("%s rsa is %s (%v)", s1, s2, err)

	assert.Equal(t, s2, str)
}

func TestRsaData(t *testing.T) {
	pubKeyData := "-----BEGIN RSA Public Key-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCyKlXDTKanPUo4N7mtl3zy+COW\nOHvGOu2NrtQVUgvBOypShuk2yqEBMW2rbGnV+JmMLV+GCzwjb/07Unf1g/kWlweF\nzPahZlusRVKvS768kPXjWT/T+WJf+AeShn6TTgwct2hvjmZOwKWEe/9L2Qen41NJ\np5UC5UXHStVgpEQ5owIDAQAB\n-----END RSA Public Key-----"
	str := "abcdefg"
	s1, err := coder.RSAEncryptByData(str, []byte(pubKeyData))
	if err != nil {
		t.Logf("encrypt error: %v", err)
		return
	}

	priKeyData := "-----BEGIN RSA Private Key-----\nMIICWwIBAAKBgQCyKlXDTKanPUo4N7mtl3zy+COWOHvGOu2NrtQVUgvBOypShuk2\nyqEBMW2rbGnV+JmMLV+GCzwjb/07Unf1g/kWlweFzPahZlusRVKvS768kPXjWT/T\n+WJf+AeShn6TTgwct2hvjmZOwKWEe/9L2Qen41NJp5UC5UXHStVgpEQ5owIDAQAB\nAoGALrL+C9TZkdh0zct9dczRSXZVDZj8iHcFsS90E6qPvjRd4YfPNTdjgEaOcnJW\nQ2mIBcAW27GyL6+49oWlP8s5zIv4nB7Dw/zAT2Up9sKR6IY7Xoz8MJbvyZRGg5Ep\nqowMt1r1gqWsFpF64CdWa2xGyTXydCnbQOX+lvHn1oeQONECQQDNAppwip3itAHO\nB+63hgMlbfPehKagDZNNL/KxtCd7kfQbb2BtjynxV9V4QpHS8SudBvn6oFQEvbLO\nZIsSbspfAkEA3np4nnFlt0YUjUSsTaHNauB55CR07NVX28B3v183As2QP4DMrI7U\n95Q9J3ctztrhyAv1mTuy0pcWiQluqPOfPQJADovCvX14Wl9/SUkSzP67NmqoxP8Q\ne4a7Dtz6EVXA/2mJsnCinONtjGw4/0Fp61elSoz2K6w4ieWTzEUiAPrPbQJAXP9u\n2jRmo2TNBHxXViAzoOBys1Y19iX8EuTyaXGgqjBJgvIRHHScO12g7pVX9abzSE8P\ne91Dk9oKVoA13LPxtQJAN5PEfLYaeGt+pabKnf73VeH7Y7gKHkUQC6kQmTkYuaUe\n15lavY2AXGykbJEKLkBkzTsYCfrV/FXX7TrlRJf1fQ==\n-----END RSA Private Key-----"
	s2, err := coder.RSADecryptByData(s1, []byte(priKeyData))

	assert.Equal(t, s2, str)
}

func TestGenerateKey(t *testing.T) {
	priKey, pubKey, _ := coder.RSAGenerateKey(coder.RSA_KEY_SIZE_1024)

	fmt.Println(priKey)
	fmt.Println("=========")
	fmt.Println(pubKey)

	oldContent := "abcdefg"
	codeData, _ := coder.RSAEncryptByData(oldContent, []byte(pubKey))
	newContent, _ := coder.RSADecryptByData(codeData, []byte(priKey))

	assert.Equal(t, newContent, oldContent)
}

func TestAes(t *testing.T) {
	// encrypt
	content := "abcdefg"
	// PKI-BRIDGE~1

	// key的长度必须是 16，24，32
	key := "simonalong123456"
	iv := "0102030405060708"
	s1 := coder.AesEncrypt(content, key, iv)
	t.Logf("%s aes is %s", content, s1)

	// decrypt
	s2 := coder.AesDecrypt(s1, key, iv)
	t.Logf("%s aes is %s", s1, s2)
}

func TestAesJava(t *testing.T) {
	content := "abcdefg"
	key := "simonalong123456"
	s1 := coder.AesEncryptECBForJava(content, key)
	t.Logf("%s aes is %s", content, s1)
	s2 := coder.AesDecryptECBForJava(s1, key)
	fmt.Println(fmt.Sprintf("%s aes is %v", s1, s2))
}
func TestAesCBC(t *testing.T) {
	content := "abcdefg"
	key := "simonalong123456"
	s1 := coder.AesEncryptCBC(content, key)
	t.Logf("%s aes is %s", content, s1)
	s2 := coder.AesDecryptCBC(s1, key)
	t.Logf("%s aes is %s", s1, s2)
}
