package main

import (
	"fmt"
	"math"
	"hamming"
)


func main() {
	fmt.Println("hello, world!")

	var instream = &IncomingBitStream{
		encodedStream: 0,
		streamString: "001100010100",
		errorBitPosition: -1,
	}

	//start := time.Now()
	//var outstream = &OutgoingBitStream{
	//	positions: 0,
	//	stream: 0,
	//	encodedStream: 0,
	//	streamString: "",
	//	encodedString: "",
	//}
	//outstream.hammingEncodeString()
	instream.hammingDecodeString()
	//elapsed := time.Since(start)
	//fmt.Printf("Input: %s ---- Output: %s ---- Took %s to execute",
	//	stream.streamString, stream.encodedString, elapsed)
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
