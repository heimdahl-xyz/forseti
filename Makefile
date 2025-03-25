
migrate_up:
	migrate -database ${FORSETI_DB_URL} -path ./migrations up

