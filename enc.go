package main

import (
	"bytes"
	"encoding/ascii85"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/mtraver/base91"
	"github.com/raiiga/flagg"
)

type encodeOptions struct {
	/*	base32  string       `flagg:"n:32, name: base32, usage: Encode text to base32"`
		base58  string       `flagg:"n:58, name: base58, usage: Encode text to base58"`
	*/
	url    bool               `flagg:"n:u, name: url, usage: Encoding mode for base64"`
	raw    bool               `flagg:"n:r, name: raw, usage: Encoding mode for base64"`
	base64 func(string) error `flagg:"n:64, name: base64, usage: Encode text to base64"`
	base85 func(string) error `flagg:"n:85, name: base85, usage: Encode text to base85"`
	base91 func(string) error `flagg:"n:91, name: base91, usage: Encode text to base91"`
	/*	hex     bool               `flagg:"n:x, name: hex, usage: Encoding mode for base32"`
		flickr  bool               `flagg:"n:F, name: flickr, usage: Encoding mode for base58"`
		bitcoin bool               `flagg:"n:B, name: bitcoin, usage: Encoding mode for base58"`
		raw     bool               `flagg:"n:r, name: raw, usage: Encoding mode for base64"`
		url     bool               `flagg:"n:u, name: url, usage: Encoding mode for base64"`*/
	//decode  bool         `flagg:"n:D, name: decode, usage: Decode mode"`
	//base16  func(string) error `flagg:"n:16, name: base16, usage: Encode text to base16"`
}

func main() {
	flags, opts := flagg.New("enc"), new(encodeOptions)

	opts.base64 = func(s string) error {
		var enc *base64.Encoding

		switch {
		case opts.url && opts.raw:
			enc = base64.RawURLEncoding
		case opts.url:
			enc = base64.URLEncoding
		case opts.raw:
			enc = base64.RawStdEncoding
		default:
			enc = base64.StdEncoding
		}

		fmt.Println(enc.EncodeToString([]byte(s)))
		return nil
	}

	opts.base85 = func(s string) error {
		buffer, err := bytes.NewBuffer(make([]byte, len(s))), error(nil)
		encoder := ascii85.NewEncoder(buffer)

		defer func() {
			if errors.Join(err, encoder.Close()) != nil {
				fmt.Println("system error..")
			} else {
				fmt.Println(buffer.String())
			}
		}()

		if _, err = encoder.Write([]byte(s)); err != nil {
			return err
		}

		return nil
	}

	opts.base91 = func(s string) error {
		fmt.Println(base91.StdEncoding.EncodeToString([]byte(s)))
		return nil
	}

	if _, _err := flags.Map(opts); _err != nil {
		return
	}
}
