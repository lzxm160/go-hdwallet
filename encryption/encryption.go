package encryption
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"io"
	// "io/ioutil"
	"crypto/sha256"
	"encoding/hex"
)

func encodeBase64(b []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(b))
}

func decodeBase64(b []byte) []byte {
	data, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		fmt.Printf("Error: Bad Key!\n")
		os.Exit(0)
	}
	return data
}

func encrypt(key, text []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	b := encodeBase64(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], b)
	return ciphertext
}
func decrypt(key, text []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(text) < aes.BlockSize {
		fmt.Printf("Error!\n")
		os.Exit(0)
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	return decodeBase64(text)
}
func byteSliceEqual(a, b []byte) bool {
    if len(a) != len(b) {
        return false
    }

    if (a == nil) != (b == nil) {
        return false
    }

    for i, v := range a {
        if v != b[i] {
            return false
        }
    }

    return true
}
//用密码对masterkey加密，对加密后的文本在app端保存
func Encrypt(key, text []byte) []byte {

	hashKey:=sha256.Sum256(key)
	fmt.Println("hashKey:",hashKey)

	// suffix:=sha256.Sum256(hashKey[:])
	prefix:=encrypt(hashKey[:],text)
	suffix:=sha256.Sum256(prefix)
	fmt.Println("prefix:",prefix)
	fmt.Println("suffix:",suffix)

	ret:=make([]byte,len(prefix)+len(suffix))
	copy(ret[:len(prefix)],prefix[:])
	copy(ret[len(prefix):],suffix[:])
	// fmt.Println("ret:",ret)
	return ret
}
//用密码对密文解密返回masterkey对应的byte数组
func Decrypt(key, text []byte) []byte {
	hashKey:=sha256.Sum256(key)
	d_des:=decrypt(hashKey[:], text[:len(text)-len(hashKey)])
	return d_des
}
//用文本来验证密码是否正确
func Validate(key, text []byte) bool {
	hashKey:=sha256.Sum256(key)
	fmt.Println("hashKey:",hashKey)

	// suffix:=sha256.Sum256(hashKey[:])
	
	// fmt.Println("suffix:",suffix)

	d_des:=decrypt(hashKey[:], text[:len(text)-len(hashKey)])
	fmt.Println("d_des:",d_des)
	// fmt.Println("d_des:",hex.EncodeToString(d_des[:len(d_des)-len(suffix)]))
	fmt.Println("d_des:",string(d_des))
	return byteSliceEqual(text[:len(text)-len(hashKey)],d_des[len(d_des)-len(suffix):])
}