from faker import Faker


def say_hi():
    fake = Faker()
    print(f"Hi {fake.name()}, from the improved hello package.")
