package main

import (
	"fmt"
	//"strconv"

	"main.go/blockchain"
)


func main() {
	

	 my_blockchain := blockchain.AddGenesis() // first step, this add the genesis blocks and also grants access to the blockchain list
	 my_blockchain.AddBlock("First block after Genesis")
	my_blockchain.AddBlock("Second block after Genesis")
	my_blockchain.AddBlock("Third block after Genesis")
	my_blockchain.AddBlock("Fourth block after Genesis")
	my_blockchain.AddBlock("Fifth block after Genesis")

	// Next, we have to iterate through the blockchain array(list) and get each block details
	for _, blocks := range my_blockchain.Blocks{
		//pow:= blockchain.NewProof(blocks)
	
		

		fmt.Printf("Previous Hash: %x\n", blocks.PrevHash)
		fmt.Printf("Data in block: %s\n", blocks.Data)
		fmt.Printf("Hash of block: %x\n", blocks.Hash)
		fmt.Println("Yipee!, a new block has been born💃")
				
	
	// 	//fmt.Printf("Proof: %s\n", strconv.FormatBool(pow.ValidateBlock()) )
		
	// 	fmt.Println()

}
}