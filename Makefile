
EXECUTABLE := bitcoinpay
GITVER := $(shell git rev-parse --short=7 HEAD )
GITDIRTY := $(shell git diff --quiet || echo '-dirty')
GITVERSION = "$(GITVER)$(GITDIRTY)"
DEV=dev
RELEASE=release
LDFLAG_DEV = -X github.com/btceasypay/bitcoinpay/version.Build=$(DEV)-$(GITVERSION)
LDFLAG_RELEASE = -X github.com/btceasypay/bitcoinpay/version.Build=$(RELEASE)-$(GITVERSION)
GOFLAGS_DEV = -ldflags "$(LDFLAG_DEV)"
GOFLAGS_RELEASE = -ldflags "$(LDFLAG_RELEASE)"
VERSION=$(shell ./build/bin/bitcoinpay --version | grep ^bitcoinpay | cut -d' ' -f3|cut -d'+' -f1)
GOBIN = ./build/bin

UNIX_EXECUTABLES := \
	build/release/darwin/amd64/bin/$(EXECUTABLE) \
	build/release/linux/amd64/bin/$(EXECUTABLE)
WIN_EXECUTABLES := \
	build/release/windows/amd64/bin/$(EXECUTABLE).exe

EXECUTABLES=$(UNIX_EXECUTABLES) $(WIN_EXECUTABLES)
	
COMPRESSED_EXECUTABLES=$(UNIX_EXECUTABLES:%=%.tar.gz) $(WIN_EXECUTABLES:%.exe=%.zip) $(WIN_EXECUTABLES:%.exe=%.cn.zip)

RELEASE_TARGETS=$(EXECUTABLES) $(COMPRESSED_EXECUTABLES)

ZMQ = FALSE

.PHONY: bitcoinpay bx release

bitcoinpay: bitcoinpay-build
	@echo "Done building."
	@echo "  $(shell $(GOBIN)/bitcoinpay --version))"
	@echo "Run \"$(GOBIN)/bitcoinpay\" to launch."

bitcoinpay-build:
    ifeq ($(ZMQ),TRUE)
		@echo "Enalbe ZMQ"
		@go build -o $(GOBIN)/bitcoinpay $(GOFLAGS_DEV) -tags=zmq "github.com/btceasypay/bitcoinpay/cmd/bitcoinpay"
    else
		@go build -o $(GOBIN)/bitcoinpay $(GOFLAGS_DEV) "github.com/btceasypay/bitcoinpay/cmd/bitcoinpay"
    endif
bx:
	@go build -o $(GOBIN)/bx "github.com/btceasypay/bitcoinpay/cmd/bx"

checkversion: bitcoinpay-build
#	@echo version $(VERSION)

all: bitcoinpay-build bx

# amd64 release
build/release/%: OS=$(word 3,$(subst /, ,$(@)))
build/release/%: ARCH=$(word 4,$(subst /, ,$(@)))
build/release/%/$(EXECUTABLE):
	@echo Build $(@)
	@GOOS=$(OS) GOARCH=$(ARCH) go build $(GOFLAGS_RELEASE) -o $(@) "github.com/btceasypay/bitcoinpay/cmd/bitcoinpay"
build/release/%/$(EXECUTABLE).exe:
	@echo Build $(@)
	@GOOS=$(OS) GOARCH=$(ARCH) go build $(GOFLAGS_RELEASE) -o $(@) "github.com/btceasypay/bitcoinpay/cmd/bitcoinpay"

%.zip: %.exe
	@echo zip $(EXECUTABLE)-$(VERSION)-$(OS)-$(ARCH)
	@zip $(EXECUTABLE)-$(VERSION)-$(OS)-$(ARCH).zip "$<"

%.cn.zip: %.exe
	@echo Build $(@).cn.zip
	@echo zip $(EXECUTABLE)-$(VERSION)-$(OS)-$(ARCH)
	@zip -j $(EXECUTABLE)-$(VERSION)-$(OS)-$(ARCH).cn.zip "$<" script/win/start.bat

%.tar.gz : %
	@echo tar $(EXECUTABLE)-$(VERSION)-$(OS)-$(ARCH)
	@tar -zcvf $(EXECUTABLE)-$(VERSION)-$(OS)-$(ARCH).tar.gz "$<"
release: clean checkversion
	@echo "Build release version : $(VERSION)"
	@$(MAKE) $(RELEASE_TARGETS)
	@shasum -a 512 $(EXECUTABLES) > $(EXECUTABLE)-$(VERSION)_checksum.txt
	@shasum -a 512 $(EXECUTABLE)-$(VERSION)-* >> $(EXECUTABLE)-$(VERSION)_checksum.txt
checksum: checkversion
	@cat $(EXECUTABLE)-$(VERSION)_checksum.txt|shasum -c
clean:
	@rm -f *.zip
	@rm -f *.tar.gz
	@rm -f ./build/bin/bx
	@rm -f ./build/bin/bitcoinpay
	@rm -rf ./build/release
