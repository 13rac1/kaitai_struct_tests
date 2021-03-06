<?php
// Autogenerated from KST: please remove this line if doing any edits by hand!

namespace Kaitai\Struct\Tests;

class SwitchIntegers2Test extends TestCase {
    public function testSwitchIntegers2() {
        $r = SwitchIntegers2::fromFile(self::SRC_DIR_PATH . '/switch_integers.bin');

        $this->assertEquals(1, $r->code());
        $this->assertEquals(7, $r->len());
        $this->assertEquals("\x02\x40\x40\x04\x37\x13\x00", $r->ham());
        $this->assertEquals(0, $r->padding());
        $this->assertEquals("13", $r->lenModStr());
    }
}
