dist-clean:
	rm -rf dist
dist: dist-clean
	mkdir -p dist/alpine-linux/amd64 && GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -a -tags netgo -installsuffix netgo -o dist/alpine-linux/amd64/confin
	mkdir -p dist/linux/amd64 && GOOS=linux GOARCH=amd64 go build -ldflags "$(LDFLAGS)" -o dist/linux/amd64/confin
	mkdir -p dist/linux/armel && GOOS=linux GOARCH=arm GOARM=5 go build -ldflags "$(LDFLAGS)" -o dist/linux/armel/confin
	mkdir -p dist/linux/armhf && GOOS=linux GOARCH=arm GOARM=6 go build -ldflags "$(LDFLAGS)" -o dist/linux/armhf/confin

release: dist
	tar -cvzf confin-alpine-linux-amd64-$(TAG).tar.gz -C dist/alpine-linux/amd64 confin
	tar -cvzf confin-linux-amd64-$(TAG).tar.gz -C dist/linux/amd64 confin
	tar -cvzf confin-linux-armel-$(TAG).tar.gz -C dist/linux/armel confin
	tar -cvzf confin-linux-armhf-$(TAG).tar.gz -C dist/linux/armhf confin