include .env
export

migrate:
	@for file in ./db/migrations/up/*; \
	do \
	  	psql postgresql://$$DB_USER:$$DB_PASS@$$DB_HOST/$$DB_NAME -a -f "$$file"; \
	done

migrate-down:
	@for file in `ls ./db/migrations/down/* | sort -r`; \
	do \
	  psql postgresql://$$DB_USER:$$DB_PASS@$$DB_HOST/$$DB_NAME -a -f "$$file"; \
	done

test:
	curl -X POST http://localhost:3000/articles -d '{"title": "test", "body": "long body"}'

interactive-db:
	psql postgresql://$$DB_USER:$$DB_PASS@$$DB_HOST/$$DB_NAME