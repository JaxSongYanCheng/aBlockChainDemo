package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64  //
	Data          []byte // current block info
	PrevBlockHash []byte // prev block hash
	Hash          []byte // current block hash
}

/**
Block bind a function
*/

func (this *Block) SetHash() {
	// timestamp+data+prevBlock hash
	timestamp := []byte(strconv.FormatInt(this.Timestamp, 10))
	headers := bytes.Join([][]byte{this.PrevBlockHash, this.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	this.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) (block *Block) {
	block = &Block{Timestamp: time.Now().Unix(), Data: []byte(data), PrevBlockHash: prevBlockHash}
	block.SetHash()
	return
}

type BlockChain struct {
	Blocks []*Block
}

func (this *BlockChain) AddBlock(data string) {
	prevBlock := this.Blocks[len(this.Blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	this.Blocks = append(this.Blocks, newBlock)
}

func NewGenesisBlock() (genesisBlock *Block) {
	genesisBlock = NewBlock("Genesis Block", nil)
	return
}

func NewBlockChain() (blockChain *BlockChain) {
	blockChain = &BlockChain{}
	blockChain.Blocks = []*Block{NewGenesisBlock()}
	return
}

func main() {
	bc := NewBlockChain()
	var cmd string
	for {
		fmt.Println("input 1 add a information to block chain")
		fmt.Println("input 2 traverse the block chain")
		fmt.Println("input other exit")
		fmt.Scanf("%s", &cmd)
		switch cmd {
		case "1":
			fmt.Println("input information:")
			input := make([]byte, 1024)
			n, _ := os.Stdin.Read(input)
			bc.AddBlock(string(input[:n]))
		case "2":
			for i, block := range bc.Blocks {
				fmt.Printf("%d is %s\n", i, string(block.Data))
			}
		default:
			return
		}
	}
}
