from datetime import datetime
import unittest

from model_session import generate_event_locale_set, generate_event_personal_info_submitted, generate_event_solution_accepted, generate_event_solution_rejected


class TestGenerateEvents(unittest.TestCase):
    def test_generate_event_solution_accepted(self):
        date_start =datetime.fromisoformat("2020-01-01T00:00:00")
        date_ends = datetime.fromisoformat("2020-01-02T00:00:00")
        event = generate_event_solution_accepted(
            session_id="GUID",
            date_start=date_start,
            date_ends=date_ends,
        )

        self.assertEqual(event["session_id"], "GUID")
        self.assertNotEqual(event["language"], "")
        self.assertNotEqual(event["scratchpad"], "")
        self.assertAlmostEqual(event["serialized_test_results"], "")
        # FIXME: 1577861081 not less than or equal to 6
        self.assertLessEqual(event["question_number"], 6)
        self.assertGreaterEqual(event["question_number"], 1)
        self.assertGreaterEqual(event["time"], date_start)
        self.assertLessEqual(event["time"], date_ends)

    def test_generate_event_solution_rejected(self):
        date_start = datetime.fromisoformat("2020-01-01T00:00:00")
        date_ends = datetime.fromisoformat("2020-01-02T00:00:00")

        event = generate_event_solution_rejected(
            session_id="GUID",
            date_start=date_start,
            date_ends=date_ends
        )

        self.assertEqual(event["session_id"], "GUID")
        self.assertNotEqual(event["language"], "")
        self.assertNotEqual(event["scratchpad"], "")
        self.assertAlmostEqual(event["serialized_test_results"], "")
        # FIXME: 1577843211 not less than or equal to 6
        self.assertLessEqual(event["question_number"], 6)
        self.assertGreaterEqual(event["question_number"], 1)
        self.assertGreaterEqual(event["time"], date_start)
        self.assertLessEqual(event["time"], date_ends)

    def test_generate_event_locale_set(self):
        date_start = datetime.fromisoformat("2020-01-01T00:00:00")
        date_ends = datetime.fromisoformat("2020-01-02T00:00:00")

        event = generate_event_locale_set(
            session_id="GUID",
            date_start=date_start,
            date_ends=date_ends,
        )

        self.assertEqual(event["session_id"], "GUID")
        self.assertNotEqual(event["locale"], "")
        # FIXME: event["time"] is a type of integer, expected a date time
        self.assertGreaterEqual(event["time"], date_start)
        # FIXME: event["time"] is a type of integer, expected a date time
        self.assertLessEqual(event["time"], date_ends)

    def test_generate_event_personal_info_submitted(self):
        date_start = datetime.fromisoformat("2020-01-01T00:00:00")
        date_ends = datetime.fromisoformat("2020-01-02T00:00:00")

        event = generate_event_personal_info_submitted(
            session_id="GUID",
            date_start=date_start,
            date_ends=date_ends,
        )

        self.assertEqual(event["session_id"], "GUID")
        self.assertNotEqual(event["student_number"], "")
        self.assertGreaterEqual(event["years_of_experience"], 1)
        self.assertLessEqual(event["years_of_experience"], 10)
        self.assertGreaterEqual(event["hours_of_practice"], 0)
        # FIXME: 3534 not less than or equal to 24
        self.assertLessEqual(event["hours_of_practice"], 24)
        self.assertNotEqual(event["familiar_languages"], "")
        self.assertGreaterEqual(event["time"], date_start)
        self.assertLessEqual(event["time"], date_ends)
