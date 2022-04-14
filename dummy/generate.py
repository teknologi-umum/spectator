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

from datetime import datetime
import json
import random
from model_event import (
    generate_keystroke_event,
    generate_mouseclick_event,
    generate_mousemove_event,
    generate_window_sized_event,
)
from model_session import (
    generate_after_exam_SAM_submited_event,
    generate_deadline_passed_event,
    generate_before_exam_SAM_submited_event,
    generate_exam_forfeited_event,
    generate_exam_ended_event,
    generate_exam_ide_reloaded_event,
    generate_exam_started_event,
    generate_locale_set_event,
    generate_personal_info_submitted_event,
    generate_session_started_event,
    generate_solution_accepted_event,
    generate_solution_rejected_event,
)

from model_user import generate_user
from utils import random_date

INPUT_EVENTS = ["keystroke", "mousemove", "window_sized", "mouseclick"]
MINUTES_TO_MILLIS = 60 * 1000


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
        time = random_date(
            datetime(2022, 1, 1, 0, 0, 0), datetime(2022, 1, 1, 20, 0, 0)
        )

        # a user always starts a session
        event = generate_session_started_event(
            current_session, time
        )
        current_session_events.append(event)

        # a user need to submit their personal info before starting the exam
        # let's assume every personal info requires at least 1 minute and at most 10 minutes
        time = time + random.randint(1, 10) * MINUTES_TO_MILLIS
        event = generate_personal_info_submitted_event(
            current_session, time
        )
        current_session_events.append(event)

        # after that, they will submit a SAM test result before the exam
        # let's assume every sam test requires at least 5 minutes and at most 15 minutes
        time = time + random.randint(5, 15) * MINUTES_TO_MILLIS
        event = generate_before_exam_SAM_submited_event(
            current_session, time
        )
        current_session_events.append(event)

        # and then they will start the exam
        # let's just say they need 30 to 60 seconds to start the exam
        time = time + random.randint(30, 60) * 1000  # s to ms conversion
        event = generate_exam_started_event(
            current_session, time
        )
        current_session_events.append(event)

        # both of these will appear randomly
        for i in range(0, 10):
            # this event could happen at any time through the exam parallel to other events
            # hence why we don't want to mutate the original `time` variable
            random_time = time + random.randint(1, 90) * MINUTES_TO_MILLIS
            random_int = random.randint(0, 4)
            if random_int == 0:
                event = generate_exam_ide_reloaded_event(
                    current_session, random_time
                )
                current_session_events.append(event)
            elif random_int == 1:
                event = generate_locale_set_event(
                    current_session, random_time
                )
                current_session_events.append(event)

        # these are the events that will occur in the exam
        for _ in range(random.randint(4000, 5000)):
            # these events will also happen parallel to each other so we shouldn't mutate the original timestamp
            # the events will be generated in the range of 1 to 90 minutes
            # and the delta will be randomised between 1ms to (9 * 60 * 1000)ms
            random_time = time + random.randint(1, 90 * MINUTES_TO_MILLIS)
            # generate random input event.
            choice = random.choice(INPUT_EVENTS)
            if choice == "keystroke":
                event = generate_keystroke_event(current_session, random_time)
                current_input_events.append(event)
            elif choice == "mousemove":
                event = generate_mousemove_event(current_session, random_time)
                current_input_events.append(event)
            elif choice == "window_sized":
                event = generate_window_sized_event(current_session, random_time)
                current_input_events.append(event)
            elif choice == "mouseclick":
                event = generate_mouseclick_event(current_session, random_time)
                current_input_events.append(event)

        # there will be 6 questions
        for _ in range(6):
            # the question will be submitted in the range of 1 to 80 minutes
            # these events will also happen parallel to the other ones so don't mutate the original timestamp
            random_time = time + random.randint(1, 80) * MINUTES_TO_MILLIS
            random_int = random.randint(0, 1)
            if random_int == 0:
                event = generate_solution_accepted_event(
                    current_session, random_time
                )
                current_input_events.append(event)
            elif random_int == 1:
                event = generate_solution_rejected_event(
                    current_session, random_time
                )
                current_input_events.append(event)

        # there are 3 ways of ending the test
        random_int = random.randint(0, 2)
        if random_int == 0:
            # deadline passed means the time limit has ended
            time = time + 90 * MINUTES_TO_MILLIS
            event = generate_deadline_passed_event(
                current_session, time
            )
            current_session_events.append(event)
        elif random_int == 1:
            # exam ended means the user pressed the 'finish' button
            # let's say the minimum time required to finish the test is 30 minutes and the maximum is 89 minutes
            # we don't want it to be 90 minutes otherwise it's going to be the same as deadline passed
            time = time + random.randint(30, 89) * MINUTES_TO_MILLIS
            event = generate_exam_ended_event(
                current_session, time
            )
            current_session_events.append(event)
        elif random_int == 2:
            # exam forfeited means the user pressed the 'surrender' button
            # the logic is going to be the same as exam ended, just with a different range
            time = time + random.randint(10, 80) * MINUTES_TO_MILLIS
            event = generate_exam_forfeited_event(
                current_session, time
            )
            current_session_events.append(event)

        # finally, they will submit a SAM test after the exam
        # let's assume every sam test requires at least 5 minutes and at most 15 minutes
        time = time + random.randint(5, 15) * MINUTES_TO_MILLIS
        event = generate_after_exam_SAM_submited_event(
            current_session, time
        )
        current_session_events.append(event)

        input_events.extend(current_input_events)
        session_event.extend(current_session_events)

    print(f"Generated { len(input_events) } input events. Writing into file.")
    write_into_file("input_events.json", input_events)

    print(f"Generated { len(session_event) } session events. Writing into file.")
    write_into_file("session_events.json", session_event)


if __name__ == "__main__":
    main()
