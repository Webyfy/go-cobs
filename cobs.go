package cobs

// GetEncodedBufferSize calculates the size of encoded message
// from the size of unencoded buffer size
func GetEncodedBufferSize(unencodedBufferSize int) int {
	return unencodedBufferSize + unencodedBufferSize/254 + 1
}

// Encode encodes a byte array using COBS
func Encode(input []byte) []byte {
	if len(input) == 0 {
		return nil
	}
	readIndex := 0
	writeIndex := 1

	codeIndex := 0
	distance := byte(1)

	length := len(input)
	output := make([]byte, GetEncodedBufferSize(length))

	for readIndex < length {
		// If we encounter a zero for the current value of the input
		if input[readIndex] == 0 {
			// Write the value of the distance to the next zero back in output where we last saw a zero
			output[codeIndex] = distance

			// Set the distance index to the latest index plus one
			codeIndex = writeIndex
			writeIndex++

			// Reset the distance
			distance = 1

			// Keep the read and write indexes the same
			readIndex++
		} else {
			// Simply copy the value over from the input (increment the indexes up one value)
			output[writeIndex] = input[readIndex]
			writeIndex++
			readIndex++

			// Increment the distance
			distance++

			// If the distance reaches maximum valve
			if distance == 0xFF {
				// Set the distance variable to its maximum value
				output[codeIndex] = distance

				// Set the distance index to the latest index plus one
				codeIndex = writeIndex
				writeIndex++

				// Reset the distance
				distance = 1
			}
		}
	}

	if codeIndex != 255 && len(output) > 0 {
		output[codeIndex] = distance
	}

	return append(output[:writeIndex], 0)
}

// Decode a cobs frame to a slice of bytes
func Decode(enc []byte) []byte {
	encLen := len(enc)
	if enc[encLen-1] == 0 { // remove trailing 0
		encLen--
	}

	dest := make([]byte, encLen)
	destLen := encLen
	ptr := 0
	pos := 0

	if encLen == 0 {
		return nil
	}

	for ptr < encLen {
		code := enc[ptr]

		if ptr+int(code) > encLen {
			return nil
		}
		ptr++

		if pos+int(code) > destLen {
			return nil
		}

		for i := 1; i < int(code); i++ {
			dest[pos] = enc[ptr]
			pos++
			ptr++
		}
		if code < 0xFF {
			dest[pos] = 0
			pos++
		}
	}

	return dest[:pos-1] // trim phantom zero
}
