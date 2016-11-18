package hamming

import (
	"fmt"
	"strconv"
	"math"
)

type OutgoingBitStream struct {
	positions int
	stream uint64
	encodedStream uint64
	streamString string
	encodedString string
}



func (b *OutgoingBitStream) hammingEncodeString() {
	//Input bit streams will have 1 to 64 bits - Should take a stream and return a
	//hamming encoded stream


	//Pad the string with 0's for parity so they dont interfere with & operation but
	//so & is accurate at the same time
	paddedStream := b.padParityPositions(2)

	stream, err := strconv.ParseInt(paddedStream, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	b.stream = uint64(stream)
	b.hammingEncode()
}



func (b *OutgoingBitStream) hammingEncode() {
	//Input bit streams will have 1 to 64 bits - Should take a stream and return a
	//hamming encoded stream

	parityPositions := int64ParityPositions()
	hamEncoded := []rune(b.encodedString)

	//For each position perform a bitwise AND between the padded stream and that positions
	//checkstream.  the number of 1 bits in the result should be analyzed and if an odd number
	//of 1 bits, that positions perity value is 1, if even, that positions parity value is 0
	for i := 0; i < b.positions; i++ {
		parity := parityPositions[i]
		adjustedCheckStream := parity.checkStream << uint(parity.position - 1)
		checkResult := b.stream & adjustedCheckStream
		evenOdd := bitCount(checkResult)
		if evenOdd % 2 == 0 {
			//even number of 1's => 0 parity bit
			parity.checkBit = 0
			hamEncoded[parity.position - 1] = 48
		} else {
			//odd number of 1's => 1 pairty bit
			parity.checkBit = 1
			hamEncoded[parity.position - 1] = 49
		}
	}
	b.encodedString = string(hamEncoded)
}




//adds 0 padding to parity positions
func (b *OutgoingBitStream) padParityPositions(power int) string {

	var bitStringArr []rune
	var paddedArr []rune

	bitStringArr = []rune(b.streamString)

	for index, powIndex, bitIndex := 0, 0, 0; index < len(bitStringArr) + b.positions; index++ {

		expIndex := int(math.Exp2(float64(powIndex)))
		if index == expIndex - 1 {
			//48 == rune value for "0"
			paddedArr = append(paddedArr, 48)
			powIndex ++
			b.positions ++
		} else {
			paddedArr = append(paddedArr, rune(bitStringArr[bitIndex]))
			bitIndex ++
		}
	}
	b.encodedString = string(paddedArr)
	return string(reverseRune(paddedArr))
}
