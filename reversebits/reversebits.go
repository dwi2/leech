package reversebits

func ReverseBits(num uint32) uint32 {
	var maskL, maskR uint32 = 0x1 << 31, 0x1
	var index uint32 = 0
	result := uint32(0)

	for i := 31; i >= 15; i-- {
		tempL := (num & maskL) >> (uint32(i) - index)
		tempR := (num & maskR) << (uint32(i) - index)
		result = result | (tempL | tempR)
		maskL = maskL >> 1
		maskR = maskR << 1
		index++
	}
	return result
}
