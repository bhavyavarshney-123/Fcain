package essential

type Blockchain struct {
	Blocks []*Block
}

func (Chain *Blockchain) AddBlock(data string) {
	PrevBlock := Chain.Blocks[len(Chain.Blocks)-1]
	prevHash := PrevBlock.BlockHash
	block := NewBlock(data, prevHash)
	Chain.Blocks = append(Chain.Blocks, block)

}
func NewBlockchain(genesis string) *Blockchain {
	genesisblock := NewBlock(genesis, []byte{})
	return &Blockchain{[]*Block{genesisblock}}
}
