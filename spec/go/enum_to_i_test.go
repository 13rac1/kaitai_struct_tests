// Autogenerated from KST: please remove this line if doing any edits by hand!

package spec

import (
	"os"
	"testing"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	. "test_formats"
	"github.com/stretchr/testify/assert"
)

func TestEnumToI(t *testing.T) {
	f, err := os.Open("../../src/enum_0.bin")
	if err != nil {
		t.Fatal(err)
	}
	s := kaitai.NewStream(f)
	var r EnumToI
	err = r.Read(s, &r, &r)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, EnumToI_Animal__Cat, r.Pet1)
	assert.EqualValues(t, EnumToI_Animal__Chicken, r.Pet2)
	tmp1, err := r.Pet1I()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, 7, tmp1)
	tmp2, err := r.OneLtTwo()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, true, tmp2)
}
