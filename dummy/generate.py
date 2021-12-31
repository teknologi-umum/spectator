# Code documentation soon

from datetime import datetime, timedelta
import json
import random
from model_event import generate_event_keystroke, generate_event_mouseclick, generate_event_mousemove
from model_user import generate_user
from utils import random_date

<<<<<<< HEAD
class User():
    type: str
    session_id: str
    student_number: str
    hours_of_practice: int
    years_of_experience: int
    familiar_language: str
    
    def __init__(self, type: str, session_id: str, student_number: str,
                 hours_of_practice: int, years_of_experience: int,
                 familiar_language: str):
        self.type = type
        self.session_id = session_id
        self.student_number = student_number
        self.hours_of_practice = hours_of_practice
        self.years_of_experience = years_of_experience
        self.familiar_language = familiar_language
    
    def asdict(self):
        return {
            "type": self.type,
            "session_id": self.session_id,
            "student_number": self.student_number,
            "hours_of_practice": self.hours_of_practice,
            "years_of_experience": self.years_of_experience,
            "familiar_language": self.familiar_language
        }

def generate_user() -> dict[str, any]:
    languages = ["python", "javascript", "php", "", "c", "c++", "c#", "pascal"]
    letters = string.ascii_lowercase
    numbers = string.digits
    user = User(
        "personal_info", 
        "".join(random.choice(letters) for _ in range(6)),
        "".join(random.choice(numbers) for _ in range(8)),
        random.randint(0, 5),
        random.randint(0, 24*7),
        " ".join(random.sample(languages, k=random.randint(1,3)))
    )
    return user.asdict()

class KeystrokeEvent():
    session_id: str
    type: str
    question_number: int
    key_char: str
    key_code: str
    shift: bool
    alt: bool
    control: bool
    unrelated_key: bool
    meta: bool

    def __init__(self,type: str,question_number: int, key_char: str, key_code: str, shift: bool,alt: bool,control: bool,unrelated_key: bool,meta: bool):
        self.type = type
        self.session_id = session_id
        self.question_number = question_number
        self.key_char = key_char
        self.key_code = key_code
        self.shift = shift
        self.alt = alt 
        self.control = control
        self.unrelated_key = unrelated_key
        self.meta = meta 

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
            "unrelated_key": self.unrelated_key,
            "meta": self.meta
        }

def generate_keystroke() -> dict[str,any]:
    letters = string.ascii_lowercase
    numbers = string.digits
    keystroke = KeystrokeEvent(
        "coding_event_keystroke",
        "".join(random.choice(letters) for _ in range(6)),
        random.choice(numbers),
        "".join(random.choice(letters) for _ in range(1)),
        "".join(random.choice(numbers) for _ in range(2)),
        random.choice([True,False]),
        random.choice([True,False]),
        random.choice([True,False]),
        random.choice([True,False])
    )
    return keystroke.asdict()

=======
>>>>>>> 4f78690afa6bb2224c37fefd047537f90e0037fb
def write_into_file(filename: str, data):
    with open("generated/" + filename, "w") as f:
        data = json.dumps(data, sort_keys=True, indent=2, ensure_ascii=False)
        f.write(data)

def main():
    users: list[dict[str, any]] = []

    for _ in range(5):
        user = generate_user()
        users.append(user)

    write_into_file("user_personal.json", users)

    events: list[dict[str, any]] = []

    for i in range(len(users)):
        # Get current session
        current_session = users[i]["session_id"]
        current_events: list[dict[str, any]] = []

        # Generate 2 random dates that are close to each other
        date_start_int: int = random_date(datetime(2021, 6, 1, 0, 0, 0), datetime(2021, 12, 29, 23, 59, 59))
        date_start: datetime = datetime.fromtimestamp(date_start_int)
        additional_duration: timedelta = timedelta(minutes=random.randint(20, 90))
        date_ends: datetime = datetime.fromtimestamp(date_start_int + additional_duration.total_seconds())

        for _ in range(random.randint(200, 500)):
            # Generate random number between 1 to 3
            rand = random.randint(1, 3)
            if rand == 1:
                # Generate a keystroke event
                event = generate_event_keystroke(current_session, date_start, date_ends)
                current_events.append(event)
                continue
            elif rand == 2:
                # Generate a mouse move event
                event = generate_event_mousemove(current_session, date_start, date_ends)
                current_events.append(event)
                continue
            else:
                # Generate a mouse click event
                event = generate_event_mouseclick(current_session, date_start, date_ends)
                current_events.append(event)
                continue

        # Add the current events to the list of events
        events.extend(current_events)

    write_into_file("events.json", events)

if __name__ == "__main__":
    main()
