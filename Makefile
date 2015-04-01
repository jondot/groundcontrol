SRC = *.go
PKG = groundcontrol README.md groundcontrol.json.sample web support
PKG_NAME = groundcontrol
DEBPKG = groundcontrol README.md web
VERSION=$(shell cat version.go | perl -n -e'/VERSION = "(.*?)"/ && print $$1')


build: $(SRC)
	go build

package: $(PKG)
	mkdir -p groundcontrol-$(VERSION)
	cp -r $(PKG) groundcontrol-$(VERSION)/
	tar -cvzf groundcontrol-$(VERSION).tar.gz  groundcontrol-$(VERSION)
	rm -rf groundcontrol-$(VERSION)/

package_deb: clean build
	mkdir -p ./build/opt/groundcontrol
	mkdir -p ./build/etc/init.d
	cp -r $(DEBPKG) ./build/opt/groundcontrol/.
	cp groundcontrol.json.sample ./build/etc/groundcontrol.json
	cp support/init.d/groundcontrol ./build/etc/init.d/.
	fpm -s dir \
            -t deb \
            -n $(PKG_NAME) \
            -v $(VERSION) \
            --license "MIT" \
            -m "Jochen Breuer <brejoc@gmail.com>" \
            --description "Manage and monitor your Raspberry Pi with ease." \
            --url "http://jondot.github.io/groundcontrol" \
            --deb-user root \
            --deb-group root \
            -C ./build \
	    etc opt

clean:
	rm -rf ./build
	rm -f groundcontrol
	rm -f groundcontrol-$(VERSION).tar.gz
	rm -f groundcontrol_$(VERSION)*.deb

.PHONY: clean
