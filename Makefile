all:
	go build -o rvspub ./cmd/rvspub

clean:
	rm -f rvspub
