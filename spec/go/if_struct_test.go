// Autogenerated from KST: please remove this line if doing any edits by hand!

package spec

import (
	"os"
	"testing"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	. "test_formats"
	"github.com/stretchr/testify/assert"
)

func TestIfStruct(t *testing.T) {
	f, err := os.Open("../../src/if_struct.bin")
	if err != nil {
		t.Fatal(err)
	}
	s := kaitai.NewStream(f)
	var r IfStruct
	err = r.Read(s, &r, &r)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, 83, r.Op1.Opcode)
	assert.EqualValues(t, "foo", r.Op1.ArgStr.Str)
	assert.EqualValues(t, 84, r.Op2.Opcode)
	assert.EqualValues(t, 66, r.Op2.ArgTuple.Num1)
	assert.EqualValues(t, 67, r.Op2.ArgTuple.Num2)
	assert.EqualValues(t, 83, r.Op3.Opcode)
	assert.EqualValues(t, "bar", r.Op3.ArgStr.Str)
}
