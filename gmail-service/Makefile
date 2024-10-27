run:
	go mod tidy
	clear
	go run cmd/main.go

proto-gen:
	protoc --go_out=./ --go-grpc_out=./ gmailapp-submodule/*.proto

migrate_up:
	migrate -path internal/migrations -database postgres://mrbek:QodirovCoder@localhost:5432/gmail_db -verbose up

migrate_down:
	migrate -path internal/migrations -database postgres://mrbek:QodirovCoder@localhost:5432/gmail_db -verbose down

migrate_force:
	migrate -path internal/migrations -database postgres://mrbek:QodirovCoder@localhost:5432/gmail_db -verbose force 1

migrate_file:
	migrate create -ext sql -dir internal/migrations -seq create_table

make spush:
	cd gmailapp-submodule && git add . && git commit -m "protos changed" && git push
	make proto-gen