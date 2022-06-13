package main

import (
	"crypto/rand"
	"encoding/ascii85"
	"encoding/base64"
	"flag"
)

func main() {
	var encoding = flag.String("encode", "b", "encode type: a[scii85] or b[ase64]")
	var length = flag.Int("len", 12, "length of password")
	flag.Parse()

	var buffer = make([]byte, *length)
	if _, err := rand.Read(buffer); err == nil {
		var outputs = make([]byte, *length*2)
		switch *encoding {
		case "a":
			{
				ascii85.Encode(outputs, buffer)
			}
		case "b":
			{
				base64.StdEncoding.Encode(outputs, buffer)
			}
		}
		println(string(outputs[:*length]))
	}

}
