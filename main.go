package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"

	"github.com/teamprex/proof-of-reserves/types"
)

func verifyMerkleProof(leafIndex int, leafHash []byte, root []byte, proofs [][]byte) bool {
	hash := leafHash
	// Process each level of the proof.
	for _, proof := range proofs {
		// Prepare a buffer for concatenation.
		var combined []byte
		if leafIndex%2 == 0 {
			// Current node is a left child; sibling goes to the right.
			combined = append(combined, hash...)
			combined = append(combined, proof...)
		} else {
			// Current node is a right child; sibling goes to the left.
			combined = append(combined, proof...)
			combined = append(combined, hash...)
		}
		// Compute the hash of the concatenated pair.
		h := sha256.Sum256(combined)
		hash = h[:]
		// Move to the parent node.
		leafIndex /= 2
	}
	fmt.Printf("Computed Hash: %x\n", hash)
	// If the computed hash equals the given root, the proof is valid.
	return bytes.Equal(hash, root)
}

func main() {
	content, err := os.ReadFile("data/merkle_proof.json")
	if err != nil {
		panic(err.Error())
	}
	merkleProof := &types.MerkleProof{}
	if err := json.Unmarshal(content, merkleProof); err != nil {
		panic(err.Error())
	}
	proofsBytes := merkleProof.GetProofsBytes()
	leafHashBytes := merkleProof.GetLeafHash()
	rootBytes := merkleProof.GetRootBytes()

	fmt.Printf("Computed Leaf Hash: %x\n", leafHashBytes)
	fmt.Printf("Expected Merkle Root: %x\n", rootBytes)

	// Verify the Merkle proof.
	if verifyMerkleProof(merkleProof.LeafIndex, leafHashBytes, rootBytes, proofsBytes) {
		fmt.Println("Merkle proof is valid!")
	} else {
		fmt.Println("Merkle proof is invalid!")
	}
}
