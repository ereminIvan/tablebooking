ifneq ($(GOBINPATH),)
	GOROOT:=$(subst /bin/go,,$(GOBINPATH))
	PATH:=$(GOROOT)/bin:$(PATH)
	SHELL:=PATH=$(PATH) GOROOT=$(GOROOT) $(SHELL)
endif

MIN_GO_VERSION:=1.7.0
GOPATH:=$(lastword $(subst :, ,$(GOPATH)))
PROJ_PATH:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
BINDIR?=$(lastword $(subst :, ,$(GOPATH)))/bin
GLIDE_PATH:=$(GOPATH)/src/github.com/Masterminds/glide
GLIDE_VERSION:=v0.12.3
GLIDE_BIN:=$(GLIDE_PATH)/glide-$(GLIDE_VERSION)
VENDOR_DIR=vendor
GLIDE_CHANGE=vendor/glide_change

DATE:=$(shell date -u "+%Y-%m-%d %H:%M:%S")
GOVER:=$(shell go version | cut -f3 -d " " | sed 's/go//')
APPVERSION:=$(shell git branch|grep '*'| cut -f2 -d' ')
GITREV:=$(shell git rev-parse --short HEAD)
GITLOG:=$(shell git log --decorate --oneline -n1| sed -e "s/'/ /g" -e "s/\"/ /g" -e "s/\#/\№/g" -e 's/`/ /g')
LDFLAGS=-X 'lazada_api/common.AppVersion=$(APPVERSION)' -X 'lazada_api/common.GitRev=$(GITREV)' -X 'lazada_api/common.GoVersion=$(GOVER)' -X 'lazada_api/common.BuildDate=$(DATE)' -X 'lazada_api/common.GitLog=$(GITLOG)'
LDFLAGS_STATIC= $(LDFLAGS) -extldflags '-static'
tags?=all
subpackage?=...
GO=go
NOVENDOR_DIRS=$$($(GLIDE_BIN) novendor)

IS_DESIRE_VERSION = $(shell expr $(GOVER) \> $(MIN_GO_VERSION))
ifeq ($(IS_DESIRE_VERSION),0)
$(error You have go version $(GOVER), need at least $(MIN_GO_VERSION))
endif

$(info GO VERSION: $(GOVER))

ifneq ($(fast),)
$(info TESTS: tables recreation disabled)
export GO_FAST_FIXTURES=1
endif

$(VENDOR_DIR):
	@mkdir -p $(VENDOR_DIR)

.PHONY: deps
deps:
	@$(MAKE) -B $(GLIDE_CHANGE)

$(GLIDE_BIN):
ifeq ($(wildcard $(GLIDE_BIN)),)
	$(info #Installing glide version $(GLIDE_VERSION)...)
ifeq ($(wildcard $(GLIDE_PATH)),)
	mkdir -p $(GLIDE_PATH) && cd $(GLIDE_PATH) ;\
	git clone https://github.com/Masterminds/glide.git .
endif
	cd $(GLIDE_PATH) && rm -rf vendor/ ;\
	git fetch --all && git checkout -f $(GLIDE_VERSION) ;\
	make clean && make build ;\
	mv glide $(GLIDE_BIN)
endif

.PHONY: get-glide
get-glide: $(GLIDE_BIN)

.PHONY: build
build:
	@cd $(PROJ_PATH)
	$(info Starting build...)
	@${GO} build -ldflags "$(LDFLAGS)" -o $(BINDIR) main.go

.PHONY: run
run: $(GLIDE_CHANGE)
	$(info Run...)
	@${GO} run -ldflags "$(LDFLAGSLDFLAGS)" main.go -config_dir=./bob_api/etc -env=dev
