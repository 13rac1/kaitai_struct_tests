// Autogenerated from KST: please remove this line if doing any edits by hand!

package spec

import (
	"os"
	"testing"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	. "test_formats"
	"github.com/stretchr/testify/assert"
)

func TestStrLiterals2(t *testing.T) {
	f, err := os.Open("../../src/fixed_struct.bin")
	if err != nil {
		t.Fatal(err)
	}
	s := kaitai.NewStream(f)
	var r StrLiterals2
	err = r.Read(s, &r, &r)
	if err != nil {
		t.Fatal(err)
	}

	tmp1, err := r.Dollar1()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, "$foo", tmp1)
	tmp2, err := r.Dollar2()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, "${foo}", tmp2)
	tmp3, err := r.Hash()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, "#{foo}", tmp3)
	tmp4, err := r.AtSign()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, "@foo", tmp4)
}
