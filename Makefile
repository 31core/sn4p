all:
	go build -o sn4p *.go

run:
	make -s all
	./sn4p