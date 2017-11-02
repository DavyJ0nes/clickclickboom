# App:     clickclickboom
# Author:  DavyJ0nes
# Github:  https://github.com/davyj0nes/clickclickboom
#

#### SET DEFAULT BEHAVIOUR ####
all: build_osx

#### DEFINE VARIABLES ####
app = clickclickboom
app_version =
goversion ?= 1.9
git_hash = $(shell git rev-parse HEAD | cut -c 1-6)
build_date = $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')
go_build_cmd = go get && GOOS=darwin GOARCH=amd64 go build --ldflags=" -X main.version=${app_version} -X main.gitHash=${git_hash} -X main.date=${build_date}" -o releases/${app}-${app_version}

#### BUILD STEPS ####
build_osx: check_env
	$(call blue, "# Building OSX Binary...")
	docker run --rm -it -v "$(CURDIR)":/go/src/app -w /go/src/app golang:${goversion} sh -c '${go_build_cmd}-osx'

build_linux: check_env
	$(call blue, "# Building Linux Binary...")
	docker run --rm -it -v "$(CURDIR)":/go/src/app -w /go/src/app golang:${goversion} sh -c '${go_build_cmd}-linux'

build_win: check_env
	$(call blue, "# Building Windows Binary...")
	docker run --rm -it -v "$(CURDIR)":/go/src/app -w /go/src/app golang:${goversion} sh -c '${go_build_cmd}.exe'

build_all: build_osx build_linux build_win
	$(call blue, "!! Built For all Platforms")

check_env:
ifndef app_version
	$(error app_version is undefined)
endif

define blue
	@tput setaf 4
	@echo $1
	@tput sgr0
endef
