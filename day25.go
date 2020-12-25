package main

import (
	"fmt"
)

func main() {
	Day25()
}

func Day25() {
	fmt.Println("==============  Day25 =============")
	fmt.Println("==============  TEST  =============")
	cardPubKey := 5764801
	doorPubKey := 17807724
	encryptionKey := determineEncryptionKey(cardPubKey, doorPubKey)
	if encryptionKey == 14897079 {
		fmt.Println("Encryption Found:", encryptionKey)
	} else {
		fmt.Println("Different encryptions Found:", encryptionKey)
	}
	fmt.Println("============== OUTPUT =============")
	fmt.Println("Encryption Key:", determineEncryptionKey(2069194, 16426071))
}

func determineEncryptionKey(cardPubKey int, doorPubKey int) int {
	cardLoopSize, doorLoopSize := findLoopSize(cardPubKey), findLoopSize(doorPubKey)
	cardEncryptionKey := 1
	for loopCount := 0; loopCount < doorLoopSize; loopCount++ {
		cardEncryptionKey = loop(cardEncryptionKey, cardPubKey)
	}

	doorEncryptionKey := 1
	for loopCount := 0; loopCount < cardLoopSize; loopCount++ {
		doorEncryptionKey = loop(doorEncryptionKey, doorPubKey)
	}
	if cardEncryptionKey == doorEncryptionKey {
		return cardEncryptionKey
	} else {
		return -1
	}
}

func loop(input int, subjectNumber int) int {
	output := 0
	output = input * subjectNumber
	output = output % 20201227
	return output
}

func findLoopSize(pubKey int) int {
	pubKeyOutput := 1
	pubKeyLoop := 0
	for pubKeyOutput != pubKey {
		//fmt.Println(pubKeyOutput)
		pubKeyOutput = loop(pubKeyOutput, 7)
		pubKeyLoop++
	}
	return pubKeyLoop
}
