package main

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

const targetBits = 16 

type ProofOfWork struct {
	*Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	return &ProofOfWork{b, target}
}

func (pow *ProofOfWork) data(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Prehash,
			pow.Data,
			IntToHex(pow.Timestamp),
			IntToHex(uint64(targetBits)),
			IntToHex(uint64(nonce)),
		},
		[]byte{},
	)
	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	var nonce = 0
	for nonce < maxNonce {
		hash = sha256.Sum256(pow.data(nonce))
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	return nonce, hash[:]
}

func (pow *ProofOfWork) Verify() bool {
	var hashInt big.Int
	hash := sha256.Sum256(pow.data(pow.Block.Nonce))
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(pow.target) == -1
}
