// Autogenerated from KST: please remove this line if doing any edits by hand!

package spec

import (
	"os"
	"testing"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	. "test_formats"
	"github.com/stretchr/testify/assert"
)

func TestBytesPadTerm(t *testing.T) {
	f, err := os.Open("../../src/str_pad_term.bin")
	if err != nil {
		t.Fatal(err)
	}
	s := kaitai.NewStream(f)
	var r BytesPadTerm
	err = r.Read(s, &r, &r)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, []uint8{115, 116, 114, 49}, r.StrPad)
	assert.EqualValues(t, []uint8{115, 116, 114, 50, 102, 111, 111}, r.StrTerm)
	assert.EqualValues(t, []uint8{115, 116, 114, 43, 43, 43, 51, 98, 97, 114, 43, 43, 43}, r.StrTermAndPad)
	assert.EqualValues(t, []uint8{115, 116, 114, 52, 98, 97, 122, 64}, r.StrTermInclude)
}
