package password

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type ResetToken struct {
	Email  string
	Expire int64
}

const key = "asdfghjkl;zxcvbnm,"

// var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func GenerateToken(email string) (string, error) {
	var resettoken ResetToken
	resettoken.Email = email
	resettoken.Expire = time.Now().Unix() + 86400

	jsonData, err := json.Marshal(resettoken)
	if err != nil {
		return "", err
	}

	token, err := EncryptionToken(jsonData)
	if err != nil {
		return "", err
	}
	return token, nil
}

func EncryptionToken(plaintext []byte) (string, error) {
	var iv []byte
	var ciphertext = make([]byte, len(plaintext))

	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key), err)
		return "", err
	}

	if iv, err = loadIV(); err != nil {
		iv = generateIV()
	}
	err = saveIV(iv)
	if err != nil {
		return "", err
	}

	cbc := cipher.NewCBCDecrypter(c, iv)
	cbc.CryptBlocks(ciphertext, plaintext)

	cb64 := base64.StdEncoding.EncodeToString(ciphertext)

	return cb64, nil
}

func DecryptionToken(cb64 string) (resettoken ResetToken, err error) {
	decodedCiphertext, err := base64.StdEncoding.DecodeString(cb64)
	if err != nil {
		return resettoken, err
	}

	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(key), err)
		return resettoken, err
	}

	decrypted := make([]byte, len(decodedCiphertext))

	if iv, err := loadIV(); err == nil {
		cbc := cipher.NewCBCEncrypter(c, iv)
		cbc.CryptBlocks(decrypted, []byte(decodedCiphertext))

		err = json.Unmarshal(decrypted, &resettoken)
		if err != nil {
			return resettoken, nil
		}
	}

	return resettoken, nil
}

func generateIV() []byte {
	iv := make([]byte, aes.BlockSize)
	_, err := rand.Read(iv)
	if err != nil {
		panic(err)
	}
	return iv
}

func saveIV(iv []byte) error {
	ivb64 := base64.StdEncoding.EncodeToString(iv)

	f, err := os.OpenFile(".env", os.O_WRONLY|os.O_CREATE|os.O_RDWR, 0760)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(ivb64)
	if err != nil {
		return err
	}

	return nil
}

func loadIV() ([]byte, error) {
	f, err := os.Open(".env")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	iv, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	return iv, err
}
