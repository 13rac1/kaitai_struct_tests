import unittest

from multiple_use import MultipleUse

class TestMultipleUse(unittest.TestCase):
    def test_multiple_use(self):
        r = MultipleUse.from_file("src/position_abs.bin")

        self.assertEquals(r.t1.first_use.value, 0x20)
        self.assertEquals(r.t2.second_use.value, 0x20)
