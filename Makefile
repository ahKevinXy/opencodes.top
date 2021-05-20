PROJECT:=opencodes

.PHONY: build
build:
	if [ ! -d log ]; then mkdir log; fi
	gofmt -w -s .
	CGO_ENABLED=0  go build -o ./bin/opencodes  main.go
	@echo "构建成功"

reload:
	kill -USR2 `cat pid/*.pid`
	echo '重启成功'
start:
	if [ ! -d pid ]; then mkdir pid; fi
	export GOTRACEBACK=crash
	ulimit -c unlimited

	bin/opencodes  server  -c config/settings.dev.yml -p 8000 -m dev >> log/panic.log 2>&1 &

	@echo "start successfully"
stop:
	kill `cat pid/*.pid`
	sleep 1
	rm -rf pid/*.pid

	@echo "停止成功"
