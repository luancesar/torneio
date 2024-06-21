package blockchain

import "time"

type Blockchain struct {
    Blocks []Block
}

func (bc *Blockchain) AddBlock(data string) {
    prevBlock := bc.Blocks[len(bc.Blocks)-1]
    newBlock := NewBlock(prevBlock, data)
    bc.Blocks = append(bc.Blocks, newBlock)
}

func GenesisBlock() Block {
    return Block{
        Index:     0,
        Timestamp: time.Now().String(),
        Data:      "Genesis Block",
        PrevHash:  "",
        Hash:      "",
    }
}

func NewBlockchain() *Blockchain {
    genesisBlock := GenesisBlock()
    genesisBlock.Hash = genesisBlock.CalculateHash()
    return &Blockchain{[]Block{genesisBlock}}
}
