package hamming

import (
	"fmt"
	"strconv"
)

/*
Method: 1) Mark all bit positions that are powers of 2 as parity bits (ex. 1, 2, 4, 8, 16, 32, 64)
	2) All other bit positions belong to the actual data
	3) To decide the value of the parity bit - for each position n, check n bits and skip n bits
		ex. position 1 - check 1 bit skit 1 bit
		position 8 - check 8 bits skip 8 bits
	4) XOR the total bit values
 */

type IncomingBitStream struct {
	encodedStream uint64
	streamString string
	errorBitPosition int

}


func (b *IncomingBitStream) hammingDecodeString() {

	reversedString := string(reverseRune([]rune(b.streamString)))
	sixtyFour, err := strconv.ParseInt(reversedString, 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	b.encodedStream = uint64(sixtyFour)
	b.hammingDecode()
}

func (b *IncomingBitStream) hammingDecode() {

	var accumulation Parity

	parityPositions := int64ParityPositions()

	for i := 0; i < len(parityPositions); i++ {
		parity := parityPositions[i]
		adjustedCheckStream := parity.checkStream << uint(parity.position - 1)
		fmt.Println(bitString(b.encodedStream))
		fmt.Println("&'d with")
		fmt.Println(bitString(adjustedCheckStream))
		checkResult := b.encodedStream & adjustedCheckStream
		evenOdd := bitCount(checkResult)
		if evenOdd % 2 == 0 {
			//even number of 1's => 0 parity bit - do nothing
		} else {
			//odd number of 1's => 1 pairty bit - push by i + 1 and add another 1
			accumulation.push(uint(i))
			accumulation.addOne()
		}

		fmt.Println("Accumulation:")
		fmt.Println(bitString(accumulation.checkStream))

	}

}


func (b *IncomingBitStream) stripParityPositions() {

}


