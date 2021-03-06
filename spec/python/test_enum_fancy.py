import unittest

from enum_fancy import EnumFancy

class TestEnumFancy(unittest.TestCase):
    def test_enum_fancy(self):
        with EnumFancy.from_file("src/enum_0.bin") as r:

            self.assertEqual(r.pet_1, EnumFancy.Animal.cat)
            self.assertEqual(r.pet_2, EnumFancy.Animal.chicken)
