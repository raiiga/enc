package main

import (
	"bytes"
	"encoding/ascii85"
	"encoding/base32"
	"encoding/base64"
	"encoding/hex"
	"github.com/akamensky/base58"
	"github.com/mtraver/base91"
	"io"
)

func base32encoding(hex bool) *base32.Encoding {
	switch {
	case hex:
		return base32.HexEncoding
	default:
		return base32.StdEncoding
	}
}

func base64encoding(url, raw bool) *base64.Encoding {
	switch {
	case url && raw:
		return base64.RawURLEncoding
	case url:
		return base64.URLEncoding
	case raw:
		return base64.RawStdEncoding
	default:
		return base64.StdEncoding
	}
}

/**/

func base16encode(s string) (string, error) {
	return hex.EncodeToString([]byte(s)), nil
}

func base32encode(s string, hex bool) (string, error) {
	enc := base32encoding(hex)
	return enc.EncodeToString([]byte(s)), nil
}

func base58encode(s string) (string, error) {
	return base58.Encode([]byte(s)), nil
}

func base64encode(s string, url, raw bool) (string, error) {
	enc := base64encoding(url, raw)
	return enc.EncodeToString([]byte(s)), nil
}

func base85encode(s string) (string, error) {
	buffer := bytes.NewBuffer(make([]byte, len(s)))
	encoder := ascii85.NewEncoder(buffer)

	if _, err := io.WriteString(encoder, s); err != nil {
		return "", err
	}

	if err := encoder.Close(); err != nil {
		return "", err
	}

	return buffer.String(), nil
}

func base91encode(s string) (string, error) {
	return base91.StdEncoding.EncodeToString([]byte(s)), nil
}

/**/

func base16decode(s string) (string, error) {
	d, err := hex.DecodeString(s)

	if err != nil {
		return "", err
	}

	return bytes.NewBuffer(d).String(), nil
}

func base32decode(s string, hex bool) (string, error) {
	enc := base32encoding(hex)

	d, err := enc.DecodeString(s)
	if err != nil {
		return "", err
	}

	return bytes.NewBuffer(d).String(), nil
}

func base58decode(s string) (string, error) {
	d, err := base58.Decode(s)

	if err != nil {
		return "", err
	}

	return bytes.NewBuffer(d).String(), nil
}

func base64decode(s string, url, raw bool) (string, error) {
	enc := base64encoding(url, raw)

	d, err := enc.DecodeString(s)
	if err != nil {
		return "", err
	}

	return bytes.NewBuffer(d).String(), nil
}

func base85decode(s string) (string, error) {
	buffer := bytes.NewBufferString(s)
	decoder := ascii85.NewDecoder(buffer)

	d, err := io.ReadAll(decoder)
	if err != nil {
		return "", err
	}

	return bytes.NewBuffer(d).String(), nil
}

func base91decode(s string) (string, error) {
	enc := base91.StdEncoding

	d, err := enc.DecodeString(s)
	if err != nil {
		return "", err
	}

	return bytes.NewBuffer(d).String(), nil
}
