from datetime import datetime
import unittest

from model_session import (
    fake_basic_info_generate,
    generate_after_exam_SAM_submited_event,
    generate_before_exam_SAM_submited_event,
    generate_deadline_passed_event,
    generate_exam_ended_event,
    generate_exam_forfeited_event,
    generate_exam_ide_reloaded_event,
    generate_exam_started_event,
    generate_personal_info_submitted_event,
    generate_session_started_event,
    generate_solution_accepted_event,
    generate_solution_rejected_event,
    generate_locale_set_event,
)


class TestGenerateSessions(unittest.TestCase):
    def test_generate_solution_accepted_event(self):
        time = int(datetime.fromisoformat("2020-01-01T00:00:00").timestamp())
        event = generate_solution_accepted_event(
            session_id="GUID",
            time=time,
        )

        self.assertEqual(event["session_id"], "GUID")
        self.assertIn(event["language"], ["Undefined", "C", "CPP", "PHP", "Javascript", "Java", "Python"])
        self.assertNotEqual(event["scratchpad"], "")
        self.assertAlmostEqual(event["serialized_test_results"], "")
        self.assertLessEqual(event["question_number"], 6)
        self.assertGreaterEqual(event["question_number"], 1)
        self.assertEqual(event["time"], time)

    def test_generate_solution_rejected_event(self):
        time = int(datetime.fromisoformat("2020-01-01T00:00:00").timestamp())

        event = generate_solution_rejected_event(session_id="GUID", time=time)

        self.assertEqual(event["type"], "solution_rejected")
        self.assertEqual(event["session_id"], "GUID")
        self.assertIn(event["language"], ["Undefined", "C", "CPP", "PHP", "Javascript", "Java", "Python"])
        self.assertNotEqual(event["scratchpad"], "")
        self.assertAlmostEqual(event["serialized_test_results"], "")
        self.assertLessEqual(event["question_number"], 6)
        self.assertGreaterEqual(event["question_number"], 1)
        self.assertEqual(event["time"], time)

    def test_generate_locale_set_event(self):
        time = int(datetime.fromisoformat("2020-01-01T00:00:00").timestamp())

        event = generate_locale_set_event(session_id="GUID", time=time)

        self.assertEqual(event["type"], "locale_set")
        self.assertEqual(event["session_id"], "GUID")
        self.assertIn(event["locale"], [0, 1])
        self.assertEqual(event["time"], time)

    def test_generate_personal_info_submitted_event(self):
        time = int(datetime.fromisoformat("2020-01-01T00:00:00").timestamp())
        user = {
            "session_id": "GUID",
            "student_number": "123456789",
            "hours_of_practice": 5,
            "years_of_experience": 1,
            "familiar_languages": "python javascript",
        }
        event = generate_personal_info_submitted_event(user=user, time=time)

        self.assertEqual(event["type"], "personal_info_submitted")
        self.assertEqual(event["session_id"], "GUID")
        self.assertEqual(event["student_number"], "123456789")
        self.assertEqual(event["years_of_experience"], 1)
        self.assertEqual(event["hours_of_practice"], 5)
        self.assertEqual(event["familiar_languages"], "python javascript")
        self.assertEqual(event["time"], time)

    def test_generate_session_started_event(self):
        time = int(datetime.fromisoformat("2020-01-01T00:00:00").timestamp())

        event = generate_session_started_event(session_id="GUID", time=time)

        self.assertEqual(event["type"], "session_started")
        self.assertEqual(event["session_id"], "GUID")
        self.assertEqual(event["time"], time)
        self.assertIn(event["locale"], [1, 0])

    def test_generate_deadline_passed_event(self):
        time = int(datetime.fromisoformat("2020-01-01T00:00:00").timestamp())

        event = generate_deadline_passed_event(session_id="GUID", time=time)

        self.assertEqual(event["type"], "deadline_passed")
        self.assertEqual(event["session_id"], "GUID")
        self.assertEqual(event["time"], time)

    def test_generate_exam_ended_event(self):
        time = int(datetime.fromisoformat("2020-01-01T00:00:00").timestamp())

        event = generate_exam_ended_event(session_id="GUID", time=time)

        self.assertEqual(event["type"], "exam_ended")
        self.assertEqual(event["session_id"], "GUID")
        self.assertEqual(event["time"], time)

    def test_generate_exam_forfeited_event(self):
        time = int(datetime.fromisoformat("2020-01-01T00:00:00").timestamp())

        event = generate_exam_forfeited_event(session_id="GUID", time=time)

        self.assertEqual(event["type"], "exam_forfeited")
        self.assertEqual(event["session_id"], "GUID")
        self.assertEqual(event["time"], time)

    def test_generate_exam_ide_reloaded_event(self):
        time = int(datetime.fromisoformat("2020-01-01T00:00:00").timestamp())

        event = generate_exam_ide_reloaded_event(session_id="GUID", time=time)

        self.assertEqual(event["type"], "exam_ide_reloaded")
        self.assertEqual(event["session_id"], "GUID")
        self.assertEqual(event["time"], time)

    def test_generate_exam_started_event(self):
        time = int(datetime.fromisoformat("2020-01-01T00:00:00").timestamp())

        event = generate_exam_started_event(session_id="GUID", time=time)

        self.assertEqual(event["type"], "exam_started")
        self.assertEqual(event["session_id"], "GUID")
        self.assertEqual(event["time"], time)
        self.assertEqual(event["deadline"], time + (90 * 60 * 1000))
        self.assertEqual(event["question_numbers"], [1, 2, 3, 4, 5, 6])

    def test_generate_before_exam_SAM_submitted_event(self):
        time = int(datetime.fromisoformat("2020-01-01T00:00:00").timestamp())

        event = generate_before_exam_SAM_submited_event(session_id="GUID", time=time)

        self.assertEqual(event["type"], "before_exam_sam_submitted")
        self.assertEqual(event["session_id"], "GUID")
        self.assertEqual(event["time"], time)
        self.assertLessEqual(event["aroused_level"], 5)
        self.assertGreaterEqual(event["aroused_level"], 0)
        self.assertLessEqual(event["pleased_level"], 5)
        self.assertGreaterEqual(event["pleased_level"], 0)

    def test_generate_after_exam_SAM_submitted_event(self):
        time = int(datetime.fromisoformat("2020-01-01T00:00:00").timestamp())

        event = generate_after_exam_SAM_submited_event(session_id="GUID", time=time)

        self.assertEqual(event["type"], "after_exam_sam_submitted")
        self.assertEqual(event["session_id"], "GUID")
        self.assertEqual(event["time"], time)
        self.assertLessEqual(event["aroused_level"], 5)
        self.assertGreaterEqual(event["aroused_level"], 0)
        self.assertLessEqual(event["pleased_level"], 5)
        self.assertGreaterEqual(event["pleased_level"], 0)


