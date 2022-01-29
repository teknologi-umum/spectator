from pprint import pprint
import random
from datetime import datetime,timedelta
from uuid import uuid4
from utils import random_date

#const

_Language = [ "Undefined", "C", "CPP", "PHP", "Javascript", "Java", "Python" ] 
_Locale = [ "EN", "ID" ]

class SessionEventBase:
    session_id: str
    type: str
    _time: int

    def __init__(self, session_id: str, time:int) -> None:
        self.session_id = session_id
        self._time = time
    def as_dictionary(self)-> dict[str, any]:
        return {
            "type": self.type,
            "session_id": self.session_id,
            "time": self._time
        }
class EventSolution(SessionEventBase):
    question_number: int
    language: int
    solution: str
    scratchpad: str
    serialized_test_results: str

    def __init__( self, session_id: str, time: int, question_number:int , language: str , solution: str, scratchpad:str , serialized_test_results: str) -> None:
        super().__init__(session_id, time)
        self.question_number = question_number
        self.language = language
        self.solution = solution
        self.scratchpad = scratchpad
        self.serialized_test_results = serialized_test_results

    def as_dictionary(self):
        return {
            "question_number": self.question_number,
            "language": self.language,
            "solution": self.solution,
            "scratchpad": self.scratchpad,
            "serialized_test_results": self.serialized_test_results,
        } | super().as_dictionary()
class EventSolutionAccepted(EventSolution):

    def __init__( self, session_id: str, time: int, question_number:int , language: str , solution: str, scratchpad:str , serialized_test_results: str) -> None:
        super().__init__(  session_id, question_number, time, language, solution, scratchpad, serialized_test_results)
        self.type = "solution_accepted"
class EventSolutionRejected(EventSolution):
    def __init__( self, session_id: str, time: int, question_number:int , language: int , solution: str, scratchpad:str , serialized_test_results: str) -> None:
        super().__init__( session_id, question_number, time, language, solution, scratchpad, serialized_test_results)
        self.type = "solution_rejected"
class EventLocaleSet(SessionEventBase):
    locale: str

    def __init__( self, session_id: str, time: int, locale:int ) -> None:
        super().__init__( session_id, time)
        self.type = "exam_started"
        self.locale = _Locale[locale]

    def as_dictionary(self):
        return {
            "locale" :self.locale
        } | super().as_dictionary()
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
    def as_dictionary(self):
        return {
            "student_number": self.student_number,
            "years_of_experience": self.years_of_experience,
            "hours_of_practice": self.hours_of_practice,
            "familiar_languages": self.familiar_languages,
        } | super().as_dictionary()
class EventSessionStarted(SessionEventBase):
    locale: str

    def __init__( self, session_id: str, time: int, locale:int ) -> None:
        super().__init__( session_id, time)
        self.type = "exam_started"
        self.locale = _Locale[locale]

    def as_dictionary(self):
        return {
            "locale" :self.locale
        } | super().as_dictionary()
class EventDeadlinePassed(SessionEventBase):
    def __init__(self, session_id: str, time: int) -> None:
        super().__init__(session_id, time)
        self.type = "deadline_passed"
class EventExamEnded(SessionEventBase):
    def __init__(self, session_id: str, time: int) -> None:
        super().__init__(session_id, time)
        self.type = "exam_ended"
class EventExamForfeited(SessionEventBase):
    def __init__(self, session_id: str, time: int) -> None:
        super().__init__(session_id, time)
        self.type = "exam_forfeited"
class EventExamIDEReloaded(SessionEventBase):
    def __init__(self, session_id: str, time: int) -> None:
        super().__init__(session_id, time)
        self.type = "exam_ide_reloaded"
class EventExamStarted(SessionEventBase):
    question_numbers=list[int]
    deadline=str
    def __init__(self, session_id: str, time: int, question_numbers: list[int], deadline) -> None:
        super().__init__(session_id, time)
        self.type = "exam_started"
        self.question_numbers=question_numbers
        self.deadline=deadline
    def as_dictionary(self):
        return{
            "question_numbers": self.question_numbers,
            "deadline": self.deadline
        } | super().as_dictionary()
class EventExamSAMSubmitted(SessionEventBase):
    aroused_level:int
    pleased_level:int

    def __init__(self, session_id: str, time: int, aroused_level: int, pleased_level: int) -> None:
        super().__init__(session_id, time)
        self.aroused_level=aroused_level
        self.pleased_level=pleased_level
    def as_dictionary(self):
        return{
            "aroused_level": self.aroused_level,
            "pleased_level": self.pleased_level
        } | super().as_dictionary()
class EventAfterExamSAMSubmitted(EventExamSAMSubmitted):
    def __init__(self, session_id: str, time: int, aroused_level: int, pleased_level: int) -> None:
        super().__init__(session_id, time, aroused_level, pleased_level)
        self.type="after_exam_sam_submitted"
class EventBeforeExamSAMSubmitted(EventExamSAMSubmitted):
    def __init__(self, session_id: str, time: int, aroused_level: int, pleased_level: int) -> None:
        super().__init__(session_id, time, aroused_level, pleased_level)
        self.type="before_exam_sam_submitted"

