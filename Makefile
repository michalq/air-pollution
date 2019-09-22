
local-db:
	docker run \
		--rm \
		--name air-pollution \
		-e POSTGRES_DB=docker \
		-e POSTGRES_USER=docker \
		-e POSTGRES_PASSWORD=docker \
		-d -p 5433:5432 \
		-v /home/michal/Dropbox/air-pollution:/var/lib/postgresql/data \
		postgres:10

local-db-config:
	docker exec -it air-pollution cat /etc/postgresql/9.3/main/postgresql.conf
