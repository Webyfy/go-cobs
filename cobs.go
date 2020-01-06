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

	return output[:writeIndex]
}

// Decode a cobs frame to a slice of bytes
func Decode(input []byte) []byte {
	length := len(input)
	if length == 0 {
		return nil
	}
	output := make([]byte, length)
	readIndex := 0
	writeIndex := 0

	var distance byte
	var i byte

	for readIndex < length {
		// Copy the current input value to the distance value
		distance = input[readIndex]

		// If the index of the next distance value is greater than the length of the input
		// AND the distance is not equal to one
		if readIndex+int(distance) > length && distance != 1 {
			return nil
		}

		// Increment to the next not zero value
		readIndex++

		// Copy the input to the output for the distance
		for i = 1; i < distance; i++ {
			output[writeIndex] = input[readIndex]
			writeIndex++
			readIndex++
		}

		// Determine if the
		if distance != 0xFF && readIndex != length {
			output[writeIndex] = 0x00
			writeIndex++
		}
	}

	return output[:writeIndex]
}
