dbinit:
	@docker run -e POSTGRES_PASSWORD=bootdotdev --name=pg-blogator --rm -d -p 5432:5432 postgres && sleep 3
	@docker exec -u postgres -it pg-blogator psql -c "CREATE DATABASE blogator;"
migrate:
	@cd sql/schema && goose postgres postgres://postgres:bootdotdev@172.17.240.1:5432/blogator up && cd ../..

build:
	@go build -o ./bin/app
run: build
	@./bin/app