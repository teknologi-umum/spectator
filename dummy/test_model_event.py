from datetime import datetime
import random
import unittest

from model_event import (
    generate_keystroke_event,
    generate_mouseclick_event,
    generate_mousemove_event,
    generate_window_sized_event,
)


class TestGenerateEvents(unittest.TestCase):
    def test_generate_keystroke_event(self):
        time = int(datetime.fromisoformat("2020-01-01T00:00:00").timestamp())

        event = generate_keystroke_event(session_id="GUID", time=time)

        self.assertEqual(event["type"], "keystroke")
        self.assertEqual(event["session_id"], "GUID")
        self.assertNotIn(event["key_char"], ["", None])
        self.assertNotIn(event["key_code"], ["", None])
        self.assertIn(event["shift"], [True, False])
        self.assertIn(event["alt"], [True, False])
        self.assertIn(event["control"], [True, False])
        self.assertIn(event["meta"], [True, False])
        self.assertIn(event["unrelated_key"], [True, False])
        self.assertEqual(event["time"], time)
        self.assertLessEqual(event["question_number"], 6)
        self.assertGreaterEqual(event["question_number"], 1)

    def test_generate_mousemove_event(self):
        time = int(datetime.fromisoformat("2020-01-01T00:00:00").timestamp())

        event = generate_mousemove_event(session_id="GUID", time=time)

        self.assertEqual(event["type"], "mouse_moved")
        self.assertIn(event["direction"], ["up", "down", "left", "right"])
        self.assertGreaterEqual(event["x"], 0)
        self.assertLessEqual(event["x"], 1920)
        self.assertGreaterEqual(event["y"], 0)
        self.assertLessEqual(event["y"], 1080)
        self.assertEqual(event["session_id"], "GUID")
        self.assertEqual(event["time"], time)
        self.assertLessEqual(event["question_number"], 6)
        self.assertGreaterEqual(event["question_number"], 1)

    def test_generate_mouseclick_event(self):
        time = int(datetime.fromisoformat("2020-01-01T00:00:00").timestamp())

        events = generate_mouseclick_event(session_id="GUID", time=time)

        self.assertEqual(len(events), 2)

        # the first event should be mouse_down
        self.assertEqual(events[0]["type"], "mouse_down")
        self.assertEqual(events[0]["time"], time)

        # and should always be followed by mouse_up
        self.assertEqual(events[1]["type"], "mouse_up")

        # the interval should be around 1 - 5 seconds in milliseconds precision
        self.assertTrue(time <= events[1]["time"] <= time + 5000)

        for event in events:
            self.assertIn(event["button"], ["Left", "Right", "Middle"])
            self.assertGreaterEqual(event["x"], 0)
            self.assertLessEqual(event["x"], 1920)
            self.assertGreaterEqual(event["y"], 0)
            self.assertLessEqual(event["y"], 1080)
            self.assertEqual(event["session_id"], "GUID")
            self.assertLessEqual(event["question_number"], 6)
            self.assertGreaterEqual(event["question_number"], 1)

    def test_generate_window_sized_event(self):
        time = int(datetime.fromisoformat("2020-01-01T00:00:00").timestamp())

        event = generate_window_sized_event(session_id="GUID", time=time)

        self.assertEqual(event["type"], "window_sized")
        self.assertGreaterEqual(event["width"], 400)
        self.assertLessEqual(event["width"], 1920)
        self.assertGreaterEqual(event["height"], 200)
        self.assertLessEqual(event["height"], 1080)
        self.assertEqual(event["session_id"], "GUID")
        self.assertEqual(event["time"], time)
        self.assertLessEqual(event["question_number"], 6)
        self.assertGreaterEqual(event["question_number"], 1)
