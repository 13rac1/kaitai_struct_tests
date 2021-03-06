// Autogenerated from KST: please remove this line if doing any edits by hand!

package spec

import (
	"os"
	"testing"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	. "test_formats"
	"github.com/stretchr/testify/assert"
)

func TestEnumOfValueInst(t *testing.T) {
	f, err := os.Open("../../src/enum_0.bin")
	if err != nil {
		t.Fatal(err)
	}
	s := kaitai.NewStream(f)
	var r EnumOfValueInst
	err = r.Read(s, &r, &r)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, EnumOfValueInst_Animal__Cat, r.Pet1)
	assert.EqualValues(t, EnumOfValueInst_Animal__Chicken, r.Pet2)
	tmp1, err := r.Pet3()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, EnumOfValueInst_Animal__Dog, tmp1)
	tmp2, err := r.Pet4()
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, EnumOfValueInst_Animal__Dog, tmp2)
}
