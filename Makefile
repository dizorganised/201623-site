SASS_COMPILER ?= sassc

SASS_DIR    := sass/
SASS_SRC    := $(shell find sass/ -name '*.scss')
MAIN_FILE   := sass/main.scss
OUTPUT_FILE := public/css/main.scss.css

all: build-sass
	go build

run: all
	./2016-site

build-sass: $(SASS_SRC)
	$(SASS_COMPILER) $(MAIN_FILE) $(OUTPUT_FILE)

clean:
	rm -f $(OUTPUT_FILE)

.PHONY: build-sass run clean
