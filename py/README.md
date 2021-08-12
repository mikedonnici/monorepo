# py

Contains Python code.

Uses [`poetry`](https://python-poetry.org/) for package management.

To [install `poetry`](https://python-poetry.org/docs/master/#installation)

```shell
curl -sSL https://raw.githubusercontent.com/python-poetry/poetry/master/install-poetry.py | python -
```

For a complete list of commands check out the [poetry cli](https://python-poetry.org/docs/master/cli/)

## Virtual Env

Poetry will create its own virtual environment by default which can cause confusion if an environment already exists.

See poetry [virtual environment](https://python-poetry.org/docs/master/basic-usage/#using-your-virtual-environment)
docs for details.

The easy way:

- Open the top-level project folder in the IDE.
- Run `poetry shell` to initialise / start the poetry venv - there's
  a [plugin for PyCharm](https://plugins.jetbrains.com/plugin/14307-poetry)
- `poetry add foo` to add a new dependency
- `poetry update` to update dependencies

## Create a local package

Run this command in the dir the package is to be created:

```shell
poetry new pkg_name
```

For an existing package run this command in the package root:

```shell
poetry init
```

## Add a dependency

Dependencies can be added in numerous ways - see
[poetry add](https://python-poetry.org/docs/master/cli/#add) docs for examples.

Local packages are added by specifying the relative path from the project root:

```shell
poetry add ../../pkg/pkg_name
```
