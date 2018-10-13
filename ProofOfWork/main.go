// This function implements a simplied block chain with Proof of Work
// level of difficulty: 24 bits
// This code is based on works from Jeiwan Github repo

package main

func main() {
	bc := NewBlockChain()
	bc.AddBlock("sent to J 100 million dollars")
	bc.AddBlock("charged J 110 million dollars")
	bc.Print()
	bc.Verify()
}

/*

$ go build
$ ./ProofOfWork
Block 0
Data: Genesis Block
Previous Hash:
Hash: 000000F6542C5410D5BB6CD1018CCB878F51276F01460AE4517A288ADED4FAF7 Nonce:25714619
Block 1
Data: sent to J 100 million dollars
Previous Hash: 000000F6542C5410D5BB6CD1018CCB878F51276F01460AE4517A288ADED4FAF7
Hash: 0000004E9625C63F2F01A24017E7375322761D5B7E9F474F91FB7AA82E28E39C Nonce:231331
Block 2
Data: charged J 110 million dollars
Previous Hash: 0000004E9625C63F2F01A24017E7375322761D5B7E9F474F91FB7AA82E28E39C
Hash: 000000E331695E595FA251905D4DBF327F59D759CE5EAC0982A53628B5C7EF3B Nonce:18632408
Block 0 verified: true
Block 1 verified: true
Block 2 verified: true

real  3m8.381s
user  3m11.780s
sys   0m0.920s

*/
