package utils

/**
 * 转换[]uint8类型为[]byte 类型
 */
func Bytes(reply any) []byte {
	switch reply := reply.(type) {

	case []uint8:
		return I2B(reply)
	case any:
		return nil
	case nil:
		return nil

	}
	return nil
}

func I2B(bs []uint8) []byte {
	ba := []byte{}
	for _, b := range bs {
		ba = append(ba, b)
	}
	return ba
}
