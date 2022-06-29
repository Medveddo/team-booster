docker:
	bash -c "docker start mongo mongo-express"
	
back:
	bash -c "cd backend && go run cmd/server/main.go"

dev:
	bash -c "cd backend && go run cmd/dev/main.go"
	