import hello.hello
from hello import __version__


def test_version():
    assert __version__ == '0.1.1'


def test_say_hi():
    hello.hello.say_hi()
