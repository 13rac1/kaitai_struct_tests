// Autogenerated from KST: please remove this line if doing any edits by hand!

package spec

import (
	"os"
	"testing"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	. "test_formats"
	"github.com/stretchr/testify/assert"
)

func TestNavParent(t *testing.T) {
	f, err := os.Open("../../src/nav.bin")
	if err != nil {
		t.Fatal(err)
	}
	s := kaitai.NewStream(f)
	var r NavParent
	err = r.Read(s, &r, &r)
	if err != nil {
		t.Fatal(err)
	}

	assert.EqualValues(t, 2, r.Header.QtyEntries)
	assert.EqualValues(t, 8, r.Header.FilenameLen)
	assert.EqualValues(t, 2, len(r.Index.Entries))
	assert.EqualValues(t, "FIRST___", r.Index.Entries[0].Filename)
	assert.EqualValues(t, "SECOND__", r.Index.Entries[1].Filename)
}
