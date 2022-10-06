package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

type block struct {
	transaction   string
	nonce_x       int
	previous_hash string
	current_hash  string
}

type blockchain struct {
	list []*block
}

func (temp *block) GetAllAttributesInString() string {

	tempstring := ""
	tempstring += strconv.Itoa(temp.nonce_x) // converting integer nonce to string
	tempstring += temp.transaction + temp.previous_hash

	return tempstring

}

func newblock(x int, transact string) *block {
	tempblock := new(block)
	tempblock.nonce_x = x
	tempblock.transaction = transact
	return tempblock
}

func verifypreviousblocks(chain *blockchain) bool {
	var temp = ""
	var check = true
	for i := 0; i < len(chain.list); i++ {
		total_sum := sha256.Sum256([]byte(chain.list[i].GetAllAttributesInString()))
		temp = fmt.Sprintf("%x", total_sum) // formating to string

		if temp != chain.list[i].current_hash {
			check = false
			fmt.Printf("Previous block has been tampered, i.e. Block # %d\n", i)
			break
		}
	}

	if check == false {
		fmt.Println("error occured")
	} else {
		fmt.Printf("Blocks verified. No tampering\n")
	}

	return check
}

func hashcalculate(chain *blockchain) {

	for i := 0; i < len(chain.list); i++ {
		total_sum := sha256.Sum256([]byte(chain.list[i].GetAllAttributesInString()))
		chain.list[i].current_hash = fmt.Sprintf("%x", total_sum) // formating to string
		if i < len(chain.list)-1 {
			chain.list[i+1].previous_hash = fmt.Sprintf("%x", total_sum) //storing current block hash to next block in its previous hash var
		}

	}
}

func (blocklist *blockchain) addblock(x int, transact string) *block {
	tempblock := newblock(x, transact)

	if verifypreviousblocks(blocklist) {
		blocklist.list = append(blocklist.list, tempblock)
		hashcalculate(blocklist)

		fmt.Printf("block addition in chain successful\n")
	} else {
		fmt.Printf(" error. block addition unsuccessful.\n")
		return nil
	}
	return tempblock
}

func displayblocks(blocklist *blockchain) {
	fmt.Println("")

	for i := 0; i < len(blocklist.list); i++ {
		fmt.Printf("Block id:%d\n\n", i)
		fmt.Println("transaction: \n", blocklist.list[i].transaction)
		fmt.Println("Nonce x value : \n", blocklist.list[i].nonce_x)
		fmt.Println("current hash: \n", blocklist.list[i].current_hash)
		fmt.Println("previous hash: \n", blocklist.list[i].previous_hash)

	}

	fmt.Println("")

}

func updateblock(chain *blockchain, x int, transact string) { // updating on basis of nonce value as identifier

	found := false
	for i := 0; i < len(chain.list); i++ {

		if x == chain.list[i].nonce_x {
			chain.list[i].transaction = transact
			fmt.Println("updated successfully\n")
			found = true
		}
	}
	if found == false {
		fmt.Println("error. Couldnt update. block not found")
	}
	return
}

func main() {

	fmt.Println("Welcome to Asif Mujeeb's CryptoCurrency \n")

	chain := new(blockchain)
	var x = 50
	chain.addblock(x, "transaction#1")
	chain.addblock(x+10, "transaction#2")
	chain.addblock(x+20, "transaction#3")
	chain.addblock(x+30, "transaction#4")
	chain.addblock(x+40, "transaction#5")
	chain.addblock(x+50, "transaction#6")
	displayblocks(chain)

	fmt.Println("updating block with nonce value :", x+30)
	updateblock(chain, x+30, "transaction updated")

	displayblocks(chain)

}
