package pow

import (
	"strings"
	"time"

	"github.com/MuxN4/pow/block"
)

type ProofOfWork struct {
	block    *block.Block
	target   string
	maxNonce int
}

// This function initializes a new Proof of Work instance with some ground rules
func NewProofOfWork(block *block.Block) *ProofOfWork {
	target := strings.Repeat("0", block.Difficulty)
	maxNonce := calculateMaxNonce(block.Difficulty)

	return &ProofOfWork{
		block:    block,
		target:   target,
		maxNonce: maxNonce,
	}
}

// determines how many attempts we'll tolerate
func calculateMaxNonce(difficulty int) int {
	return 1_000_000 * difficulty
}

// Mine attempts to find a valid hash meeting the difficulty criteria
func (pow *ProofOfWork) Mine() (bool, int, time.Duration) {
	startTime := time.Now()

	for pow.block.Nonce < pow.maxNonce {
		hash := pow.block.CalculateHash()

		if strings.HasPrefix(hash, pow.target) {
			pow.block.Hash = hash
			return true, pow.block.Nonce, time.Since(startTime)
		}

		pow.block.Nonce++
	}

	return false, -1, time.Since(startTime)
}

// checks if the block's hash meets difficulty requirements
func (pow *ProofOfWork) Validate() bool {
	return strings.HasPrefix(pow.block.Hash, pow.target)
}
