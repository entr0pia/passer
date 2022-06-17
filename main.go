package main

import (
	"crypto/rand"
	"encoding/ascii85"
	"encoding/base64"
	"errors"
	"github.com/akamensky/argparse"
	"os"
	"strconv"
)

func main() {
	// init parser
	var parser = argparse.NewParser("passer", "A simple random password generator")
	var encoding = parser.String("e", "encode", &argparse.Options{Help: "Encode type: a[scii85] or b[ase64]", Default: "b"})
	var length = parser.Int("l", "length", &argparse.Options{Help: "Length of password", Default: 12, Validate: func(args []string) error {
		if l, err := strconv.Atoi(args[0]); l < 6 || err != nil {
			return errors.New("the length must be >= 6")
		}
		return nil
	}})

	// err massage
	if err := parser.Parse(os.Args); err != nil {
		println(parser.Usage(err))
		return
	}

	// gen password
	var buffer = make([]byte, *length)
	if _, err := rand.Read(buffer); err == nil {
		var outputs = make([]byte, *length*2)
		switch (*encoding)[0] {
		case 0x41, 0x61:
			ascii85.Encode(outputs, buffer)
		case 0x42, 0x62:
			// as default
			fallthrough
		default:
			base64.StdEncoding.Encode(outputs, buffer)
		}
		println(string(outputs[:*length]))
	}

}
