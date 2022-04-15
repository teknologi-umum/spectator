import string
import random
import uuid


class User:
    type: str
    session_id: str
    student_number: str
    hours_of_practice: int
    years_of_experience: int
    familiar_languages: str

    def __init__(
        self,
        session_id: str,
        student_number: str,
        hours_of_practice: int,
        years_of_experience: int,
        familiar_languages: str,
    ):
        self.type = "personal_info_submitted"
        self.session_id = session_id
        self.student_number = student_number
        self.hours_of_practice = hours_of_practice
        self.years_of_experience = years_of_experience
        self.familiar_languages = familiar_languages

    def as_dictionary(self):
        return {
            "type": self.type,
            "session_id": self.session_id,
            "student_number": self.student_number,
            "hours_of_practice": self.hours_of_practice,
            "years_of_experience": self.years_of_experience,
            "familiar_language": self.familiar_languages,
        }


def generate_user() -> dict[str, any]:
    languages = ["python", "javascript", "php", "", "c", "c++", "c#", "pascal"]
    numbers = string.digits
    session_id = str(uuid.uuid4())
    user = User(
        session_id=session_id,
        student_number="".join(random.choice(numbers) for _ in range(8)),
        hours_of_practice=random.randint(0, 5),
        years_of_experience=random.randint(0, 5),
        familiar_languages=" ".join(random.sample(languages, k=random.randint(1, 3))),
    )
    return user.as_dictionary()
