package stringcompression

import "strconv"

func compress(chars []byte) int {
	writeIndex, readIndex := 0, 0
	for readIndex < len(chars) {
		curr := chars[readIndex]
		count := 0

		for readIndex < len(chars) && chars[readIndex] == curr {
			readIndex++
			count++
		}

		chars[writeIndex] = curr
		writeIndex++

		if count > 1 {
			for _, char := range strconv.Itoa(count) {
				chars[writeIndex] = byte(char)
				writeIndex++
			}
		}
	}
	return writeIndex
}
