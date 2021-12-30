import time
from random import randrange
from datetime import datetime, timedelta

def random_date(start: datetime, end: datetime) -> int:
    """
    This function will return a random datetime between two datetime
    objects.
    """
    delta: timedelta = end - start
    int_delta: int = (delta.days * 24 * 60 * 60) + delta.seconds
    random_second: int = randrange(int_delta)
    time_addition = start + timedelta(seconds=random_second)
    return int(time.mktime(time_addition.timetuple()))
