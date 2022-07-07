.DEFAULT_GOAL := up

docker:
	bash -c "docker start mongo mongo-express"
	
back:
	bash -c "cd backend && go run cmd/server/main.go"

front:
	bash -c "cd frontend && npm run dev"

dev:
	bash -c "cd backend && go run cmd/dev/main.go"

echo:
	bash -c "echo 'sam echo'"

up:
	bash -c "docker-compose up --build"
