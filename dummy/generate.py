"""
This file contains the code for generating the data.
The data that being generated is about:
    - User's personal info (including their session ID)
      which from the latest conversation is in the format
      of UUID.
    - Keystroke events
    - Mouse move events
    - Mouse click events

As the generation process is relatively quick, there's
no much need for verbosity in this one.
"""

from datetime import datetime, timedelta
import json
import random
from model_event import generate_event_keystroke, generate_event_mouseclick, generate_event_mousemove, generate_event_window_sized
from model_user import generate_user
from utils import random_date

def write_into_file(filename: str, data):
    with open("generated/" + filename, "w") as f:
        data = json.dumps(data, sort_keys=True, indent=2, ensure_ascii=False)
        f.write(data)

def main():
    users: list[dict[str, any]] = []

    # begini bukan
    number_of_data = 5
    date_range = 4

    for _ in range(5):
        user = generate_user()
        users.append(user)

    print("Generating user personal data...")
    write_into_file("user_personal.json", users)

    events: list[dict[str, any]] = []

    for i in range(len(users)):
        # Get current session
        current_session = users[i]["session_id"]
        current_events: list[dict[str, any]] = []

        # Generate 2 random dates that are close to each other
        date_start_int: int = random_date(datetime(2021, 6, 1, 0, 0, 0), datetime(2021, 12, 29, 23, 59, 59))
        date_start: datetime = datetime.fromtimestamp(date_start_int)
        additional_duration: timedelta = timedelta(minutes=random.randint(4, (4 + date_range)))
        date_ends: datetime = datetime.fromtimestamp(date_start_int + additional_duration.total_seconds())

        for _ in range(random.randint(420 * number_of_data, 666 * number_of_data)):
            # Generate random number between 1 to 3
            rand = random.randint(1, 4)
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
            elif rand == 3:
                # Generate a window resize event
                event = generate_event_window_sized(current_session, date_start, date_ends)
                current_events.append(event)
                continue
            else:
                # Generate a mouse click event
                event = generate_event_mouseclick(current_session, date_start, date_ends)
                current_events.append(event)
                continue

        # Add the current events to the list of events
        events.extend(current_events)

    print("Generated {} events. Writing into file.".format(len(events)))
    write_into_file("events.json", events)

if __name__ == "__main__":
    main()
