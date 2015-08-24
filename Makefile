full:
	go test . -bench=. -cover -v

test:
	go test . -count 500

benchmark:
	go test . -bench=.

cover:
	go test . -cover -race
