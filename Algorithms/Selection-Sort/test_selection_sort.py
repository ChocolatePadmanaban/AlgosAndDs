import unittest
from selection_sort import selection_sort

class TestBubbleSort(unittest.TestCase):

    def test_already_sorted(self):
        self.assertEqual(selection_sort([1, 2, 3, 4, 5]), [1, 2, 3, 4, 5])

    def test_unsorted_list(self):
        self.assertEqual(selection_sort([5, 4, 3, 2, 1]), [1, 2, 3, 4, 5])

    def test_list_with_duplicates(self):
        self.assertEqual(selection_sort([3, 1, 2, 3, 1]), [1, 1, 2, 3, 3])

    def test_single_element(self):
        self.assertEqual(selection_sort([1]), [1])

    def test_empty_list(self):
        self.assertEqual(selection_sort([]), [])

    def test_negative_numbers(self):
        self.assertEqual(selection_sort([-2, -3, -1, -4, -5]), [-5, -4, -3, -2, -1])

    def test_all_elements_same(self):
        self.assertEqual(selection_sort([2, 2, 2, 2, 2]), [2, 2, 2, 2, 2])

if __name__ == '__main__':
    unittest.main()