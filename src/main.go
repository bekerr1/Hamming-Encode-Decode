package main

import (
	"fmt"
	"os"
	"hamming"
	"strings"
)


func main() {
	fmt.Printf("Hello, Person!!\n\n")

	//hamming.DecodeUsingStreamString("011100101110")

	if len(os.Args) != 3 {
		argError()
		os.Exit(1)
	}

	hammingMethod := os.Args[1]
	hammingMethod = strings.Title(hammingMethod)
	stringToUse := os.Args[2]

	switch hammingMethod {
	case "Decode":
		//hamming decode
		decoded := hamming.DecodeUsingStreamString(stringToUse)
		fmt.Printf("You gave me the stream %s \nI hamming decoded it to %s.\n", stringToUse, decoded)


	case "Encode":
		//hamming encode
		encoded := hamming.EncodeUsingStreamString(stringToUse)
		fmt.Printf("You gave me the stream %s \nI hamming encoded it to %s.\n", stringToUse, encoded)

	default:
		//error in argument
		argError()
	}

}

func argError() {
	fmt.Printf("Args must be 'Encode' or 'Decode' \nand should include a bitstring " +
		"in \nBigEndian (Network Byte) format.\n")
}











////Utility method to create the int64 in the correct 1/0 patterns - currently not used
//func createPositionBuffers(bitCount int) []*Parity {
//	var positions []*Parity
//
//	var exp int = 1
//	//for every position - 1, 2, 4, 8, 16, 32, 64
//	for position := 0; exp <= bitCount; exp = int(math.Exp2(float64(position))) {
//
//		var par = &Parity{
//			checkStream: 0,
//			checkBit: -1,
//			position: exp,
//		}
//		//for each position x - push x 1 bits and ignore x 1 bits
//		//position 1: 00001 -> 00101 -> 10101
//		//position 2: 00001 -> 000110 -> 1100110011
//		for {
//			par.addon(exp)
//			var bstring = bitString(par.checkStream)
//
//			//when the bit string is of the length of the bit stream - the amount of bits
//			//that get added and pushed each time - (since leading 0's arent kept)
//			//8 bits 110011 - 0's assumed when bit operations used later
//			if stringlen := len(bstring); stringlen >= bitCount - exp {
//				break
//			}
//			par.push(uint(exp))
//		}
//		position++
//		positions = append(positions, par)
//	}
//	return positions
//}
