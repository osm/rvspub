all:
	go build -o rvspub ./cmd/rvspub
	go build -o rvscrypto ./cmd/rvscrypto

rvspub:
	go build -o rvspub ./cmd/rvspub

rvscrypto:
	go build -o rvscrypto ./cmd/rvscrypto

clean:
	rm -f rvspub rvscrypto
