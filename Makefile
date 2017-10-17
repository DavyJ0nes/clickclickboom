# Applicaiton: clickclickboom
# Author:      DavyJ0nes
# Github:      https://github.com/davyj0nes/clickclickboom
#
all: build_osx

app = clickclickboom
app_version =
goversion ?= 1.9
git_hash = $(shell git rev-parse HEAD | cut -c 1-6)
build_date = $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')

build_osx: check_env
	$(call blue, "# Building OSX Binary...")
	docker run --rm -it -v "$(CURDIR)":/go/src/app -w /go/src/app golang:${goversion} sh -c 'go get && GOOS=darwin GOARCH=amd64 go build -X main.version=${app_version} -X main.gitHash=${git_hash} -X main.date=${build_date} -o releases/${app}_osx'

build_linux: check_env
	$(call blue, "# Building Linux Binary...")
	docker run --rm -it -v "$(CURDIR)":/go/src/app -w /go/src/app golang:${goversion} sh -c 'go get && GOOS=linux GOARCH=amd64 go build -X main.version=${app_version} -X main.gitHash=${git_hash} -X main.date=${build_date} -o releases/${app}_linux'

build_win: check_env
	$(call blue, "# Building Windows Binary...")
	docker run --rm -it -v "$(CURDIR)":/go/src/app -w /go/src/app golang:${goversion} sh -c 'go get && GOOS=windows GOARCH=amd64 go build -X main.version=${app_version} -X main.gitHash=${git_hash} -X main.date=${build_date} -o releases/${app}.exe'

install: build_osx
	$(call blue, "# Moving Binary to bin...")
	chmod +x ${app}
	mv ${app} "$(HOME)/bin"

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
