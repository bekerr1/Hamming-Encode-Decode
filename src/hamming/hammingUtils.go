package hamming

import "strconv"


type Parity struct {
	checkStream uint64
	checkBit int
	position int
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

func (p *Parity) addOne() {
	p.checkStream += 1
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
