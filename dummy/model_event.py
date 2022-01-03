import random
import datetime
from utils import random_date

class EventKeystroke:
    session_id: str
    type: str
    question_number: int
    key_char: str
    key_code: str
    shift: bool
    alt: bool
    control: bool
    meta: bool
    unrelated_key: bool
    _time: int

    def __init__(self, session_id: str, question_number: int, key_char: str,
                 key_code: str, shift: bool, alt: bool, control: bool, meta: bool,
                 unrelated_key: bool, time: int) -> None:
        self.type = "coding_event_keystroke"
        self.session_id = session_id
        self.question_number = question_number
        self.key_char = key_char
        self.key_code = key_code
        self.shift = shift
        self.alt = alt
        self.control = control
        self.meta = meta
        self.unrelated_key = unrelated_key
        self._time = time

    def asdict(self):
        return {
            "type": self.type,
            "session_id": self.session_id,
            "question_number": self.question_number,
            "key_char": self.key_char,
            "key_code": self.key_code,
            "shift": self.shift,
            "alt": self.alt,
            "control": self.control,
            "meta": self.meta,
            "unrelated_key": self.unrelated_key,
            "time": self._time
        }

class EventMouseMove:
    session_id: str
    type: str
    question_number: int
    # direction: "up" | "down" | "left" | "right"
    direction: str
    x_position: int
    y_position: int
    window_width: int
    window_height: int
    _time: int

    def __init__(self, session_id: str, question_number: int, direction: str,
                 x_position: int, y_position: int, time: int) -> None:
        self.type = "coding_event_mousemove"
        self.session_id = session_id
        self.question_number = question_number
        self.direction = direction
        self.x_position = x_position
        self.y_position = y_position
        self._time = time

    def asdict(self):
        return {
            "type": self.type,
            "session_id": self.session_id,
            "question_number": self.question_number,
            "direction": self.direction,
            "x": self.x_position,
            "y": self.y_position,
            "time": self._time
        }

class EventMouseClick:
    session_id: str
    type: str
    question_number: str
    right_click: bool
    left_click: bool
    middle_click: bool
    _time: int

    def __init__(self, session_id: str, question_number: str, right_click: bool,
                 left_click: bool, middle_click: bool, time: int) -> None:
        self.type = "coding_event_mouseclick"
        self.session_id = session_id
        self.question_number = question_number
        self.right_click = right_click
        self.left_click = left_click
        self.middle_click = middle_click
        self._time = time

    def asdict(self):
        return {
            "type": self.type,
            "session_id": self.session_id,
            "question_number": self.question_number,
            "right_click": self.right_click,
            "left_click": self.left_click,
            "middle_click": self.middle_click,
            "time": self._time
        }

class EventWindowSized:
    session_id: str
    type: str
    question_number: int
    width: int
    height: int
    _time: int

    def __init__(self, session_id: str, question_number: int, width: int,
                 height: int, time: int) -> None:
        self.type = "window_sized"
        self.session_id = session_id
        self.question_number = question_number
        self.width = width
        self.height = height
        self._time = time

    def asdict(self):
        return {
            "type": self.type,
            "session_id": self.session_id,
            "question_number": self.question_number,
            "width": self.width,
            "height": self.height,
            "time": self._time
        }

def generate_event_keystroke(session_id: str, date_start: datetime, date_ends: datetime) -> dict[str, any]:
    """Generate an EventKeystroke class with
    random values.

    Returns:
        dict[str, any]: [description]
    """
    question_number = random.randint(1, 6)
    key_code = random.randint(65, 255)
    key_char = chr(key_code)
    shift = random.choice([True, False])
    alt = random.choice([True, False])
    control = random.choice([True, False])
    meta = random.choice([True, False])
    unrelated_key = random.choice([True, False])
    time = random_date(date_start, date_ends)

    return (EventKeystroke(session_id, question_number, key_char, key_code,
                          shift, alt, control, meta, unrelated_key, time)).asdict()

def generate_event_mousemove(session_id: str, date_start: datetime, date_ends: datetime) -> dict[str, any]:
    """Generate an EventMouseMove class with random values.
    The "direction" key may only be either "up", "down", "left" or "right".

    Args:
        date_start (datetime): [description]
        date_ends (datetime): [description]

    Returns:
        dict[str, any]: [description]
    """
    question_number = random.randint(1, 6)
    direction = random.choice(["up", "down", "left", "right"])
    window_height = random.randint(0, 1080)
    window_width = random.randint(0, 1920)
    x_position = random.randint(0, window_width)
    y_position = random.randint(0, window_height)
    time = random_date(date_start, date_ends)

    return (EventMouseMove(session_id, question_number, direction, x_position,
                           y_position, time)).asdict()

def generate_event_mouseclick(session_id: str, date_start: datetime, date_ends: datetime) -> dict[str, any]:
    """Generate an EventMouseClick class with random values.

    Args:
        date_start (datetime): [description]
        date_ends (datetime): [description]

    Returns:
        dict[str, any]: [description]
    """
    question_number = random.randint(1, 6)
    right_click = random.choice([True, False])
    left_click = random.choice([True, False])
    middle_click = random.choice([True, False])
    time = random_date(date_start, date_ends)

    return (EventMouseClick(session_id, question_number, right_click, left_click,
                            middle_click, time)).asdict()

def generate_event_window_sized(session_id: str, date_start: datetime, date_ends: datetime) -> dict[str, any]:
    """Generate an EventWindowSized class with random values.

    Args:
        date_start (datetime): [description]
        date_ends (datetime): [description]

    Returns:
        dict[str, any]: [description]
    """
    question_number = random.randint(1, 6)
    width = random.randint(400, 1920)
    height = random.randint(200, 1080)
    time = random_date(date_start, date_ends)

    return (EventWindowSized(session_id, question_number, width, height, time)).asdict()
