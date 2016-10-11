import unittest

from repeat_n_strz_double import RepeatNStrzDouble

class TestRepeatNStrzDouble(unittest.TestCase):
    def test_repeat_n_strz_double(self):
        with RepeatNStrzDouble.from_file("src/repeat_n_strz.bin") as r:

            self.assertEqual(r.qty, 2)
            self.assertEqual(r.lines1, ["foo"])
            self.assertEqual(r.lines2, ["bar"])
