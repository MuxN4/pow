package main

import (
	"fmt"
	"log"

	"github.com/MuxN4/pow/block"
	"github.com/MuxN4/pow/pow"
)

type Blockchain struct {
	Blocks []*block.Block
}

// Creates a blockchain with a genesis block
func NewBlockchain() *Blockchain {
	genesisBlock := block.NewBlock(0, "Genesis Block", "", 2)
	powInstance := pow.NewProofOfWork(genesisBlock)
	success, _, _ := powInstance.Mine()

	if !success {
		log.Fatal("Failed to mine genesis block")
	}

	return &Blockchain{Blocks: []*block.Block{genesisBlock}}
}

// Where mining and adding a new block to the blockchain takes place
func (bc *Blockchain) AddBlock(data string, difficulty int) *block.Block {
	previousHash := ""
	if len(bc.Blocks) > 0 {
		previousHash = bc.Blocks[len(bc.Blocks)-1].Hash
	}

	newBlock := block.NewBlock(len(bc.Blocks), data, previousHash, difficulty)
	powInstance := pow.NewProofOfWork(newBlock)

	success, nonce, duration := powInstance.Mine()
	if !success {
		return nil
	}

	bc.Blocks = append(bc.Blocks, newBlock)

	fmt.Printf("Block mined: Difficulty=%d, Nonce=%d, Time=%v\n",
		difficulty, nonce, duration)

	return newBlock
}

func main() {
	blockchain := NewBlockchain()

	// A simulation with different difficulties
	difficulties := []int{2, 4, 6}
	for _, diff := range difficulties {
		blockchain.AddBlock(fmt.Sprintf("Transaction at difficulty %d", diff), diff)
	}

	for _, block := range blockchain.Blocks {
		fmt.Printf("Block #%d: Hash=%s, Prev=%s, Difficulty=%d\n",
			block.Index, block.Hash, block.PreviousHash, block.Difficulty)
	}
}
