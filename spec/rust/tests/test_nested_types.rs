// Autogenerated from KST: please remove this line if doing any edits by hand!

extern crate kaitai_struct;
extern crate rust;

use kaitai_struct::KaitaiStruct;
use rust::NestedTypes;

#[test]
fn test_nested_types() {
    if let Ok(r) = NestedTypes::from_file("src/fixed_struct.bin") {
        assert_eq!(r.one.typed_at_root.value_b, 80);
        assert_eq!(r.one.typed_here.value_c, 65);
        assert_eq!(r.two.value_b, 67);
    }
}