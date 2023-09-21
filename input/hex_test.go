package input

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_decodeHexString(t *testing.T) {
	spec := []struct {
		input    string
		expected []byte
		err      error
	}{
		{
			input:    "",
			expected: []byte{},
			err:      nil,
		},
		{
			input:    "00",
			expected: []byte{0},
			err:      nil,
		},
		{
			input:    "0x00",
			expected: []byte{0},
			err:      nil,
		},
		{
			input:    "0x0",
			expected: []byte{},
			err:      hex.ErrLength,
		},
	}

	for _, s := range spec {
		t.Run(s.input, func(t *testing.T) {
			actual, err := decodeHexString(s.input)
			assert.Equal(t, s.err, err)
			assert.Equal(t, s.expected, actual)
		})
	}
}

func Benchmark_decodeHexString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decodeHexString("0x12aa3caf0000000000000000000000003208684f96458c540eb08f6f01b9e9afb2b7d4f0000000000000000000000000eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000003208684f96458c540eb08f6f01b9e9afb2b7d4f0000000000000000000000000e1a14ee4daf10465e102389940fa454e90b23c8e00000000000000000000000000000000000000000000003635c9adc5dea000000000000000000000000000000000000000000000000000000000017d5124a7da000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000160000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000009230000000000000000000000000000000000000009050008d70008bd00000600206b4be0b900a0c9e75c480000000000000000060400000000000000000000000000000000000000000000000000088900068000a007e5c0d200000000000000000000000000000000000000000000065c00052c00051200a0c9e75c48000000000016100703020000000000000000000000000000000004e400039600030900027c00016e00a007e5c0d200000000000000000000000000000000000000000000000000014a00001a4041c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2d0e30db051004a585e0f7c18e2c414221d6402652d5e0990e5f8c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200a4a5dcbcdf000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7000000000000000000000000352b186090068eb35d532428676ce510e17ab5810000000000000000000000003208684f96458c540eb08f6f01b9e9afb2b7d4f0ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a007e5c0d20000000000000000000000000000000000000000000000000000ea00001a4041c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2d0e30db05100d51a44d3fae010294c616388b506acda1bfaae46c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20044394747c50000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000a007e5c0d200000000000000000000000000000000000000000000000000006900001a4041c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2d0e30db002a00000000000000000000000000000000000000000000000000000000000000001ee63c1e5016ca298d2983ab03aa1da7679389d955a4efee15cc02aaa39b223fe8d0a0e5c4f27ead9083c756cc200a007e5c0d200000000000000000000000000000000000000000000000000006900001a4041c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2d0e30db002a00000000000000000000000000000000000000000000000000000000000000001ee63c1e50111b815efb8f581194ae79006d24e0d814b7697f6c02aaa39b223fe8d0a0e5c4f27ead9083c756cc200a007e5c0d200000000000000000000000000000000000000000000000000012a00001a4041c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2d0e30db05120d17b3c9784510e33cd5b87b490e79253bcd81e2ec02aaa39b223fe8d0a0e5c4f27ead9083c756cc2004458d30ac9000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec70000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000003208684f96458c540eb08f6f01b9e9afb2b7d4f000000000000000000000000000000000000000000000000000000000651061860020d6bdbf78dac17f958d2ee523a2206206994597c13d831ec751204a585e0f7c18e2c414221d6402652d5e0990e5f8dac17f958d2ee523a2206206994597c13d831ec700a4a5dcbcdf000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000d0b2f5018b5d22759724af6d4281ac0b132663600000000000000000000000003208684f96458c540eb08f6f01b9e9afb2b7d4f0ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a0c9e75c48000000000000000030020000000000000000000000000000000000000000000000000001db00014e00a007e5c0d200000000000000000000000000000000000000000000000000012a00001a4041c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2d0e30db05100d17b3c9784510e33cd5b87b490e79253bcd81e2ec02aaa39b223fe8d0a0e5c4f27ead9083c756cc2004458d30ac9000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000000000000000000000000000003208684f96458c540eb08f6f01b9e9afb2b7d4f0000000000000000000000000000000000000000000000000000000006510618600a007e5c0d200000000000000000000000000000000000000000000000000006900001a4041c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2d0e30db002a00000000000000000000000000000000000000000000000000000000000000001ee63c1e50088e6a0c2ddd26feeb64f039a2c41296fcb3f5640c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20020d6bdbf78a0b86991c6218b36c1d19d4a2e9eb0ce3606eb4880a06c4eca27a0b86991c6218b36c1d19d4a2e9eb0ce3606eb481111111254eeb25477b68fb85ed929f73a96058200000000000000000000000000000000000000000000000000000000008b1ccac8")
	}
}
