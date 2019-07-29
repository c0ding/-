package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Block struct {
	PreHash   string
	HashCode  string
	TimeStamp string
	Data      string //交易信息
	Index     int
	Validator *PNode // 区块验证者
}

type PNode struct {
	Tokens  int
	Days    int
	Address string
}

func GenerateFirstBlock(data string) Block {
	var block Block
	block.PreHash = "0"
	block.TimeStamp = time.Now().String()
	block.Index = 1
	block.HashCode = GenerationHashValue(block)
	block.Data = data
	block.Validator = &PNode{0, 0, ""}

	return block
}

// 模拟链上 有5个节点
var nodes = make([]PNode, 5)

// 有15个地址 竞争挖矿权
var adds = make([]*PNode, 15)

func initNodes() {
	nodes[0] = PNode{5, 1, "0x12315"}
	nodes[1] = PNode{1, 1, "0x12314"}
	nodes[2] = PNode{2, 1, "0x12313"}
	nodes[3] = PNode{3, 1, "0x12312"}
	nodes[4] = PNode{4, 1, "0x12311"}

	counts := 0
	for i := 0; i < len(nodes); i++ {
		for j := 0; j < nodes[i].Tokens*nodes[i].Days; j++ {
			adds[counts] = &nodes[i]
			counts++
		}

	}

	//fmt.Print("节点【tokens，Days，Address】:\n")
	//fmt.Println("%v \n", nodes)
	for i, node := range nodes {
		fmt.Println(i, node)
	}
}

func newBlock(lastBlock *Block, data string) Block {
	var block Block
	block.Data = data
	block.PreHash = lastBlock.HashCode
	block.TimeStamp = time.Now().String()
	block.Index = lastBlock.Index + 1

	// 采用PoS找到挖矿的节点
	time.Sleep(1000000000)
	rand.Seed(time.Now().Unix())
	var rd = rand.Intn(15) //生成一个随机数，【0-15） 之间
	node := adds[rd]       //根据随机数找到矿工
	fmt.Println("查看矿工的地址", node.Address)
	block.Validator = node
	node.Tokens += 1
	block.HashCode = GenerationHashValue(block)

	return block
}

func GenerationHashValue(block Block) string {
	hashdata := block.Data + strconv.Itoa(block.Index) + block.TimeStamp
	sha := sha256.New()
	sha.Write([]byte(hashdata))
	hashed := sha.Sum(nil)
	return hex.EncodeToString(hashed)
}

func main() {
	initNodes()
	var firstBlock = GenerateFirstBlock("创世区块")

	for i := 0; i < 30; i++ {
		block := newBlock(&firstBlock, "新的区块")
		fmt.Println("新的区块信息", block)
	}

	//PoS的特点： 拥有tokens数量越多，获得记账权利的概率越大
}
