#!/usr/bin/env python3

import unittest
import main

class TestMain(unittest.TestCase):

    def test_floor_processor(self):
        floor, firstBasementFloor = main.parse_input(')))')
        self.assertEqual(firstBasementFloor, 1)
        self.assertEqual(floor, -3)

if __name__ == '__main__':
    unittest.main()
