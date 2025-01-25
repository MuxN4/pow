package main

import (
	"fmt"
	"log"
	"time"

	"github.com/MuxN4/pow/block"
	"github.com/MuxN4/pow/pow"

	"github.com/fatih/color"
)

type Blockchain struct {
	Blocks []*block.Block
}

// Creates a blockchain with a genesis block
func NewBlockchain() *Blockchain {
	color.Cyan("🚀 Starting Mining Process...")

	genesisBlock := block.NewBlock(0, "Genesis Block", "", 2)
	powInstance := pow.NewProofOfWork(genesisBlock)
	success, _, _ := powInstance.Mine()

	if !success {
		log.Fatal("Failed to mine genesis block")
	}

	return &Blockchain{Blocks: []*block.Block{genesisBlock}}
}

// converts duration to human-readable format
func formatDuration(d time.Duration) string {
	switch {
	case d < time.Millisecond:
		return fmt.Sprintf("%dµs", d.Microseconds())
	case d < time.Second:
		return fmt.Sprintf("%dms", d.Milliseconds())
	default:
		return fmt.Sprintf("%.2fs", d.Seconds())
	}
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
		color.Red("❌ Failed to mine block")
		return nil
	}

	bc.Blocks = append(bc.Blocks, newBlock)

	color.Yellow("⛏️  Block mined!")
	fmt.Printf("    %s:   %d\n", color.CyanString("Difficulty"), difficulty)
	fmt.Printf("    %s:        %d\n", color.CyanString("Nonce"), nonce)
	fmt.Printf("    %s:         %s\n", color.GreenString("Time"), formatDuration(duration))
	fmt.Printf("    %s:         %s\n", color.BlueString("Hash"), newBlock.Hash)

	// Handle previous hash display
	prevHashDisplay := previousHash
	if prevHashDisplay == "" {
		prevHashDisplay = "[Genesis Block]"
	}
	fmt.Printf("    %s:    %s\n\n", color.MagentaString("Prev Hash"), prevHashDisplay)

	return newBlock
}

func main() {
	blockchain := NewBlockchain()

	// A simulation with different difficulties
	difficulties := []int{2, 4, 6}
	for _, diff := range difficulties {
		blockchain.AddBlock(fmt.Sprintf("Transaction at difficulty %d", diff), diff)
	}

	color.Green("🎉 Mining Complete!")
}
