package main

type decodeOpts struct {
	base32  bool         `flagg:"n:32, name: base32, usage: Decode text from base32"`
	base58  bool         `flagg:"n:58, name: base58, usage: Decode text from base58"`
	base64  bool         `flagg:"n:64, name: base64, usage: Decode text from base64"`
	hex     bool         `flagg:"n:x, name: hex, usage: Encoding mode for base32"`
	flickr  bool         `flagg:"n:F, name: flickr, usage: Encoding mode for base58"`
	bitcoin bool         `flagg:"n:B, name: bitcoin, usage: Encoding mode for base58"`
	raw     bool         `flagg:"n:r, name: raw, usage: Encoding mode for base64"`
	url     bool         `flagg:"n:u, name: url, usage: Encoding mode for base64"`
	help    func() error `flagg:"n:h, name: help, usage: Show this message and exit"`
}

func main() {
	arr := []int{4, 2, 9, 1, -1, -1, 2, 7, 3, 5}
	arr = quick(arr)
	for i := 0; i < len(arr); i++ {
		print(arr[i], " ")
	}
}
