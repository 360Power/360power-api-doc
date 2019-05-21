package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	pemString := `-----BEGIN RSA PRIVATE KEY-----
MIICWgIBAAKBgFUIytlnjO9kbfSXh0D8Rkar79Nblt6sWi2SLJqMpyxlzqKrzhkW
LpEgtaCmfgUyDlxwpL38waWXRA4BHVvzRUztvH4e3gObjwZxenXpl8Au5Sc85sm6
mnyV2StjeYeWOKDyJ87/nBC8gNaMb65Z38kPmLuFESvCszmEklxRqL6xAgMBAAEC
gYA6CN4osnuFhs1keWZd+88avI3ZelDleEuzfmfisswFiRYV/5uRk4oEkoZjNj4b
3aXfgSFuaOrg0PQpeqlG8CkDJnhGEe5t4GNQQOGDI2fnQ7UXAjQSFtISJQu9I8Oz
wxRHr+B81trIyzLja+AYGrDm3/1SSBAy5+292XyaJW80gQJBAJRNPpCUtsALqcby
wUhZAU2GhdLv7tZJPTSxQLrt2vB/tw1XPC8hTOxlvg2lOBjerfyLxcYhOpT6E3lb
coq69mUCQQCSyYQWbwfQ7egcq044U+JkHWm9av6LSC0RxZj5xLqS5zwyVSXvQEu5
DbAPaiWydf5EzEtiVwoWI4bMSbYSJoxdAkBSrtpyC6f0XMxUkqX2q0ERsy3LlGAp
8v1/8k9vqQuHSP2LH5b7g+p6ZqNWwkYLf6OriVZEB+S8iMzwvW6YMHMNAkBS0T+l
KJ/QUWpUQpKvVSS2N6IhLOzQyLgk/seApG5f0/cyrrfodO5ESmS7TbhXKBt91YXy
xgj61LCJMk13kChBAkBwXwAxM5c0qPMbLs2mKDQbqb6KYgcFQZOjsj8u3T6zQvnX
4jXF5U9sgNwyC/2IYVJvMAh9hXlFtEeGS3w2XL2M
-----END RSA PRIVATE KEY-----`

	rawPrivKey, _ := pem.Decode([]byte(pemString))
	privKey, err := x509.ParsePKCS1PrivateKey(rawPrivKey.Bytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from parsing priv key: %s\n", err)
		return
	}

	toEncrypt := []byte("GET999999999999999")
	shaDigest := sha256.Sum256(toEncrypt)

	rng := rand.Reader

	signature, err := rsa.SignPKCS1v15(rng, privKey, crypto.Hash(0), shaDigest[:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
		return
	}

	fmt.Println(base64.StdEncoding.EncodeToString(signature))

}
