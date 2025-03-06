package types

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
)

// Asset represents each asset in the assets list.
type Asset struct {
	Currency string `json:"currency"`
	Balance  string `json:"balance"`
}

// MerkleProof represents the structure of the JSON file.
type MerkleProof struct {
	LeafIndex int      `json:"leaf_index"`
	Assets    []Asset  `json:"assets"`
	Root      string   `json:"root"`
	Proofs    []string `json:"proofs"`
}

func (p *MerkleProof) GetProofsBytes() [][]byte {
	proofsBytes := make([][]byte, len(p.Proofs))
	for i, proofStr := range p.Proofs {
		decoded, err := base64.StdEncoding.DecodeString(proofStr)
		if err != nil {
			panic(err.Error())
		}
		proofsBytes[i] = decoded
	}
	return proofsBytes
}

func (p *MerkleProof) GetLeafHash() []byte {
	assetsJSON, err := json.Marshal(p.Assets)
	if err != nil {
		panic(err.Error())
	}
	leafHash := sha256.Sum256(assetsJSON)
	return leafHash[:]
}

func (p *MerkleProof) GetRootBytes() []byte {
	rootBytes, err := hex.DecodeString(p.Root)
	if err != nil {
		panic(err.Error())
	}
	return rootBytes
}
