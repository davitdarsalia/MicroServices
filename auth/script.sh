# Initialization of migrations
menuAPI % migrate create -ext sql -dir ./schema -seq init
# Migration UP
menuAPI % migrate -path ./schema -database 'postgres://postgres:asdASD123@localhost:5436/postgres?sslmode=disable' up
