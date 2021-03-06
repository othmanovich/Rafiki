package rafiki

import (
	"bytes"
	"code.google.com/p/go.crypto/openpgp"
	"code.google.com/p/go.crypto/openpgp/armor"
	"io/ioutil"
)

const blockType = "PGP SIGNATURE"

func EncryptString(key []byte, clearText string) (string, error) {

	encBuf := bytes.NewBuffer(nil)
	w, err := armor.Encode(encBuf, blockType, nil)
	if err != nil {
		return "", err
	}

	plaintext, err := openpgp.SymmetricallyEncrypt(w, key, nil, nil)
	if err != nil {
		return "", err
	}
	message := []byte(clearText)
	_, err = plaintext.Write(message)

	plaintext.Close()
	w.Close()

	return encBuf.String(), nil

}

func DecryptString(encryptionKey []byte, cypherText string) (string, error) {

	decbuf := bytes.NewBuffer([]byte(cypherText))
	result, err := armor.Decode(decbuf)
	if err != nil {
		return "", err
	}

	md, err := openpgp.ReadMessage(result.Body, nil, func(keys []openpgp.Key, symmetric bool) ([]byte, error) {
		return encryptionKey, nil
	}, nil)
	if err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadAll(md.UnverifiedBody)
	return string(bytes), nil

}
