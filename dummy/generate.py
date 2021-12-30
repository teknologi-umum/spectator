# Code documentation soon

import random
import string
import json

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

if __name__ == "__main__":
    main()