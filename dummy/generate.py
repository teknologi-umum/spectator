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
from pprint import pprint
import random
from model_event import (
    generate_event_keystroke,
    generate_event_mouseclick,
    generate_event_mousemove,
    generate_event_window_sized,
)
from model_session import (
    generate_event_after_exam_SAM_Submited,
    generate_event_deadline_passed,
    generate_event_before_exam_SAM_Submited,
    generate_event_exam_forfeited,
    generate_event_exam_ended,
    generate_event_exam_ide_reloaded,
    generate_event_exam_started,
    generate_event_locale_set,
    generate_event_personal_info_submitted,
    generate_event_session_started,
    generate_event_solution_accepted,
    generate_event_solution_rejected,
)

from model_user import generate_user
from utils import random_date


def write_into_file(filename: str, data):
    # return
    with open("generated/" + filename, "w") as f:
        data = json.dumps(data, sort_keys=True, indent=2, ensure_ascii=False)
        f.write(data)


def main():
    users: list[dict[str, any]] = []

    for _ in range(5):
        user = generate_user()
        print(user["session_id"])
        users.append(user)

    print("Generating user personal data...")
    write_into_file("user_personal.json", users)

    input_events: list[dict[str, any]] = []
    session_event: list[dict[str, any]] = []

    for user in users:
        # Get current session
        current_session = user["session_id"]
        current_input_events: list[dict[str, any]] = []
        current_session_events: list[dict[str, any]] = []
        # Generate 2 random dates that are close to each other
        date_start_int = random_date(
            datetime(2021, 6, 1, 0, 0, 0), datetime(2021, 12, 29, 23, 59, 59)
        )
        date_start = datetime.fromtimestamp(date_start_int)
        additional_duration = timedelta(minutes=random.randint(6, 21))
        date_ends = datetime.fromtimestamp(
            date_start_int + additional_duration.total_seconds()
        )

        INPUT_EVENTS = ["keystroke", "mousemove", "window_sized", "mouseclick"]
        SESSION_EVENTS = [
            "solution_accepted",
            "solution_rejected",
            "locale_set",
            "personal_info_submitted",
            "session_started",
            "deadline_passed",
            "exam_ended",
            "exam_forfeited",
            "exam_ide_reloaded",
            "exam_started",
            "session_started",
            "exam_before_sam_submited",
            "exam_after_sam_submitted",
        ]

        for _ in range(random.randint(1000, 10000)):
            # generate random input event.
            choice = random.choice(INPUT_EVENTS)
            if choice == "keystroke":
                event = generate_event_keystroke(current_session, date_start, date_ends)
            elif choice == "mousemove":
                event = generate_event_mousemove(current_session, date_start, date_ends)
            elif choice == "window_sized":
                event = generate_event_window_sized(
                    current_session, date_start, date_ends
                )
            elif choice == "mouseclick":
                event = generate_event_mouseclick(
                    current_session, date_start, date_ends
                )
            current_input_events.append(event)
        # Add the current events to the list of events
        input_events.extend(current_input_events)

        # a user always starts a session
        event = generate_event_session_started(
            current_session, date_start, date_ends
        )
        current_session_events.append(event)

        # first and foremost, personal is going to be submitted
        event = generate_event_personal_info_submitted(
            current_session, date_start, date_ends
        )
        current_session_events.append(event)

        # after that, they will submit a SAM test result before the exam
        event = generate_event_before_exam_SAM_Submited(
            current_session, date_start, date_ends
        )
        current_session_events.append(event)

        # and then they will start the exam
        event = generate_event_exam_started(
            current_session, date_start, date_ends
        )
        current_session_events.append(event)

        # both of these will appear randomly
        for i in range(0, 10):
            random_int = random.randint(0, 4)
            if random_int == 0:
                event = generate_event_exam_ide_reloaded(
                    current_session, date_start, date_ends
                )
                current_session_events.append(event)
            elif random_int == 1:
                event = generate_event_locale_set(
                    current_session, date_start, date_ends
                )
                current_session_events.append(event)

        # there are 3 ways of ending the test
        random_int = random.randint(0, 2)
        if random_int == 0:
            event = generate_event_deadline_passed(
                current_session, date_start, date_ends
            )
            current_session_events.append(event)
        elif random_int == 1:
            event = generate_event_exam_ended(
                current_session, date_start, date_ends
            )
            current_session_events.append(event)
        elif random_int == 2:
            event = generate_event_exam_forfeited(
                current_session, date_start, date_ends
            )
            current_session_events.append(event)

        # finally, they will submit a SAM test after the exam
        event = generate_event_after_exam_SAM_Submited(
            current_session, date_start, date_ends
        )
        current_session_events.append(event)

        session_event.extend(current_session_events)
    print(f"Generated { len(input_events) } input events. Writing into file.")
    write_into_file("input_events.json", input_events)

    print(f"Generated { len(session_event) } session events. Writing into file.")
    write_into_file("session_events.json", session_event)


if __name__ == "__main__":
    main()
