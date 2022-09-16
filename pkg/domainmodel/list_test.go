package domainmodel

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartitionedList(t *testing.T) {

	a := "a"
	b := "b"
	c := "c"
	d := "d"
	e := "e"
	f := "f"
	g := "g"

	args := []struct {
		List []*string
		ChunkSize int
		Expect [][]*string
	}{
		{[]*string{&a, &b, &c},1, [][]*string{{&a},{&b},{&c}} },
		{[]*string{&a, &b, &c},2, [][]*string{{&a,  &b},{&c}} },
		{[]*string{&a, &b, &c, &d, &e, &f},1, [][]*string{{&a},{&b},{&c},{&d},{&e},{&f}} },
		{[]*string{&a, &b, &c, &d, &e, &f},2, [][]*string{{&a,  &b},{&c,  &d},{&e,  &f}} },
		{[]*string{&a, &b, &c, &d, &e, &f},3, [][]*string{{&a,  &b,  &c},{&d,  &e,  &f}} },
		{[]*string{&a, &b, &c, &d, &e, &f},6, [][]*string{{&a,  &b,  &c,  &d,  &e,  &f}} },
		{[]*string{&a, &b, &c, &d, &e, &f, &g},4, [][]*string{{&a,  &b,  &c, &d}, {&e,  &f, &g}} },
	}

	for _, arg := range args {
		t.Run("", func(t *testing.T) {
			actual := PartitionedList[string](arg.List, arg.ChunkSize)
			assert.Equal(t, actual, arg.Expect)
		})
	}
}

