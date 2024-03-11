.PHONY: profile
profile:
	go run ./... -cpuprofile=cpu.prof
	go tool pprof -http=: cpu.prof
