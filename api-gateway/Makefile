run:
	go mod tidy;
	make proto-gen;
	make swag-gen;
	clear;
	go run main.go;

migrate_up:
	migrate -path internal/migrations -database postgres://mrbek:QodirovCoder@localhost:5432/gmail_db -verbose up

migrate_down:
	migrate -path internal/migrations -database postgres://mrbek:QodirovCoder@localhost:5432/gmail_db -verbose down

migrate_force:
	migrate -path internal/migrations -database postgres://mrbek:QodirovCoder@localhost:5432/gmail_db -verbose force 1

migrate_file:
	migrate create -ext sql -dir internal/migrations -seq create_table

swag-gen:
	~/go/bin/swag init -g ./internal/http/router.go -o ./internal/http/docs force 1

proto-gen:
	protoc --go_out=./ --go-grpc_out=./ gmailapp-submodule/*.proto
	
push:
	cd muallimah-submodule && git add . &&	git commit -m "protos changed" &&git push origin dev

drun:
	docker run -p 9000:9000 -p 9001:9001 --name minio -v ~/minio/data:/data -e "MINIO_ROOT_USER=user" -e "MINIO_ROOT_PASSWORD=password" quay.io/minio/minio server /data --console-address ":9001"

spull:
	cd gmailapp-submodule && git pull
	make proto-gen