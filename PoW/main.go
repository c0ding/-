package main

import (
	"ConsensusAlgorithm/PoW/Block"
	"ConsensusAlgorithm/PoW/Blockchain"
)

func main() {
	first := Block.GenerateFirstBlock("创世")
	second := Block.GenerateNextBlock("第二", first)
	headerNode := Blockchain.GenerateHeaderNode(&first)
	Blockchain.AddNode(&second, headerNode)

	Blockchain.ShowNodes(headerNode)
}
