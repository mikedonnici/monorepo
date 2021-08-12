from faker import Faker


def say_hi():
    fake = Faker()
    print(f"Hi {fake.name()}, from the all new and improved hello package.")
