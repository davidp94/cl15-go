package main

import (
	"crypto/rand"
	"log"
	"math/big"
)

func generateLargePrime(k int) (p *big.Int, err error) {
	return rand.Prime(rand.Reader, k)
}

func main() {
	log.Println("CL15 Golang implementation")
	log.Println("Credits to Prof Buchanan")
	bits := 128

	p, err := generateLargePrime(bits)
	if err != nil {
		log.Fatalln("could not generate p")
	}

	q, err := generateLargePrime(bits)
	if err != nil {
		log.Fatalln("could not generate q")
	}

	n := new(big.Int).Mul(p, q)

	e := new(big.Int).SetInt64(65537)

	phiL := new(big.Int).Sub(p, big.NewInt(1))
	phiR := new(big.Int).Sub(q, big.NewInt(1))

	phi := new(big.Int).Mul(phiL, phiR)

	d := new(big.Int).ModInverse(e, phi)

	uid := new(big.Int).SetInt64(43)
	xa := new(big.Int).SetInt64(11)
	xb := new(big.Int).SetInt64(10)

	id1 := new(big.Int).Exp(uid, xa, n)
	id2 := new(big.Int).Exp(uid, xb, n)

	log.Println("id 1", id1)
	log.Println("id 2", id2)

	log.Println("xa", xa)
	log.Println("xb", xb)

	log.Println("e", e)

	log.Println("N", n)

	log.Println("p", p)

	log.Println("q", q)

	log.Println("RSA Encryption parameters. Pubkey [e, N]")

	cipherOne := new(big.Int).Exp(id1, e, n)

	val := new(big.Int).Exp(cipherOne, xb, n)
	modInvxaphi := new(big.Int).ModInverse(xa, phi)
	log.Println(val)
	if modInvxaphi == nil {
		log.Fatalf("failed to compute modInvxaphi")
	}
	log.Println(modInvxaphi)
	log.Println(n)

	val = new(big.Int).Exp(val, modInvxaphi, n)

	val = new(big.Int).Exp(val, d, n)

	log.Println("Derived ID2:", val)

}
