package Block

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	PreHash   string
	HashCode  string
	TimeStamp string
	Diff      int
	Noce      int
	Index     int
	Data      string
}

func GenerateFirstBlock(data string) Block {
	var block Block
	block.Data = data
	block.PreHash = "0"
	block.TimeStamp = time.Now().String()
	block.Diff = 4
	block.Index = 1
	block.Noce = 0
	block.HashCode = GenerationHashValue(block)
	return block
}

func GenerateNextBlock(data string, oldBlock Block) Block {
	var block Block
	block.Data = data
	block.PreHash = oldBlock.HashCode
	block.TimeStamp = time.Now().String()
	block.Diff = 4
	block.Index = 1
	block.Noce = 0
	//block.HashCode = GenerationHashValue(block)
	//需要挖矿 得到 满足条件的 hash
	block.HashCode = pow(block.Diff, &block)
	return block
}

func pow(diff int, block *Block) string {
	for {

		hash := GenerationHashValue(*block)
		fmt.Println("hashIng======", hash)
		if strings.HasPrefix(hash, strings.Repeat("0", diff)) {
			return hash
		} else {
			block.Noce++
		}
	}
}

func GenerationHashValue(block Block) string {
	hashdata := strconv.Itoa(block.Noce) + strconv.Itoa(block.Diff) + strconv.Itoa(block.Index) + block.TimeStamp + block.Data
	sha := sha256.New()
	sha.Write([]byte(hashdata))
	hashed := sha.Sum(nil)
	return hex.EncodeToString(hashed)
}
