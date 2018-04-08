# Autogenerated from KST: please remove this line if doing any edits by hand!

import unittest

from params_call_short import _schema

class TestParamsCallShort(unittest.TestCase):
    def test_params_call_short(self):
        r = _schema.parse_file('src/term_strz.bin')
        self.assertEqual(r.buf1.body, u"foo|b")
        self.assertEqual(r.buf2.body, u"ar|ba")
        self.assertEqual(r.buf2.trailer, 122)