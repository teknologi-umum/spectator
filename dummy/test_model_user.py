import unittest

from model_user import generate_user

class TestGenerateUser(unittest.TestCase):
    def test_generate_user(self):
        user = generate_user()

        self.assertNotEqual(user["session_id"], "")
        self.assertEqual(len(user["student_number"]), 8)
        self.assertGreaterEqual(user["hours_of_practice"], 0)
        self.assertLessEqual(user["hours_of_practice"], 5)
        self.assertGreaterEqual(user["years_of_experience"], 0)
        self.assertLessEqual(user["years_of_experience"], 5)
        self.assertNotEqual(user["familiar_languages"], "")
