package main

import (
	"crypto/rand"
	"os"
	"strconv"
)

func main() {
	const ALPHABET = `ABCDEFGHJKLMNPQRSTUVWXYZ0123456789abcdefghijkmnopqrstuvwxyz0123456789@#$%^&*,./_-+=0123456789`
	const MOD = len(ALPHABET)
	// set length
	var length = 12
	if len(os.Args) >= 2 {
		var res, err = strconv.ParseInt(os.Args[1], 0, 0)
		if err == nil {
			if res > 6 {
				length = int(res)
			} else {
				length = 6
				println(">> The minimum length is 6 (actually processed as 6)\n")
			}
		} else {
			println(">> Usage: passer [length=12]\n")
		}
	}

	// gen password
	var password = ""
	var buffer = make([]byte, length)
	var iv_buf = make([]byte, 1)
	rand.Read(buffer)
	rand.Read(iv_buf)
	var iv = MOD - (int(iv_buf[0]) % MOD)
	for _, n := range buffer {
		var i = (int(n) + iv) % MOD
		iv = MOD - i
		password += string(ALPHABET[i])
	}
	println(password)
}
