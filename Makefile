frontend:
	cd frontend && npm run build

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
