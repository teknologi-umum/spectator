import random
from generate_key_event import event_which_to_event_code
import numpy as np

MOUSE_BUTTON = ["Left", "Right", "Middle"]
MOUSE_DIRECTIONS = ["up", "down", "left", "right"]

rng = np.random.default_rng()


class InputEventBase:
    session_id: str
    type: str
    _time: int
    question_number: int

    def __init__(self, session_id: str, time: int, question_number: int) -> None:
        self.session_id = session_id
        self._time = time
        self.question_number = question_number

    def as_dictionary(self) -> dict[str, any]:
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

    def __init__(
        self,
        session_id: str,
        question_number: int,
        key_char: str,
        key_code: str,
        shift: bool,
        alt: bool,
        control: bool,
        meta: bool,
        unrelated_key: bool,
        time: int,
    ) -> None:
        super().__init__(session_id, time, question_number)
        self.type = "keystroke"
        self.key_char = key_char
        self.key_code = key_code
        self.shift = shift
        self.alt = alt
        self.control = control
        self.meta = meta
        self.unrelated_key = unrelated_key

    def as_dictionary(self) -> dict[str, any]:
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

    def __init__(
        self,
        session_id: str,
        question_number: int,
        direction: str,
        x_position: int,
        y_position: int,
        time: int,
    ) -> None:
        super().__init__(session_id, time, question_number)
        self.type = "mouse_moved"
        self.direction = direction
        self.x_position = x_position
        self.y_position = y_position

    def as_dictionary(self):
        return {
            "direction": self.direction,
            "x": self.x_position,
            "y": self.y_position,
        } | super().as_dictionary()


class EventMouseDown(InputEventBase):
    button: str
    x_position: int
    y_position: int
    _time: int

    def __init__(
        self,
        session_id: str,
        question_number: int,
        button: str,
        x_position: int,
        y_position: int,
        time: int,
    ) -> None:
        super().__init__(session_id, time, question_number)
        self.type = "mouse_down"
        self.x_position = x_position
        self.y_position = y_position
        self.button = button

    def as_dictionary(self):
        return {
            "button": self.button,
            "x": self.x_position,
            "y": self.y_position,
        } | super().as_dictionary()


class EventMouseUp(InputEventBase):
    button: str
    x_position: int
    y_position: int
    _time: int

    def __init__(
        self,
        session_id: str,
        question_number: int,
        button: str,
        x_position: int,
        y_position: int,
        time: int,
    ) -> None:
        super().__init__(session_id, time, question_number)
        self.type = "mouse_up"
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

    def __init__(
        self, session_id: str, question_number: int, width: int, height: int, time: int
    ) -> None:
        super().__init__(session_id, time, question_number)
        self.type = "window_sized"
        self.width = width
        self.height = height

    def as_dictionary(self):
        return {
            "width": self.width,
            "height": self.height,
        } | super().as_dictionary()


def generate_keystroke_event(session_id: str, time) -> dict[str, any]:
    question_number = random.randint(1, 6)
    event_code_keys = list(event_which_to_event_code.keys())
    # there are 101 keys, the first 54 keys should appear more often than the rest
    probabilities = np.array([0.4] * 54 + [0.1] * 47)
    normalised_p = np.array(probabilities / probabilities.sum())
    key_code = int(rng.choice(event_code_keys, p=normalised_p))
    key_char = event_which_to_event_code[key_code]
    shift = random.choice([True, False])
    alt = random.choice([True, False])
    control = random.choice([True, False])
    meta = random.choice([True, False])
    unrelated_key = bool(rng.choice([True, False], p=[0.2, 0.8]))

    return (
        EventKeystroke(
            session_id,
            question_number,
            key_char,
            key_code,
            shift,
            alt,
            control,
            meta,
            unrelated_key,
            time,
        )
    ).as_dictionary()


def generate_mousemove_event(session_id: str, time) -> dict[str, any]:
    question_number = random.randint(1, 6)
    direction = random.choice(MOUSE_DIRECTIONS)
    window_height = random.randint(0, 1080)
    window_width = random.randint(0, 1920)
    x_position = random.randint(0, window_width)
    y_position = random.randint(0, window_height)

    return (
        EventMouseMove(
            session_id, question_number, direction, x_position, y_position, time
        )
    ).as_dictionary()


# generates a pair of mousedown and mouse up with a random interval
# they will always come in pairs
def generate_mouseclick_event(session_id: str, time: int) -> list[dict[str, any]]:
    question_number = random.randint(1, 6)
    window_height = random.randint(0, 1080)
    window_width = random.randint(0, 1920)
    x_position = random.randint(0, window_width)
    y_position = random.randint(0, window_height)
    button = random.choice(MOUSE_BUTTON)
    interval = random.randint(1, 5000)  # 1 to 5 seconds in milliseconds precision

    result = []

    result.append(
        EventMouseDown(
            session_id, question_number, button, x_position, y_position, time
        ).as_dictionary()
    )

    # mouseup occurs randomly after mousedown with a random interval
    result.append(
        EventMouseUp(
            session_id, question_number, button, x_position, y_position, time + interval
        ).as_dictionary()
    )

    return result


def generate_window_sized_event(session_id: str, time) -> dict[str, any]:
    question_number = random.randint(1, 6)
    width = random.randint(400, 1920)
    height = random.randint(200, 1080)

    return (
        EventWindowSized(session_id, question_number, width, height, time)
    ).as_dictionary()
