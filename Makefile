all:
	go build -o rvspub ./cmd/rvspub
	go build -o rvscrypto ./cmd/rvscrypto
	go build -o rvsserver ./cmd/rvsserver

rvspub:
	go build -o rvspub ./cmd/rvspub

rvscrypto:
	go build -o rvscrypto ./cmd/rvscrypto

rvsserver:
	go build -o rvsserver ./cmd/rvsserver

clean:
	rm -f rvspub rvscrypto
