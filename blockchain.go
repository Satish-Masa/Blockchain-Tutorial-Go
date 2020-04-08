package main

import "github.com/boltdb/bolt"

const dbFile = "blockchain.db"
const blocksBucket = "blocks"

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

type BlockchainIterator struct {
	currentHash []byte
	db          *bolt.DB
}

// AddBlock saves provided data as a block in the blockchain
func (bc *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("|"))

		return nil
	})

	newBlock := NewBlock(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		err = b.Put([]byte("|"), newBlock.Hash)
		bc.tip = newBlock.Hash

		return nil
	})
}

// NewBlockchain creates a new Blockchain with genesis Block
func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("|"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("|"))
		}

		return nil
	})

	bc := Blockchain{tip, db}

	return &bc
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{bc.tip, bc.db}
}

func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	i.currentHash = block.PrevBlockHash

	return nil
}
