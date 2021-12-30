# Code documentation soon

from datetime import datetime, timedelta
import json
import random
from model_event import generate_event_keystroke, generate_event_mouseclick, generate_event_mousemove
from model_user import generate_user
from utils import random_date

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
