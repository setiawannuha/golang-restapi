APP=exampleservice
APP_EXECUTABLE="./out/$(APP)"
build:
			mkdir -p ./out && go build -o ${APP_EXECUTABLE}
run: build
			${APP_EXECUTABLE} start-server
migrate-mysql: build
			${APP_EXECUTABLE} migrate-mysql
rollback-mysql: build
			${APP_EXECUTABLE} rollback-mysql
migrate-pg: build
			${APP_EXECUTABLE} migrate-pg
rollback-pg: build
			${APP_EXECUTABLE} rollback-pg
mysql-migrate-up:
			migrate -path migration/mysql -database "mysql://root:password@tcp(127.0.0.1:3306)/db_test" -verbose up
mysql-migrate-down:
			migrate -path migration/mysql -database "mysql://root:password@tcp(127.0.0.1:3306)/db_test" -verbose down
pg-migrate-up:
			migrate -path migration/postgres -database "postgres://postgres:password@127.0.0.1:5432/db_test?sslmode=disable" -verbose up
pg-migrate-down:
			migrate -path migration/postgres -database "postgres://postgres:password@127.0.0.1:5432/db_test?sslmode=disable" -verbose down