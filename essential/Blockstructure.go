package essential

type Blockchain struct {
	block []*Block
}

func (chain *Blockchain) AddBlock(data string) {
	PrevBlock := chain.block[len(chain.block)-1]
	prevHash := PrevBlock.BlockHash
	block := NewBlock(data, prevHash)
	chain.block = append(chain.block, block)

}
func NewBlockchain(genesis string) *Blockchain {
	genesisblock := NewBlock(genesis, []byte{})
	return &Blockchain{[]*Block{genesisblock}}
}
