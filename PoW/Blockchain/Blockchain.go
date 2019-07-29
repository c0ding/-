/*
使用链表的结构 实现区块链
*/
package Blockchain

import (
	"ConsensusAlgorithm/PoW/Block"
	"fmt"
)

type Node struct {
	NextNode *Node
	Data     *Block.Block
}

func GenerateHeaderNode(data *Block.Block) *Node {
	headerNode := new(Node)
	headerNode.NextNode = nil
	headerNode.Data = data
	return headerNode
}

func AddNode(data *Block.Block, preNode *Node) *Node {
	newNode := new(Node)
	newNode.Data = data
	newNode.NextNode = nil
	preNode.NextNode = newNode
	return newNode
}

func ShowNodes(node *Node) {
	n := node
	for {
		if n.NextNode == nil {
			fmt.Println(n.Data)
			break
		} else {
			fmt.Println(n.Data)
			n = n.NextNode
		}
	}
}
