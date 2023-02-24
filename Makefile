UNAME := $(shell uname)
CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
LINTVER=latest
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}

.PHONY: lintfix lint install-lint bindir run

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
	goimports -local=./ -w ./cmd ./internal ./pkg

lint: install-lint
	${LINTBIN} run


install-lint: bindir
	test -f ${LINTBIN} || \
  (GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
  mv ${BINDIR}/golangci-lint ${LINTBIN})

bindir:
	mkdir -p ${BINDIR}

run:
	go run cmd/main.go