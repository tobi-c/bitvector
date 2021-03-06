package bitvector_test

import (
	"github.com/tobi-c/bitvector"
	"testing"
)

func TestRank1(t *testing.T) {
	b := []byte{0xFF, 0xFF}
	bv := bitvector.NewBitVector(b, 16)
	for i := uint64(0); i <= 16; i++ {
		r, err := bv.Rank1(i)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r != i {
			t.Errorf("%d(=bv.Rank1(%d)) != %d", r, i, i)
		}
	}

	b = []byte{0x55, 0x55}
	bv = bitvector.NewBitVector(b, 16)
	for i := uint64(0); i <= 16; i++ {
		r, err := bv.Rank1(i)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r != (i+1)/2 {
			t.Errorf("%d(=bv.Rank1(%d)) != %d", r, i, (i+1)/2)
		}
	}

	i := uint64(17)
	r, err := bv.Rank1(i)
	if err == nil {
		t.Errorf("Over Length error")
	}
	if r != 0 {
		t.Errorf("%d(=bv.Rank1(%d)) != %d", r, i, 0)
	}
}

func TestSelect1(t *testing.T) {
	b := []byte{0xFF, 0xFF}
	bv := bitvector.NewBitVector(b, 16)
	for i := uint64(0); i < 16; i++ {
		s, err := bv.Select1(i)
		if err != nil {
			t.Errorf(err.Error())
		}
		if s != i {
			t.Errorf("%d(=bv.Select1(%d)) != %d ", s, i, i)
		}
	}

	b = []byte{0x55, 0x55}
	bv = bitvector.NewBitVector(b, 16)
	for i := uint64(0); i < 8; i++ {
		s, err := bv.Select1(i)
		if err != nil {
			t.Errorf(err.Error())
		}
		if s != i*2 {
			t.Errorf("%d(=bv.Select1(%d)) != %d ", s, i, i*2)
		}
	}

	b = []byte{0xAA, 0xAA}
	bv = bitvector.NewBitVector(b, 16)
	for i := uint64(0); i < 8; i++ {
		s, err := bv.Select1(i)
		if err != nil {
			t.Errorf(err.Error())
		}
		if s != i*2+1 {
			t.Errorf("%d(=bv.Select1(%d)) != %d ", s, i, i*2+1)
		}
	}

	i := uint64(8)
	r, err := bv.Select1(i)
	if err == nil {
		t.Errorf("Over rank error")
	}
	if r != 0 {
		t.Errorf("%d(=bv.Select1(%d)) != %d", r, i, 0)
	}
}

func TestSelect0(t *testing.T) {
	b := []byte{0x00, 0x00}
	bv := bitvector.NewBitVector(b, 16)
	for i := uint64(0); i < 16; i++ {
		s, err := bv.Select0(i)
		if err != nil {
			t.Errorf(err.Error())
		}
		if s != i {
			t.Errorf("%d(=bv.Select0(%d)) != %d ", s, i, i)
		}
	}

	b = []byte{0xAA, 0xAA}
	bv = bitvector.NewBitVector(b, 16)
	for i := uint64(0); i < 8; i++ {
		s, err := bv.Select0(i)
		if err != nil {
			t.Errorf(err.Error())
		}
		if s != i*2 {
			t.Errorf("%d(=bv.Select0(%d)) != %d ", s, i, i*2)
		}
	}

	b = []byte{0x55, 0x55}
	bv = bitvector.NewBitVector(b, 16)
	for i := uint64(0); i < 8; i++ {
		s, err := bv.Select0(i)
		if err != nil {
			t.Errorf(err.Error())
		}
		if s != i*2+1 {
			t.Errorf("%d(=bv.Select0(%d)) != %d ", s, i, i*2+1)
		}
	}

	i := uint64(8)
	r, err := bv.Select0(i)
	if err == nil {
		t.Errorf("Over rank error")
	}
	if r != 0 {
		t.Errorf("%d(=bv.Select0(%d)) != %d", r, i, 0)
	}
	b = []byte{0xFF, 0xFF}
	bv = bitvector.NewBitVector(b, 16)
	r, err = bv.Select0(0)
	if err == nil {
		t.Errorf("Over rank error")
	}
}

var cacheBV = make(map[uint64](*bitvector.BitVector))

func newBitVectorForBench(length uint64) *bitvector.BitVector {
	bv := cacheBV[length]
	if bv == nil {
		s := make([]byte, length)
		for i := 0; i < len(s); i++ {
			s[i] = 0x55
		}
		bv = bitvector.NewBitVector(s, uint64(len(s)*8))
		cacheBV[length] = bv
	}
	return bv
}

func BenchmarkRank1_100000(b *testing.B) {
	bv := newBitVectorForBench(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bv.Rank1(bv.Length - 1)
	}
}

func BenchmarkRank1_100000000(b *testing.B) {
	bv := newBitVectorForBench(100000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bv.Rank1(bv.Length - 1)
	}
}

func BenchmarkSelect1_100000(b *testing.B) {
	bv := newBitVectorForBench(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bv.Select1(bv.Length / 2)
	}
}

func BenchmarkSelect1_100000000(b *testing.B) {
	bv := newBitVectorForBench(100000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bv.Select1(bv.Length / 2)
	}
}

func TestSparseBitVector(t *testing.T) {
	b := []byte{0xFF, 0xFF}
	bitvector.NewSparseBitVector(b, 16)
}

func TestSparseBitVectorRank1(t *testing.T) {
	b := []byte{0xFF, 0xFF}
	bv := bitvector.NewSparseBitVector(b, 16)
	for i := uint64(0); i <= 16; i++ {
		r, err := bv.Rank1(i)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r != i {
			t.Errorf("%d(=bv.Rank1(%d)) != %d", r, i, i)
		}
	}

	b = []byte{0x55, 0x55}
	bv = bitvector.NewSparseBitVector(b, 16)
	for i := uint64(0); i <= 16; i++ {
		r, err := bv.Rank1(i)
		if err != nil {
			t.Errorf(err.Error())
		}
		if r != (i+1)/2 {
			t.Errorf("%d(=bv.Rank1(%d)) != %d", r, i, (i+1)/2)
		}
	}

	i := uint64(17)
	r, err := bv.Rank1(i)
	if err == nil {
		t.Errorf("Over Length error")
	}
	if r != 0 {
		t.Errorf("%d(=bv.Rank1(%d)) != %d", r, i, 0)
	}
}

func TestSparseBitVectorSelect1(t *testing.T) {
	b := []byte{}
	for i := 0; i < 33; i++ {
		b = append(b, 0xFF)
	}
	bv := bitvector.NewSparseBitVector(b, uint64(len(b)*8))
	for r := uint64(1); r < uint64(len(b)*8); r++ {
		s, err := bv.Select1(r)
		if err != nil {
			t.Errorf("Over rank error: %v", err)
		}
		if s != r {
			t.Errorf("err: %d %d ", s, r)
		}
	}

	b = []byte{}
	for i := 0; i < 33; i++ {
		if i%10 == 0 {
			b = append(b, 0x1)
		} else {
			b = append(b, 0x0)
		}
	}
	bv = bitvector.NewSparseBitVector(b, uint64(len(b)*8))
	for r := uint64(1); r < uint64(len(b)/10); r++ {
		s, err := bv.Select1(r)
		if err != nil {
			t.Errorf("Over rank error: %v", err)
		}
		if s != (r * 80) {
			t.Errorf("err: %d %d ", s, r)
		}
	}
}
