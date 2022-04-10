from datetime import datetime
import unittest

from model_session import generate_event_solution_accepted


class TestGenerateEvents(unittest.TestCase):
    def test_generate_event_solution_accepted(self):
        event = generate_event_solution_accepted(
            session_id="GUID",
            date_start=datetime.fromisoformat("2020-01-01T00:00:00"),
            date_ends=datetime.fromisoformat("2020-01-02T00:00:00"),
        )

        self.assertEqual(event["session_id"], "GUID")
        self.assertNotEqual(event["language"], "")
        self.assertNotEqual(event["scratchpad"], "")
        self.assertAlmostEqual(event["serialized_test_results"], "")
