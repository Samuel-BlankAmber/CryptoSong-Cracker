package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"sync"
)

func parseEncryptedText(encryptedB64 string) ([]byte, []byte, []byte, error) {
	encrypted, err := base64.StdEncoding.DecodeString(encryptedB64)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("Base64 decoding error: %v", err)
	}

	nonce := encrypted[:12]
	tag := encrypted[12:28]
	ciphertext := encrypted[28:]
	return nonce, tag, ciphertext, nil
}

func decryptAesGcm(key, nonce, ciphertext, tag []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("NewCipher error: %v", err)
	}

	aesGcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("NewGCM error: %v", err)
	}

	ciphertextWithTag := append(ciphertext, tag...)
	plaintext, err := aesGcm.Open(nil, nonce, ciphertextWithTag, nil)
	if err != nil {
		return nil, fmt.Errorf("Open error: %v", err)
	}

	return plaintext, nil
}

func processKey(key string, nonce, ciphertext, tag []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	aesKey := sha256.Sum256([]byte(key))
	plaintext, err := decryptAesGcm(aesKey[:], nonce, ciphertext, tag)
	if err != nil {
		return
	}
	result := fmt.Sprintf("Decrypted: '%s' with key: '%s'", plaintext, key)
	fmt.Println(result)
}

func main() {
	encryptedB64 := flag.String("encrypted", "", "Base64 encoded encrypted text")
	flag.Parse()

	if *encryptedB64 == "" {
		fmt.Println("Error: encrypted text is required")
		return
	}

	nonce, tag, ciphertext, err := parseEncryptedText(*encryptedB64)
	if err != nil {
		fmt.Println("Error processing encrypted text:", err)
		return
	}

	file, err := os.Open("keys.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var wg sync.WaitGroup
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		key := scanner.Text()
		wg.Add(1)
		go processKey(key, nonce, ciphertext, tag, &wg)
	}

	wg.Wait()

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
