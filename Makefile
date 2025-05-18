build:
	go build -o altica_dht

docker-build:
	docker build -t altica_dht .

docker-run:
	docker run --rm -p 4001:4001 -v $(PWD)/peer.key:/app/peer.key altica_dht

run:
	./altica_dht
