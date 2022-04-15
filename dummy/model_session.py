import random
from datetime import datetime, timedelta
from uuid import uuid4
from utils import random_date

LANGUAGE = {
    "Undefined": 0,
    "C": 1,
    "CPP": 2,
    "PHP": 3,
    "Javascript": 4,
    "Java": 5,
    "Python": 6,
}

LOCALE = {"EN": 0, "ID": 1}


class SessionEventBase:
    session_id: str
    type: str
    _time: int

    def __init__(self, session_id: str, time: int) -> None:
        self.session_id = session_id
        self._time = time

    def as_dictionary(self) -> dict[str, any]:
        return {
            "type": self.type,
            "session_id": self.session_id,
            "time": self._time
        }


class SolutionEventBase(SessionEventBase):
    question_number: int
    language: int
    solution: str
    scratchpad: str
    serialized_test_results: str

    def __init__(
        self,
        session_id: str,
        time: int,
        question_number: int,
        language: str,
        solution: str,
        scratchpad: str,
        serialized_test_results: str,
    ) -> None:
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


class SolutionAcceptedEvent(SolutionEventBase):
    def __init__(
        self,
        session_id: str,
        time: int,
        question_number: int,
        language: str,
        solution: str,
        scratchpad: str,
        serialized_test_results: str,
    ) -> None:
        super().__init__(
            session_id,
            question_number,
            time,
            language,
            solution,
            scratchpad,
            serialized_test_results,
        )
        self.type = "solution_accepted"


class SolutionRejectedEvent(SolutionEventBase):
    def __init__(
        self,
        session_id: str,
        time: int,
        question_number: int,
        language: int,
        solution: str,
        scratchpad: str,
        serialized_test_results: str,
    ) -> None:
        super().__init__(
            session_id,
            question_number,
            time,
            language,
            solution,
            scratchpad,
            serialized_test_results,
        )
        self.type = "solution_rejected"


class LocaleSetEvent(SessionEventBase):
    locale: str

    def __init__(self, session_id: str, time: int, locale: int) -> None:
        super().__init__(session_id, time)
        self.type = "locale_set"
        self.locale = LOCALE[locale]

    def as_dictionary(self):
        return {
            "locale": self.locale
        } | super().as_dictionary()


class PersonalInfoSubmitedEvent(SessionEventBase):
    student_number: str
    years_of_experience: int
    hours_of_practice: int
    familiar_languages: str

    def __init__(
        self,
        session_id: str,
        time: int,
        student_number: int,
        years_of_experience: int,
        hours_of_practice: int,
        familiar_languages: str,
    ) -> None:
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


class SessionStartedEvent(SessionEventBase):
    locale: str

    def __init__(self, session_id: str, time: int, locale: int) -> None:
        super().__init__(session_id, time)
        self.type = "session_started"
        self.locale = LOCALE[locale]

    def as_dictionary(self):
        return {
            "locale": self.locale
        } | super().as_dictionary()


class DeadlinePassedEvent(SessionEventBase):
    def __init__(self, session_id: str, time: int) -> None:
        super().__init__(session_id, time)
        self.type = "deadline_passed"


class ExamEndedEvent(SessionEventBase):
    def __init__(self, session_id: str, time: int) -> None:
        super().__init__(session_id, time)
        self.type = "exam_ended"


class ExamForfeitedEvent(SessionEventBase):
    def __init__(self, session_id: str, time: int) -> None:
        super().__init__(session_id, time)
        self.type = "exam_forfeited"


class ExamIDEReloadedEvent(SessionEventBase):
    def __init__(self, session_id: str, time: int) -> None:
        super().__init__(session_id, time)
        self.type = "exam_ide_reloaded"


class ExamStartedEvent(SessionEventBase):
    question_numbers = list[int]
    deadline = str

    def __init__(
        self, session_id: str, time: int, question_numbers: list[int], deadline: int
    ) -> None:
        super().__init__(session_id, time)
        self.type = "exam_started"
        self.question_numbers = question_numbers
        self.deadline = deadline

    def as_dictionary(self):
        return {
            "question_numbers": self.question_numbers,
            "deadline": self.deadline,
        } | super().as_dictionary()


class SAMSubmittedEventBase(SessionEventBase):
    aroused_level: int
    pleased_level: int

    def __init__(
        self, session_id: str, time: int, aroused_level: int, pleased_level: int
    ) -> None:
        super().__init__(session_id, time)
        self.aroused_level = aroused_level
        self.pleased_level = pleased_level

    def as_dictionary(self):
        return {
            "aroused_level": self.aroused_level,
            "pleased_level": self.pleased_level,
        } | super().as_dictionary()


class AfterExamSAMSubmittedEvent(SAMSubmittedEventBase):
    def __init__(
        self, session_id: str, time: int, aroused_level: int, pleased_level: int
    ) -> None:
        super().__init__(session_id, time, aroused_level, pleased_level)
        self.type = "after_exam_sam_submitted"


