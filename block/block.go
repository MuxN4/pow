package block

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// Represents a single block in the blockchain
type Block struct {
	Index        int
	Timestamp    int64
	Data         string
	PreviousHash string
	Hash         string
	Nonce        int
	Difficulty   int
}

// NewBlock creates a new block with given parameters
func NewBlock(index int, data string, previousHash string, difficulty int) *Block {
	return &Block{
		Index:        index,
		Timestamp:    time.Now().Unix(),
		Data:         data,
		PreviousHash: previousHash,
		Nonce:        0,
		Difficulty:   difficulty,
	}
}

// Here CalculateHash generates a unique hash for the block
func (b *Block) CalculateHash() string {
	record := strconv.Itoa(b.Index) +
		strconv.FormatInt(b.Timestamp, 10) +
		b.Data +
		b.PreviousHash +
		strconv.Itoa(b.Nonce)

	h := sha256.New()
	h.Write([]byte(record))
	return hex.EncodeToString(h.Sum(nil))
}
