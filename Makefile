NODE_ENV ?=

frontend:
	cd frontend && NODE_ENV=${NODE_ENV} npm run build

.PHONY: frontend

icons:
	(echo "//+build !windows" && go run github.com/cratonica/2goarray iconData main < icon.png) > icon.go
	convert icon.png icon.ico && \
		go run github.com/cratonica/2goarray iconData main < icon.ico > icon_windows.go && \
		rm -f icon.ico

.PHONY: icons

dll:
	curl -sSLO https://github.com/webview/webview/raw/master/dll/x64/WebView2Loader.dll
	curl -sSLO https://github.com/webview/webview/raw/master/dll/x64/webview.dll

.PHONY: dll

release:
	if [ "$(shell go env GOOS)" = "windows" ]; then make dll; fi
	NODE_ENV=production make frontend
	go build -o gui .
	7z a gui.7z gui
	if [ "$(shell go env GOOS)" = "windows" ]; then 7z a gui.7z *.dll; fi
	make clean

.PHONY: release

clean:
	rm -f gui *.dll

.PHONY: clean

distclean: clean
	rm -f gui.7z

.PHONY: distclean
