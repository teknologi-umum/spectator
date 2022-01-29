import random
import datetime
from utils import random_date

_MouseButton=[
    "Left",
    "Right",
    "Middle"
]
class InputEventBase:
    session_id: str
    type: str
    _time: int
    question_number: str

    def __init__(self, session_id: str, time: int, question_number: int) -> None:
        self.session_id = session_id
        self._time = time
        self.question_number = question_number
    def as_dictionary(self)-> dict[str, any]:
        return {
            "type": self.type,
            "session_id": self.session_id,
            "time": self._time,
            "question_number": self.question_number,
        }
class EventKeystroke(InputEventBase):
    key_char: str
    key_code: str
    shift: bool
    alt: bool
    control: bool
    meta: bool
    unrelated_key: bool
    def __init__(self, session_id: str, question_number: int, key_char: str,
                 key_code: str, shift: bool, alt: bool, control: bool,
                 meta: bool,
                 unrelated_key: bool, time: int) -> None:
        super().__init__(session_id, time, question_number)
        self.type = "coding_event_keystroke"
        self.key_char = key_char
        self.key_code = key_code
        self.shift = shift
        self.alt = alt
        self.control = control
        self.meta = meta
        self.unrelated_key = unrelated_key

    def as_dictionary(self) -> dict[str,any]:
        return {
            "key_char": self.key_char,
            "key_code": self.key_code,
            "shift": self.shift,
            "alt": self.alt,
            "control": self.control,
            "meta": self.meta,
            "unrelated_key": self.unrelated_key,
        } | super().as_dictionary()

class EventMouseMove(InputEventBase):
    # direction: "up" | "down" | "left" | "right"
    direction: str
    x_position: int
    y_position: int
    window_width: int
    window_height: int

    def __init__(self, session_id: str, question_number: int, direction: str,
                 x_position: int, y_position: int, time: int) -> None:
        super().__init__(session_id, time, question_number)
        self.type = "coding_event_mousemove"
        self.direction = direction
        self.x_position = x_position
        self.y_position = y_position

    def as_dictionary(self):
        return {
            "direction": self.direction,
            "x": self.x_position,
            "y": self.y_position,
        } | super().as_dictionary()

class EventMouseClick(InputEventBase):
    button: str
    x_position: int
    y_position: int
    _time: int

    def __init__(self, session_id: str, question_number: str, button: str,
                 x_position: int, y_position: int, time: int) -> None:
        super().__init__(session_id, time, question_number)
        self.type = "coding_event_mouseclick"
        self.x_position = x_position
        self.y_position = y_position
        self.button = button

    def as_dictionary(self):
        return {
            "button": self.button,
            "x": self.x_position,
            "y": self.y_position,
        } | super().as_dictionary()

class EventWindowSized(InputEventBase):
    width: int
    height: int

    def __init__(self, session_id: str, question_number: str, width: int,
                 height: int, time: int) -> None:
        super().__init__(session_id, time, question_number)
        self.type = "window_sized"
        self.width = width
        self.height = height

    def as_dictionary(self):
        return {
            "width": self.width,
            "height": self.height,
        } | super().as_dictionary()

def generate_event_keystroke(session_id: str, date_start: datetime,
                             date_ends: datetime) -> dict[str, any]:
    """Generate an EventKeystroke class with
    random values.

    Args:
        session_id (str): [desc]
        date_start (datetime): [desc]
        date_ends (datetime): [desc]

    Returns:
        dict[str, any]: [description]
    """
    question_number = random.randint(1, 6)
    key_code = random.randint(0, 255)
    key_char = chr(key_code)
    shift = random.choice([True, False])
    alt = random.choice([True, False])
    control = random.choice([True, False])
    meta = random.choice([True, False])
    unrelated_key = random.choice([True, False])
    time = random_date(date_start, date_ends)

    return (EventKeystroke(session_id, question_number, key_char, key_code,
                           shift, alt, control, meta, unrelated_key,
                           time)).as_dictionary()

def generate_event_mousemove(session_id: str, date_start: datetime,
                             date_ends: datetime) -> dict[str, any]:
    """Generate an EventMouseMove class with random values.
    The "direction" key may only be either "up", "down", "left" or "right".

    Args:
        session_id (str): [description]
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
                           y_position, time)).as_dictionary()

def generate_event_mouseclick(session_id: str, date_start: datetime,
                              date_ends: datetime) -> dict[str, any]:
    """Generate an EventMouseClick class with random values.

    Args:
        session_id (str): [description]
        date_start (datetime): [description]
        date_ends (datetime): [description]

    Returns:
        dict[str, any]: [description]
    """
    question_number = random.randint(1, 6)
    window_height = random.randint(0, 1080)
    window_width = random.randint(0, 1920)
    x_position = random.randint(0, window_width)
    y_position = random.randint(0, window_height)
    button = random.choice(_MouseButton)
    time = random_date(date_start, date_ends)

    return (EventMouseClick(session_id, str(question_number), button, x_position,
                            y_position, time)).as_dictionary()

def generate_event_window_sized(session_id: str, date_start: datetime,
                                date_ends: datetime) -> dict[str, any]:
    """Generate an EventWindowSized class with random values.

    Args:
        session_id (str): a user's session id
        date_start (datetime): the start date range
        date_ends (datetime): the end date range

    Returns:
        dict[str, any]: [description]
    """
    question_number = random.randint(1, 6)
    width = random.randint(400, 1920)
    height = random.randint(200, 1080)
    time = random_date(date_start, date_ends)

    return (EventWindowSized(session_id, str(question_number), width, height,
                             time)).as_dictionary()
