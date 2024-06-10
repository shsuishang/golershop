package utility

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/farmerx/gorsa"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"log"
)

func RsaDecrypt(encrypted string, privateKey string) (oriText string, err error) {
	cipherText, _ := base64.StdEncoding.DecodeString(encrypted)

	//1、block pem格式
	//privateKey, _ := ioutil.ReadFile("private.pem")
	//bytePrivateKey := []byte(privateKey)

	//2、单行数据格式
	decodeString, err := base64.StdEncoding.DecodeString(privateKey)

	// 创建一个 PEM 格式的块
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: []byte(decodeString),
	}

	// 编码 PEM 块为 PEM 格式的字符串
	bytePrivateKey := pem.EncodeToMemory(block)

	priKey, _ := pem.Decode(bytePrivateKey)

	if priKey == nil {
		return oriText, errors.New("private key error!")
	}

	//fmt.Println(fmt.Printf("%v ", gconv.String(priKey)))

	priv, err := x509.ParsePKCS1PrivateKey(priKey.Bytes)

	if err != nil {
		//尝试是否为 ParsePKCS8PrivateKey
		pkcs8PrivateKey, err := x509.ParsePKCS8PrivateKey(priKey.Bytes)

		if err != nil {
			return oriText, err
		}

		switch pkcs8PrivateKey.(type) {
		case *rsa.PrivateKey:
			priv = pkcs8PrivateKey.(*rsa.PrivateKey)
		case *ecdsa.PrivateKey:
			// ...
		case ed25519.PrivateKey:
			// ...
		default:
			panic("unknown key")
		}
	}

	v15Bytes, err := rsa.DecryptPKCS1v15(rand.Reader, priv, cipherText)

	return gconv.String(v15Bytes), err
}

// 获取公钥
func PublicKey(ctx context.Context) string {
	// 附件存储路径
	publicKey, _ := g.Cfg().Get(ctx, "secure.publicKey")

	return publicKey.String()
}

// 获取私钥
func PrivateKey(ctx context.Context) string {
	// 附件存储路径
	privateKey, _ := g.Cfg().Get(ctx, "secure.privateKey")

	return privateKey.String()
}

func test() {
	//生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	//生成公钥
	publicKey := privateKey.PublicKey
	//根据公钥加密
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&publicKey,
		[]byte("测试哈哈哈"), //需要加密的字符串
		nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("encrypted bytes: ", encryptedBytes)
	//根据私钥解密
	decryptedBytes, err := privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}
	fmt.Println("decrypted message: ", string(decryptedBytes))
}

func PubkeyDecrypt(encrypted string, publicKey string) (oriText string, err error) {
	//2、单行数据格式
	decodeString, err := base64.StdEncoding.DecodeString(publicKey)

	// 创建一个 PEM 格式的块
	block := &pem.Block{
		Type:  "公钥",
		Bytes: decodeString,
	}

	// 编码 PEM 块为 PEM 格式的字符串
	publicStr := pem.EncodeToMemory(block)

	if err := gorsa.RSA.SetPublicKey(gconv.String(publicStr)); err != nil {
		log.Fatalln(`set public key :`, err)
	}

	cipherText, _ := base64.StdEncoding.DecodeString(encrypted)
	pubdecrypt, err := gorsa.RSA.PubKeyDECRYPT(cipherText)
	if err != nil {
		return "", err
	}

	return gconv.String(pubdecrypt), err
}
