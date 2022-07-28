package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)


type Block struct {
	Index int
	Hash []byte
	Data []byte
	PrevHash []byte
	Nonce int

}
type BlockChain struct{
	Blocks []*Block
	
}
type Proof struct{
	BlockData *Block
	Target *big.Int
}

var Difficulty = 10



func (b *Block) CalHash(){
	data := bytes.Join([][]byte{
		IntToBytes(int64(b.Index)),
		b.Data,
		b.PrevHash,
		IntToBytes(int64(b.Nonce)),
	}, []byte{})
	hash := sha256.Sum256(data)
	b.Hash = hash[:]
	//b.Nonce = nonce
	
}

func IntToBytes(num int64) []byte{
	buffer := new(bytes.Buffer)
	data := binary.Write(buffer, binary.BigEndian, num)
	if data != nil{
		fmt.Println(data)
	}
	return buffer.Bytes()
}
func CreateBlock(data string, prevHash []byte)*Block{
	block := &Block{0,[]byte{},[]byte(data), prevHash,0}
	new := DefTarget(block)
	nonce := new.Proof()
	block.Nonce = nonce
	block.CalHash()
	
	return block
}
//func (block *BlockChain) GetBlock() {
//	for _, blocks := range block.Blocks{
		//fmt.Println(blocks.Data)
//	}

//}
func DefTarget(block *Block) *Proof{
	target := big.NewInt(0)
	calTarget := target.Lsh(target, uint(256-Difficulty))
	return &Proof{block, calTarget}

}
// We want users to be able to add blocks on the blockchain
// How: If the hash of the block less than target, the block is valid
func (b *BlockChain) AddBlock(data string){
	prevBlock := b.Blocks[len(b.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	b.Blocks = append(b.Blocks, newBlock)
}

func (proof *Proof) Proof() int{
	var HashInt big.Int
	nonce := proof.BlockData.Nonce
	HashInt.SetBytes(proof.BlockData.Hash)
	

	for nonce < math.MaxInt64{	
		if HashInt.Cmp(proof.Target) == -1{
			break
		}else{
			nonce++
		}
		fmt.Printf("nonce: %x/r", nonce)
	}
	
	return nonce
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}
func AddGenesis() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
