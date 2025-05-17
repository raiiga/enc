package main

import (
	"fmt"
	"github.com/raiiga/flagg"
)

type encodeOptions struct {
	base16 func(string) error `flagg:"n:16, name: base16, usage: Encode text to base16"`
	base32 func(string) error `flagg:"n:32, name: base32, usage: Encode text to base32"`
	base58 func(string) error `flagg:"n:58, name: base58, usage: Encode text to base58"`

	base64 func(string) error `flagg:"n:64, name: base64, usage: Encode text to base64"`
	base85 func(string) error `flagg:"n:85, name: base85, usage: Encode text to base85"`
	base91 func(string) error `flagg:"n:91, name: base91, usage: Encode text to base91"`

	hex bool `flagg:"n:x, name: hex, usage: Encoding mode for base32"`
	url bool `flagg:"n:u, name: url, usage: Encoding mode for base64"`
	raw bool `flagg:"n:r, name: raw, usage: Encoding mode for base64"`
}

func main() {
	opts := new(encodeOptions)

	opts.base16 = func(s string) error {
		e, err := base16encode(s)

		if err != nil {
			return err
		}

		fmt.Println(e)
		return nil
	}

	opts.base32 = func(s string) error {
		e, err := base32encode(s, opts.hex)

		if err != nil {
			return err
		}

		fmt.Println(e)
		return nil
	}

	opts.base58 = func(s string) error {
		e, err := base58encode(s)

		if err != nil {
			return err
		}

		fmt.Println(e)
		return nil
	}

	opts.base64 = func(s string) error {
		e, err := base64encode(s, opts.url, opts.raw)

		if err != nil {
			return err
		}

		fmt.Println(e)
		return nil
	}

	opts.base85 = func(s string) error {
		e, err := base85encode(s)

		if err != nil {
			return err
		}

		fmt.Println(e)
		return nil
	}

	opts.base91 = func(s string) error {
		e, err := base91encode(s)

		if err != nil {
			return err
		}

		fmt.Println(e)
		return nil
	}

	flagg.New("enc").Map(opts)
}
