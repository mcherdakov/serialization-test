run:
	docker compose up --build

gen_proto:
	protoc -I./proto --go_out=internal/gen ./proto/animalshelter.proto
