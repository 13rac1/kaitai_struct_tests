// Autogenerated from KST: please remove this line if doing any edits by hand!

package spec

import (
	"os"
	"testing"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	. "test_formats"
	"github.com/stretchr/testify/assert"
)

func TestNestedSameName2(t *testing.T) {
	f, err := os.Open("../../src/nested_same_name2.bin")
	if err != nil {
		t.Fatal(err)
	}
	s := kaitai.NewStream(f)
	var r NestedSameName2
	err = r.Read(s, &r, &r)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, 66, r.Version)
	assert.EqualValues(t, 2, r.MainData.MainSize)
	assert.EqualValues(t, []uint8{17, 17, 17, 17}, r.MainData.Foo.Data1)
	assert.EqualValues(t, 3, r.Dummy.DummySize)
	assert.EqualValues(t, []uint8{34, 34, 34, 34, 34, 34}, r.Dummy.Foo.Data2)
}
