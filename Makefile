all: deps run lintfix

.PHONY: deps
deps:
	go get github.com/go-telegram-bot-api/telegram-bot-api/v5
	go get github.com/ilyakaznacheev/cleanenv

.PHONY: run
run:
	go run cmd/main.go

UNAME := $(shell uname)

.PHONY: lintfix
lintfix:
    ifeq ($(UNAME), Linux)
		find . \( -path './cmd/*' -or -path './internal/*' -or -path './pkg/*' -or -path './e2e/*' \) \
		-type f -name '*.go' -print0 | \
		xargs -0  sed -i '/import (/,/)/{/^\s*$$/d;}'
    endif
    ifeq ($(UNAME), Darwin)
		find . \( -path './cmd/*' -or -path './internal/*' -or -path './pkg/*' -or -path './e2e/*' \) \
		-type f -name '*.go' -print0 | \
		xargs -0  sed -i '' '/import (/,/)/{/^\s*$$/d;}'
    endif
