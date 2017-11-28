package encryption
import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	// "os"
	"io"
	// "io/ioutil"
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

func encodeBase64(b []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(b))
}

func decodeBase64(b []byte) ([]byte,error) {
	data, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		// fmt.Printf("Error: Bad Key!\n")
		// os.Exit(0)
		return []byte(""),err
	}
	return data,nil
}

func encrypt(key, text []byte) ([]byte,error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""),err
	}
	b := encodeBase64(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return []byte(""),err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], b)
	return ciphertext,nil
}
func decrypt(key, text []byte) ([]byte,error){
	block, err := aes.NewCipher(key)
	if err != nil {
		return []byte(""),err
	}
	if len(text) < aes.BlockSize {
		return []byte(""),errors.New("len(text) < aes.BlockSize")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	// fmt.Println("55")
	cfb.XORKeyStream(text, text)
	// fmt.Println("55")
	return decodeBase64(text)
}
func byteSliceEqual(a, b []byte) bool {
	// fmt.Println("a:",a)
	// fmt.Println("b:",b)
    if len(a) != len(b) {
    	// fmt.Println("len")
        return false
    }

    if (a == nil) != (b == nil) {
    	// fmt.Println("nil")
        return false
    }

    for i, v := range a {
        if v != b[i] {
        	// fmt.Println("range")
            return false
        }
    }

    return true
}
//用密码对masterkey加密，对加密后的文本在app端保存
func Encrypt(keystr, textstr string) (string,error) {
	// key, err := hex.DecodeString([]byte(keystr))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return ""
	// }
	// text, err := hex.DecodeString(textstr)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return ""
	// }
	key:=[]byte(keystr)
	text:=[]byte(textstr)
	hashKey:=sha256.Sum256(key)
	// fmt.Println("hashKey:",hashKey)

	// suffix:=sha256.Sum256(hashKey[:])
	encryptStr:=make([]byte,len(text)+len(hashKey))
	
	copy(encryptStr[:len(text)],text[:])
	copy(encryptStr[len(text):],hashKey[:])

	suffix:=sha256.Sum256(encryptStr)
	constEncKey:=sha256.Sum256([]byte("K3d9R9oDAj9j1PkbWuUkqi4TT2RgWqTTvgmahbNW9cxccRhEWS"))
	prefix,err:=encrypt(constEncKey[:],encryptStr)
	if err!=nil{
		return "",err
	}
	// fmt.Println("prefix:",prefix)
	// fmt.Println("suffix:",suffix)

	ret:=make([]byte,len(prefix)+len(suffix))
	copy(ret[:len(prefix)],prefix[:])
	copy(ret[len(prefix):],suffix[:])
	// fmt.Println("ret:",ret)
	return hex.EncodeToString(ret[:]),nil
}

//用密码对密文解密返回masterkey对应的byte数组
func Decrypt(textstr string) (string,error){
	// key, err := hex.DecodeString(keystr)
	// if err != nil {
	// 	return ""
	// }
	// key:=[]byte(keystr)
	text, err := hex.DecodeString(textstr)
	if err != nil {
		return "",err
	}
	// fmt.Println("text:",text)
	// hashKey:=sha256.Sum256(key)
	constEncKey:=sha256.Sum256([]byte("K3d9R9oDAj9j1PkbWuUkqi4TT2RgWqTTvgmahbNW9cxccRhEWS"))
	d_des,err:=decrypt(constEncKey[:], text[:len(text)-len(constEncKey)])
	if err!=nil{
		return "",err
	}
	return string(d_des[:len(d_des)-len(constEncKey)]),nil
}
//用文本来验证密码是否正确
func Validate(keystr, textstr string) bool {
	dec,err:=Decrypt(textstr)
	if err!=nil{
		return false
	}
	d_des:=[]byte(dec)
	key:=[]byte(keystr)
	text, err := hex.DecodeString(textstr)
	if err != nil {
		return false
	}
	hashKey:=sha256.Sum256(key)
	// fmt.Println("hashKey:",hashKey)

	// suffix:=sha256.Sum256(hashKey[:])

	// fmt.Println("suffix:",suffix)

	// d_des,err:=decrypt(hashKey[:], text[:len(text)-len(hashKey)])
	// if err!=nil{
	// 	return false
	// }
	hashstr:=make([]byte,len(d_des)+len(hashKey))
	copy(hashstr[:len(d_des)],d_des[:])
	copy(hashstr[len(d_des):],hashKey[:])

	hash_des:=sha256.Sum256(hashstr[:])
	// fmt.Println("d_des:",d_des)
	// fmt.Println("d_des:",hex.EncodeToString(d_des[:len(d_des)-len(suffix)]))
	fmt.Println("d_des:",dec)
	return byteSliceEqual(text[len(text)-len(hashKey):],hash_des[:])
}