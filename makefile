VENV := .venv
PYTHON := python3

.PHONY: setup activate install serve clean

setup:
	$(PYTHON) -m venv $(VENV)

install: setup
	. $(VENV)/bin/activate && pip install mkdocs && pip install mkdocs-material

serve: install
	. $(VENV)/bin/activate && mkdocs serve -a 0.0.0.0:2442

check:
	. $(VENV)/bin/activate && mkdocs --version

clean:
	rm -rf $(VENV)
