from faker import Faker


def say_hi():
    fake = Faker()
    print(f"Hi {fake.name()}, from hello v0.1.1.")
