package types

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
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
	assetString := ""
	for _, asset := range p.Assets {
		assetString += "|" + asset.Currency + ":" + asset.Balance
	}
	leafHash := sha256.Sum256([]byte(assetString))
	return leafHash[:]
}

func (p *MerkleProof) GetRootBytes() []byte {
	rootBytes, err := hex.DecodeString(p.Root)
	if err != nil {
		panic(err.Error())
	}
	return rootBytes
}
