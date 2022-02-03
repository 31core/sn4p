all:
	go build -o sn4p_client client/*.go
	go build -o sn4p_server server/*.go

run_client:
	make -s all
	./sn4p_client

run_server:
	make -s all
	./sn4p_server
