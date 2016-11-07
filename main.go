package main

import (
	"fmt"
	"strconv"
	"math"
	"time"
)

type BitStream struct {
	positions int
	stream uint64
	encodedStream uint64
	streamString string
	encodedString string
}

type Parity struct {
	checkStream uint64
	checkBit int
	position int
}

func main() {
	fmt.Println("hello, world!")
	start := time.Now()
	var stream = &BitStream{
		positions: 0,
		stream: 0,
		encodedStream: 0,
		streamString: "1010101",
		encodedString: "",
	}
	stream.hammingEncodeString()
	elapsed := time.Since(start)
	fmt.Printf("Input: %s ---- Output: %s ---- Took %s to execute",
		stream.streamString, stream.encodedString, elapsed)
}

/*
Method: 1) Mark all bit positions that are powers of 2 as parity bits (ex. 1, 2, 4, 8, 16, 32, 64)
	2) All other bit positions belong to the actual data
	3) To decide the value of the parity bit - for each position n, check n bits and skip n bits
		ex. position 1 - check 1 bit skit 1 bit
		position 8 - check 8 bits skip 8 bits
	4) XOR the total bit values
 */


func (b *BitStream) hammingEncodeString() {
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



func (b *BitStream) hammingEncode() {
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


//position 1 - 6148914691236517205 == 1010101 cont'd
//position 2 - 3689348814741910323 == 1100110011 cont'd
//position 4 - 1085102592571150095 == 111100011110001111 cont'd
//position 8 - 71777214294589695 == 11111110000000011111111 cont'd
//position 16 - 281470681808895 == 00000000000000001111111111111111 cont'd
//position 32 - 4294967295 == 1111111111111111111111111111111 cont'd
//position 64 - Int64.Max

func int64ParityPositions() [7]*Parity {
	return [7]*Parity {
		&Parity{
			checkStream: 6148914691236517205,
			checkBit: -1,
			position: 1,
		},
		&Parity{
			checkStream: 3689348814741910323,
			checkBit: -1,
			position: 2,
		},
		&Parity{
			checkStream: 1085102592571150095,
			checkBit: -1,
			position: 4,
		},
		&Parity{
			checkStream: 71777214294589695,
			checkBit: -1,
			position: 8,
		},
		&Parity{
			checkStream: 281470681808895,
			checkBit: -1,
			position: 16,
		},
		&Parity{
			checkStream: 4294967295,
			checkBit: -1,
			position: 32,
		},
		&Parity{
			checkStream: 0,
			checkBit: -1,
			position: 64,
		},
	}
}

//adds 0 padding to parity positions
func (b *BitStream) padParityPositions(power int) string {

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

func createPositionBuffers(bitCount int) []*Parity {
	var positions []*Parity

	var exp int = 1
	//for every position - 1, 2, 4, 8, 16, 32, 64
	for position := 0; exp <= bitCount; exp = int(math.Exp2(float64(position))) {

		var par = &Parity{
			checkStream: 0,
			checkBit: -1,
			position: exp,
		}
		//for each position x - push x 1 bits and ignore x 1 bits
		//position 1: 00001 -> 00101 -> 10101
		//position 2: 00001 -> 000110 -> 1100110011
		for {
			par.addon(exp)
			var bstring = bitString(par.checkStream)

			//when the bit string is of the length of the bit stream - the amount of bits
			//that get added and pushed each time - (since leading 0's arent kept)
			//8 bits 110011 - 0's assumed when bit operations used later
			if stringlen := len(bstring); stringlen >= bitCount - exp {
				break
			}
			par.push(uint(exp))
		}
		position++
		positions = append(positions, par)
	}
	return positions
}

//add on 1 bits and push bits over by an amount
func (p *Parity) addon(add int)  {
	for add > 0 {
		p.checkStream = p.checkStream << 1
		p.checkStream ++
		//p.checkStream << 1
		add --
	}
	bitString(p.checkStream)
}

func (p *Parity) push(by uint)  {
	p.checkStream = p.checkStream << by
}


//Utility function to count number of '1' bits in an integer
func bitCount(bits uint64) int {

	var count int = 0
	for bits != 0 {
		bitString(bits)
		bits = bits & (bits - 1)
		count ++
	}
	return count
}

func bitString(stream uint64) string {
	return strconv.FormatInt(int64(stream), 2)
}

func reverseRune(rarr []rune) []rune {
	var n = len(rarr)
	for i := 0; i < n/2; i++ {
		rarr[i], rarr[n-1-i] = rarr[n-1-i], rarr[i]
	}
	return rarr
}


