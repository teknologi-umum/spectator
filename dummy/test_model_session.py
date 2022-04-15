from datetime import datetime
import unittest

from model_session import (
    generate_personal_info_submitted_event,
    generate_solution_accepted_event,
    generate_solution_rejected_event,
    generate_locale_set_event,
)


class TestGenerateEvents(unittest.TestCase):
    def test_generate_solution_accepted_event(self):
        time = int(datetime.fromisoformat("2020-01-01T00:00:00").timestamp())
        event = generate_solution_accepted_event(
            session_id="GUID",
            time=time,
        )

        self.assertEqual(event["session_id"], "GUID")
        self.assertNotEqual(event["language"], "")
        self.assertNotEqual(event["scratchpad"], "")
        self.assertAlmostEqual(event["serialized_test_results"], "")
        # FIXME: 1577861081 not less than or equal to 6
        self.assertLessEqual(event["question_number"], 6)
        self.assertGreaterEqual(event["question_number"], 1)
        self.assertGreaterEqual(event["time"], time)

    def test_generate_solution_rejected_event(self):
        time = int(datetime.fromisoformat("2020-01-01T00:00:00").timestamp())

        event = generate_solution_rejected_event(session_id="GUID", time=time)

        self.assertEqual(event["session_id"], "GUID")
        self.assertNotEqual(event["language"], "")
        self.assertNotEqual(event["scratchpad"], "")
        self.assertAlmostEqual(event["serialized_test_results"], "")
        # FIXME: 1577843211 not less than or equal to 6
        self.assertLessEqual(event["question_number"], 6)
        self.assertGreaterEqual(event["question_number"], 1)
        self.assertGreaterEqual(event["time"], time)

    def test_generate_locale_set_event(self):
        time = int(datetime.fromisoformat("2020-01-01T00:00:00").timestamp())

        event = generate_locale_set_event(session_id="GUID", time=time)

        self.assertEqual(event["session_id"], "GUID")
        self.assertNotEqual(event["locale"], "")
        # FIXME: event["time"] is a type of integer, expected a date time
        self.assertGreaterEqual(event["time"], time)
        # FIXME: event["time"] is a type of integer, expected a date time

    def test_generate_personal_info_submitted_event(self):
        time = int(datetime.fromisoformat("2020-01-01T00:00:00").timestamp())

        event = generate_personal_info_submitted_event(session_id="GUID", time=time)

        self.assertEqual(event["session_id"], "GUID")
        self.assertNotEqual(event["student_number"], "")
        self.assertGreaterEqual(event["years_of_experience"], 1)
        self.assertLessEqual(event["years_of_experience"], 10)
        self.assertGreaterEqual(event["hours_of_practice"], 0)
        # FIXME: 3534 not less than or equal to 24
        self.assertLessEqual(event["hours_of_practice"], 24)
        self.assertNotEqual(event["familiar_languages"], "")
        self.assertGreaterEqual(event["time"], time)
