package main

import (
	"fmt"
	"github.com/raiiga/flagg"
)

type decodeOptions struct {
	base16 func(string) error `flagg:"n:16, name: base16, usage: Decode text from base16"`
	base32 func(string) error `flagg:"n:32, name: base32, usage: Decode text from base32"`
	base58 func(string) error `flagg:"n:58, name: base58, usage: Decode text from base58"`

	base64 func(string) error `flagg:"n:64, name: base64, usage: Decode text from base64"`
	base85 func(string) error `flagg:"n:85, name: base85, usage: Decode text from base85"`
	base91 func(string) error `flagg:"n:91, name: base91, usage: Decode text from base91"`

	hex bool `flagg:"n:x, name: hex, usage: Encoding mode for base32"`
	url bool `flagg:"n:u, name: url, usage: Encoding mode for base64"`
	raw bool `flagg:"n:r, name: raw, usage: Encoding mode for base64"`
}

func main() {
	opts := new(decodeOptions)

	opts.base16 = func(s string) error {
		e, err := base16decode(s)

		if err != nil {
			return err
		}

		fmt.Println(e)
		return nil
	}

	opts.base32 = func(s string) error {
		e, err := base32decode(s, opts.hex)

		if err != nil {
			return err
		}

		fmt.Println(e)
		return nil
	}

	opts.base58 = func(s string) error {
		e, err := base58decode(s)

		if err != nil {
			return err
		}

		fmt.Println(e)
		return nil
	}

	opts.base64 = func(s string) error {
		e, err := base64decode(s, opts.url, opts.raw)

		if err != nil {
			return err
		}

		fmt.Println(e)
		return nil
	}

	opts.base85 = func(s string) error {
		e, err := base85decode(s)

		if err != nil {
			return err
		}

		fmt.Println(e)
		return nil
	}

	opts.base91 = func(s string) error {
		e, err := base91decode(s)

		if err != nil {
			return err
		}

		fmt.Println(e)
		return nil
	}

	flagg.New("dec").Map(opts)
}