def _generate_event_solution ( session_id,  date_start: datetime, date_ends: datetime ) -> dict[str, any]:
    question_number = random.randint(1, 6)
    language = _Language[ random.choice(["Undefined", "C", "CPP", "PHP", "Javascript", "Java", "Python"]) ]
    solution = generate_gibberish_code()
    scratchpad = generate_gibberish_code()
    serialized_test_results = "" # add gibbrish ??
    time = random_date(date_start, date_ends)
    return { 
        "session_id" :session_id,
        "time": time, 
        "question_number": question_number, 
        "language": language, 
        "solution": solution, 
        "scratchpad": scratchpad, 
        "serialized_test_results": serialized_test_results 
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
    return ( EventSolutionAccepted ( **data ).as_dictionary() )

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
    return ( EventSolutionRejected(**data).as_dictionary() ) 

def generate_event_locale_set( session_id,  date_start: datetime, date_ends: datetime ) -> dict[str, any]:
    time = random_date(date_start, date_ends)
    locale = random.choice(["ID", "EN"])
    return (EventLocaleSet(session_id, time, locale).as_dictionary())

def generate_event_session_started( session_id,  date_start: datetime, date_ends: datetime ) -> dict[str, any]:
    time = random_date(date_start, date_ends)
    locale = random.choice(["ID", "EN"])
    return (EventSessionStarted(session_id, time, locale).as_dictionary())

def generate_event_personal_info_submitted( session_id,  date_start: datetime, date_ends: datetime ) -> dict[str, any]:
    time = random_date(date_start, date_ends)
    student_number = random.randint(999999999,9999999999)
    years_of_experience = random.randint(1,10)
    hours_of_practice = random.randint(0,10000)
    familiar_languages = ""
    return (EventPersonalInfoSubmited(session_id,time, student_number, years_of_experience,hours_of_practice,familiar_languages).as_dictionary())

def generate_event_session_started( session_id,  date_start: datetime, date_ends: datetime ) -> dict[str, any]:
    time = random_date(date_start, date_ends)
    locale = random.choice(["ID","EN"])
    return_object = EventSessionStarted(session_id, time, locale)
    return return_object.as_dictionary()

def generate_event_deadline_passed( session_id,  date_start: datetime, date_ends: datetime ) -> dict[str, any]:
    time = random_date(date_start, date_ends)
    return_object = EventDeadlinePassed(session_id, time)
    return (return_object.as_dictionary())

def generate_event_exam_ended( session_id,  date_start: datetime, date_ends: datetime ) -> dict[str, any]:
    time = random_date(date_start, date_ends)
    return_object = EventExamEnded(session_id, time)
    return (return_object.as_dictionary())

def generate_event_exam_forfeited( session_id,  date_start: datetime, date_ends: datetime ) -> dict[str, any]:
    time = random_date(date_start, date_ends)
    return_object = EventExamForfeited(session_id, time)
    return (return_object.as_dictionary())

def generate_event_exam_ide_reloaded( session_id,  date_start: datetime, date_ends: datetime ) -> dict[str, any]:
    time = random_date(date_start, date_ends)
    return_object = EventExamIDEReloaded(session_id, time)
    return (return_object.as_dictionary())

def generate_event_exam_started( session_id,  date_start: datetime, date_ends: datetime ) -> dict[str, any]:
    time = random_date(date_start, date_ends)
    question_numbers = [1, 2, 3, 4, 5, 6]
    deadline = random_date(datetime.fromtimestamp(time / 1e3), date_ends)
    return_object = EventExamStarted(session_id, time, question_numbers, deadline)
    return (return_object.as_dictionary())

def generate_event_before_exam_SAM_Submited( session_id,  date_start: datetime, date_ends: datetime ) -> dict[str, any]:
    aroused_level = random.randint(0,5)
    pleased_level = random.randint(0,5)
    time = random_date(date_start, date_ends)
    return_object = EventBeforeExamSAMSubmitted(session_id, time, aroused_level, pleased_level)
    return (return_object.as_dictionary())

def generate_event_after_exam_SAM_Submited( session_id,  date_start: datetime, date_ends: datetime ) -> dict[str, any]:
    aroused_level = random.randint(0,5)
    pleased_level = random.randint(0,5)
    time = random_date(date_start, date_ends)
    return_object = EventAfterExamSAMSubmitted(session_id, time, aroused_level, pleased_level)
    return (return_object.as_dictionary())

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

def generate_gibberish_code():
    return random.choice(["""
var
 radians = Math.PI / 4, // Pi / 4 is 45 degrees. All answers should be the same.
 degrees = 45.0,
 sine = Math.sin(radians),
 cosine = Math.cos(radians),
 tangent = Math.tan(radians),
 arcsin = Math.asin(sine),
 arccos = Math.acos(cosine),
 arctan = Math.atan(tangent);
 
// sine
window.alert(sine + " " + Math.sin(degrees * Math.PI / 180));
// cosine
window.alert(cosine + " " + Math.cos(degrees * Math.PI / 180));
// tangent
window.alert(tangent + " " + Math.tan(degrees * Math.PI / 180));
// arcsine
window.alert(arcsin + " " + (arcsin * 180 / Math.PI));
// arccosine
window.alert(arccos + " " + (arccos * 180 / Math.PI));
// arctangent
window.alert(arctan + " " + (arctan * 180 / Math.PI));
            ""","""
package main
 
import "fmt"
 
var v = []float32{1, 2, .5}
 
func main() {
    var sum float32
    for _, x := range v {
        sum += x * x
    }
    fmt.Println(sum)
}
            ""","""
require 'ipaddr'
 
 
TESTCASES = ["127.0.0.1",                "127.0.0.1:80",
                "::1",                      "[::1]:80",
                "2605:2700:0:3::4713:93e3", "[2605:2700:0:3::4713:93e3]:80"]                            
 
output = [%w(String Address Port Family Hex),
          %w(------ ------- ---- ------ ---)]
 
def output_table(rows)
  widths = []
  rows.each {|row| row.each_with_index {|col, i| widths[i] = [widths[i].to_i, col.to_s.length].max }}
  format = widths.map {|size| "%#{size}s"}.join("\t")
  rows.each {|row| puts format % row}
end
 
TESTCASES.each do |str|
  case str  # handle port; IPAddr does not.
  when /\A\[(?<address> .* )\]:(?<port> \d+ )\z/x      # string like "[::1]:80"
    address, port = $~[:address], $~[:port]
  when /\A(?<address> [^:]+ ):(?<port> \d+ )\z/x       # string like "127.0.0.1:80"
    address, port = $~[:address], $~[:port]
  else                                                 # string with no port number
    address, port = str, nil
  end
 
  ip_addr = IPAddr.new(address) 
  family = "IPv4" if ip_addr.ipv4?
  family = "IPv6" if ip_addr.ipv6?
 
  output << [str, ip_addr.to_s, port.to_s, family, ip_addr.to_i.to_s(16)]
end
 
output_table(output)
            ""","""
package main
 
import (
    "fmt"
    "math/bits"
)
 
func main() {
    fmt.Println("Pop counts, powers of 3:")
    n := uint64(1) // 3^0
    for i := 0; i < 30; i++ {
        fmt.Printf("%d ", bits.OnesCount64(n))
        n *= 3
    }
    fmt.Println()
    fmt.Println("Evil numbers:")
    var od [30]uint64
    var ne, no int
    for n = 0; ne+no < 60; n++ {
        if bits.OnesCount64(n)&1 == 0 {
            if ne < 30 {
                fmt.Printf("%d ", n)
                ne++
            }
        } else {
            if no < 30 {
                od[no] = n
                no++
            }
        }
    }
    fmt.Println()
    fmt.Println("Odious numbers:")
    for _, n := range od {
        fmt.Printf("%d ", n)
    }
    fmt.Println()
}
            ""","""
class Latin {
  constructor(size = 3) {
    this.size = size;
    this.mst = [...Array(this.size)].map((v, i) => i + 1);
    this.square = Array(this.size).fill(0).map(() => Array(this.size).fill(0));
 
    if (this.create(0, 0)) {
      console.table(this.square);
    }
  }
 
  create(c, r) {
    const d = [...this.mst];
    let s;
    while (true) {
      do {
        s = d.splice(Math.floor(Math.random() * d.length), 1)[0];
        if (!s) return false;
      } while (this.check(s, c, r));
 
      this.square[c][r] = s;
      if (++c >= this.size) {
        c = 0;
        if (++r >= this.size) {
          return true;
        }
      }
      if (this.create(c, r)) return true;
      if (--c < 0) {
        c = this.size - 1;
        if (--r < 0) {
          return false;
        }
      }
    }
  }
 
  check(d, c, r) {
    for (let a = 0; a < this.size; a++) {
      if (c - a > -1) {
        if (this.square[c - a][r] === d)
          return true;
      }
      if (r - a > -1) {
        if (this.square[c][r - a] === d)
          return true;
      }
    }
    return false;
  }
}
new Latin(5);
            """ ])

def checker():
    pprint(generate_event_solution_accepted(**fake_basic_info_generate()))
    pprint(generate_event_solution_rejected(**fake_basic_info_generate()))
    pprint(generate_event_locale_set(**fake_basic_info_generate()))
    pprint(generate_event_personal_info_submitted(**fake_basic_info_generate()))
    pprint(generate_event_session_started(**fake_basic_info_generate()))
    pprint(generate_event_deadline_passed(**fake_basic_info_generate()))
    pprint(generate_event_exam_ended(**fake_basic_info_generate()))
    pprint(generate_event_exam_forfeited(**fake_basic_info_generate()))
    pprint(generate_event_exam_ide_reloaded(**fake_basic_info_generate()))
    pprint(generate_event_exam_started(**fake_basic_info_generate()))
    pprint(generate_event_after_exam_SAM_Submited(**fake_basic_info_generate()))
    pprint(generate_event_before_exam_SAM_Submited(**fake_basic_info_generate()))
    
if __name__ == "__main__":
    checker()