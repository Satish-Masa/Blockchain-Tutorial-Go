package main

import "bytes"

type TXInput struct {
	Txid      []byte
	Vout      int
	Signature []byte
	PubKey    []byte
}

func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lookingHash := HashPubKey(in.PubKey)

	return bytes.Compare(lookingHash, pubKeyHash) == 0
}
