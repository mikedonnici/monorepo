from faker import Faker


def say_hi():
    fake = Faker()
    print(f"Hello {fake.name()}, from mods/fakomatic.foo")