class BeforeExamSAMSubmittedEvent(SAMSubmittedEventBase):
    def __init__(
        self, session_id: str, time: int, aroused_level: int, pleased_level: int
    ) -> None:
        super().__init__(session_id, time, aroused_level, pleased_level)
        self.type = "before_exam_sam_submitted"


def _generate_solution_event(
    session_id, time
) -> dict[str, any]:
    question_number = random.randint(1, 6)
    language = random.choice(list(LANGUAGE.keys()))
    solution = generate_gibberish_code()
    scratchpad = generate_gibberish_code()
    serialized_test_results = ""  # add gibbrish ??
    return {
        "session_id": session_id,
        "time": time,
        "question_number": question_number,
        "language": language,
        "solution": solution,
        "scratchpad": scratchpad,
        "serialized_test_results": serialized_test_results,
    }


def generate_solution_accepted_event(
    session_id, time
) -> dict[str, any]:
    """Generate an EventSolutionAccepted class with random values.
    The "language" key may only be either "Undefined", "C", "CPP", "PHP", "Javascript", "Java", "Python"

    Args:
        date_start (datetime): [description]
        date_ends (datetime): [description]

    Returns:
        dict[str, any]: [description]
    """
    data = _generate_solution_event(session_id, time)
    return SolutionAcceptedEvent(**data).as_dictionary()


def generate_solution_rejected_event(
    session_id, time
) -> dict[str, any]:
    """Generate an EventSolutionRejected class with random values.
    The "language" key may only be either "Undefined", "C", "CPP", "PHP", "Javascript", "Java", "Python"

    Args:
        date_start (datetime): [description]
        date_ends (datetime): [description]

    Returns:
        dict[str, any]: [description]
    """
    data = _generate_solution_event(session_id, time)
    return SolutionRejectedEvent(**data).as_dictionary()


def generate_locale_set_event(
    session_id, time
) -> dict[str, any]:
    locale = random.choice(list(LOCALE.keys()))
    return LocaleSetEvent(session_id, time, locale).as_dictionary()


def generate_personal_info_submitted_event(
    user: dict[str, any], time: int
) -> dict[str, any]:
    return PersonalInfoSubmitedEvent(
        user["session_id"],
        time,
        user["student_number"],
        user["years_of_experience"],
        user["hours_of_practice"],
        user["familiar_languages"],
    ).as_dictionary()


def generate_session_started_event(
    session_id: str, time: int
) -> dict[str, any]:
    locale = random.choice(list(LOCALE.keys()))
    return_object = SessionStartedEvent(session_id, time, locale)
    return return_object.as_dictionary()


def generate_deadline_passed_event(
    session_id: str, time: int
) -> dict[str, any]:
    return_object = DeadlinePassedEvent(session_id, time)
    return return_object.as_dictionary()


def generate_exam_ended_event(
    session_id: str, time: int
) -> dict[str, any]:
    return_object = ExamEndedEvent(session_id, time)
    return return_object.as_dictionary()


def generate_exam_forfeited_event(
    session_id: str, time: int
) -> dict[str, any]:
    return_object = ExamForfeitedEvent(session_id, time)
    return return_object.as_dictionary()


def generate_exam_ide_reloaded_event(
    session_id: str, time: int
) -> dict[str, any]:
    return_object = ExamIDEReloadedEvent(session_id, time)
    return return_object.as_dictionary()


def generate_exam_started_event(
    session_id: str, time: int
) -> dict[str, any]:
    question_numbers = [1, 2, 3, 4, 5, 6]
    deadline = time + (90 * 60 * 1000)  # 90 minutes
    return_object = ExamStartedEvent(session_id, time, question_numbers, deadline)
    return return_object.as_dictionary()


def generate_before_exam_SAM_submited_event(
    session_id: str, time: int
) -> dict[str, any]:
    aroused_level = random.randint(0, 5)
    pleased_level = random.randint(0, 5)
    return_object = BeforeExamSAMSubmittedEvent(
        session_id, time, aroused_level, pleased_level
    )
    return return_object.as_dictionary()


def generate_after_exam_SAM_submited_event(
    session_id: str, time: int
) -> dict[str, any]:
    aroused_level = random.randint(0, 5)
    pleased_level = random.randint(0, 5)
    return_object = AfterExamSAMSubmittedEvent(
        session_id, time, aroused_level, pleased_level
    )
    return return_object.as_dictionary()


# checker
def fake_basic_info_generate():
    date_start_int: int = random_date(
        datetime(2021, 6, 1, 0, 0, 0), datetime(2021, 12, 29, 23, 59, 59)
    )
    date_start: datetime = datetime.fromtimestamp(date_start_int)
    additional_duration: timedelta = timedelta(minutes=random.randint(6, 21))
    date_ends: datetime = datetime.fromtimestamp(
        date_start_int + additional_duration.total_seconds()
    )
    return {
        "session_id": str(uuid4()),
        "date_start": date_start,
        "date_ends": date_ends,
    }


def generate_gibberish_code() -> str:
    return random.choice(
        [
            """
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
            """,
            """
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
            """,
            """
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
            """,
            """
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
            """,
            """
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
            """,
        ]
    )
