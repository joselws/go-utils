package mytypes

/*
rune and byte data types are added implicitly due to them being alias for
int32 and uint8 respectively.
*/
type Number interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}
