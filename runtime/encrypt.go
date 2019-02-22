package lib

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

type Encrypt struct {
}

//char * md5(char *);
func (*Encrypt) Md5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}

//char * sha1(char *);
func (*Encrypt) Sha1(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

//char * sha256(char *);
func (*Encrypt) Sha256(s string) string {
	r := sha256.Sum256([]byte(s))
	return hex.EncodeToString(r[:])
}

//char * sha512(char *);
func (*Encrypt) Sha512(s string) string {
	r := sha512.Sum512([]byte(s))
	return hex.EncodeToString(r[:])
}

//char * base64_encode(char *);
func (*Encrypt) Base64_encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// char * base64_decode(char *)
func (*Encrypt) Base64_decode(s string) string {
	r, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		fmt.Println(err)
	}
	return string(r)
}
