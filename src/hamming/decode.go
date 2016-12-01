package hamming
import "strconv"
import "fmt"

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
	errorParityPositions []uint64
	errorBitPosition uint64
	streamString string
	decodedString string
}

func DecodeUsingStreamString(stream string) string {

	decodeStream := &IncomingBitStream{
		encodedStream: 0,
		errorParityPositions: make([]uint64, 0),
		errorBitPosition: 0,
		streamString: stream,
		decodedString: "",
	}

	decodeStream.hammingDecodeString()
	return decodeStream.decodedString
}

func (b *IncomingBitStream) hammingDecodeString() {
	//No way to reverse a string in Go, but if you convert string to []rune (go 'char' type)
	//can eaisily reverse those. (to get bits into LittleEndian order or bit operations
	reversedString := string(reverseRune([]rune(b.streamString)))
	bit_64, err := strconv.ParseInt(reversedString, 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	b.encodedStream = uint64(bit_64)
	b.hammingDecode()
}

func (b *IncomingBitStream) hammingDecode() {

	parityPositions := int64ParityPositions()
	var errorPosition int
	for i := 0; i < len(parityPositions); i++ {
		parity := parityPositions[i]
		adjustedCheckStream := parity.checkStream << uint(parity.position - 1)

		checkResult := b.encodedStream & adjustedCheckStream
		evenOdd := bitCount(checkResult)
		if evenOdd % 2 == 0 {
			//even number of 1's => 0 parity bit - do nothing
			parityPositions[i].checkBit = 0
		} else {
			//odd number of 1's => 1 pairty bit - push by i + 1 and add another 1
			errorPosition += parityPositions[i].position
		}
	}
	//fmt.Printf("Error at bit position: %i\n", errorPosition)
	b.fixErrorBit(errorPosition - 1)
}

//Check if the UTF-8 value at the error position in the string given is 0 or 1 and
//flip it
func (b *IncomingBitStream) fixErrorBit(errorPosition int) {

	streamRune := []rune(b.streamString)
	if streamRune[errorPosition] == 48 {
		streamRune[errorPosition] = 49
	} else {
		streamRune[errorPosition] = 48
	}
	b.decodedString = string(streamRune)
	//fmt.Printf("Corrected stream string is: %s", b.decodedString)
	b.stripParityPositions()

}

func (b *IncomingBitStream) stripParityPositions() {
	stripped := make([]rune, 0)
	streamRune := []rune(b.decodedString)
	var lastcheck int
	if len(streamRune) > 3 {
		lastcheck = 4
		stripped = append(stripped, streamRune[2:3]...)
	}
	if len(streamRune) > 7 {
		lastcheck = 8
		stripped = append(stripped, streamRune[4:7]...)
	}
	if len(streamRune) > 15 {
		lastcheck = 16
		stripped = append(stripped, streamRune[8:15]...)
	}
	if len(streamRune) > 31 {
		lastcheck = 32
		stripped = append(stripped, streamRune[16:31]...)
	}
	if len(streamRune) > 63 {
		stripped = append(stripped, streamRune[32:63]...)
		b.decodedString = string(stripped)
		return
	}
	stripped = append(stripped, streamRune[lastcheck:]...)
	b.decodedString = string(stripped)
}


