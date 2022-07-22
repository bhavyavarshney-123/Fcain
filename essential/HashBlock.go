package essential

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strings"
)

type Block struct {
	PrioriHash []byte
	BlockData  []byte
	BlockHash  []byte
}

func (b *Block) GenerateHash() []byte {
	input := bytes.Join([][]byte{b.PrioriHash, b.BlockData}, []byte{})
	hash := sha256.Sum224(input)
	return hash[:]
}
func NewBlock(data string, Priori []byte) *Block {
	block := &Block{
		PrioriHash: Priori,
		BlockData:  []byte(data),
	}

	block.BlockHash = block.GenerateHash()

	return block

}
func (b *Block) String() string {
	var s strings.Builder
	s.WriteString("Block Hash\n")
	s.WriteString(fmt.Sprintf("%x", b.BlockHash))
	s.WriteString("\n")

	s.WriteString("Block data\n")
	s.WriteString(string(b.BlockData))
	s.WriteString("\n")

	s.WriteString("Block Priori\n")
	s.WriteString(fmt.Sprintf("%x", b.PrioriHash))
	s.WriteString("\n")

	return s.String()
}
