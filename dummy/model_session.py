from pprint import pprint
import random
from datetime import datetime,timedelta
from time import time
from uuid import uuid4
from utils import random_date

#const

_Language = {
    "Undefined":0, 
    "C":1, 
    "CPP":2, 
    "PHP":3, 
    "Javascript":4, 
    "Java":5, 
    "Python":6
}
_Locale = {
    "EN":0,
    "ID":1
}

class SessionEventBase:
    session_id: str
    type: str
    _time: int

    def __init__(self, session_id: str, time:int) -> None:
        self.session_id = session_id
        self._time = time

class EventSolution(SessionEventBase):
    question_number: int
    language: int
    solution: str
    scratchpad: str
    serialized_test_result: str

    def __init__( self, session_id: str, time: int, question_number:int , language: str , solution: str, scratchpad:str , serialized_test_result: str) -> None:
        super().__init__(session_id, time)
        self.question_number = question_number
        self.language = language
        self.solution = solution
        self.scratchpad = scratchpad
        self.serialized_test_result = serialized_test_result

    def asdict(self):
        return {
            "type": self.type ,
            "session_id": self.session_id,
            "time": self._time,
            "question_number": self.question_number,
            "language": self.language,
            "solution": self.solution,
            "scratchpad": self.scratchpad,
            "serialized_test_result": self.serialized_test_result,
        }

class EventSolutionAccepted(EventSolution):

    def __init__( self, session_id: str, time: int, question_number:int , language: str , solution: str, scratchpad:str , serialized_test_result: str) -> None:
        super().__init__(  session_id, question_number, time, language, solution, scratchpad, serialized_test_result)
        self.type = "solution_accepted"
    
class EventSolutionRejected(EventSolution):
    def __init__( self, session_id: str, time: int, question_number:int , language: int , solution: str, scratchpad:str , serialized_test_result: str) -> None:
        super().__init__( session_id, question_number, time, language, solution, scratchpad, serialized_test_result)
        self.type = "solution_rejected"

class EventLocaleSet(SessionEventBase):
    locale: int

    def __init__( self, session_id: str, time: int, locale:str ) -> None:
        super().__init__( session_id, time)
        self.type = "exam_started"
        self.locale = _Locale[locale]

    def asdict(self):
        return {
            "type" : self.type ,
            "session_id" : self.session_id,
            "time" : self._time,
            "locale" :self.locale
        }

class EventPersonalInfoSubmited(SessionEventBase):
    student_number:str
    years_of_experience: int
    hours_of_practice: int
    familiar_languages: str

    def __init__( self, session_id: str, time: int, student_number: int, years_of_experience: int, hours_of_practice: int, familiar_languages:str ) -> None:
        super().__init__(session_id, time)
        self.type = "personal_info_submitted"
        self.student_number = student_number
        self.years_of_experience = years_of_experience
        self.hours_of_practice = hours_of_practice
        self.familiar_languages = familiar_languages
    def asdict(self):
        return {
            "type" : self.type ,
            "session_id" : self.session_id,
            "time" : self._time,
            "student_number": self.student_number,
            "years_of_experience": self.years_of_experience,
            "hours_of_practice": self.hours_of_practice,
            "familiar_languages": self.familiar_languages,
        }

def _generate_event_solution ( session_id,  date_start: datetime, date_ends: datetime, ) -> dict[str, any]:
    question_number = random.randint(1, 6)
    language = _Language[ random.choice(["Undefined", "C", "CPP", "PHP", "Javascript", "Java", "Python"]) ]
    solution = "" # add gibbrish ??
    scratchpad = "" # add gibbrish ??
    serialized_test_result = "" # add gibbrish ??
    time = random_date(date_start, date_ends)
    return { 
        "session_id" :session_id,
        "time": time, 
        "question_number": question_number, 
        "language": language, 
        "solution": solution, 
        "scratchpad": scratchpad, 
        "serialized_test_result": serialized_test_result 
    }
    
def generate_event_solution_accepted( session_id,  date_start: datetime, date_ends: datetime ) -> dict[str, any]:
    """Generate an EventSolutionAccepted class with random values.
    The "language" key may only be either "Undefined", "C", "CPP", "PHP", "Javascript", "Java", "Python"

    Args:
        date_start (datetime): [description]
        date_ends (datetime): [description]

    Returns:
        dict[str, any]: [description]
    """
    data= _generate_event_solution( session_id,  date_start, date_ends )
    return ( EventSolutionAccepted ( **data ).asdict() )

def generate_event_solution_rejected( session_id,  date_start: datetime, date_ends: datetime ) -> dict[str, any]:
    """Generate an EventSolutionRejected class with random values.
    The "language" key may only be either "Undefined", "C", "CPP", "PHP", "Javascript", "Java", "Python"

    Args:
        date_start (datetime): [description]
        date_ends (datetime): [description]

    Returns:
        dict[str, any]: [description]
    """
    data= _generate_event_solution( session_id,  date_start, date_ends )
    return ( EventSolutionRejected(**data).asdict() ) 

def generate_event_locale_set( session_id,  date_start: datetime, date_ends: datetime ) -> dict[str, any]:
    time = random_date(date_start, date_ends)
    locale = random.choice(["ID","EN"])
    return (EventLocaleSet(session_id, time, locale).asdict())

def generate_event_personal_info_submitted( session_id,  date_start: datetime, date_ends: datetime ) -> dict[str, any]:
    time = random_date(date_start, date_ends)
    student_number = random.randint(999999999,9999999999)
    years_of_experience = random.randint(1,10)
    hours_of_practice = random.randint(0,10000)
    familiar_languages = ""
    return (EventPersonalInfoSubmmited(session_id,time, student_number, years_of_experience,hours_of_practice,familiar_languages).asdict())
    
# checker
def fake_basic_info_generate():
    date_start_int: int = random_date(datetime(2021, 6, 1, 0, 0, 0), datetime(2021, 12, 29, 23, 59, 59))
    date_start: datetime = datetime.fromtimestamp(date_start_int)
    additional_duration: timedelta = timedelta(minutes=random.randint(6, 21))
    date_ends: datetime = datetime.fromtimestamp(date_start_int + additional_duration.total_seconds())
    return {
        "session_id": str(uuid4()),
        "date_start" :date_start,
        "date_ends" :date_ends
    }

def main():
    pprint(generate_event_solution_accepted(**fake_basic_info_generate()))
    pprint(generate_event_solution_rejected(**fake_basic_info_generate()))
    pprint(generate_event_locale_set(**fake_basic_info_generate()))
    pprint(generate_event_personal_info_submitted(**fake_basic_info_generate()))
    
if __name__ == "__main__":
    main()