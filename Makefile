docker-build:
	docker build -t stone-challenge .

docker-run:
	docker run -d -p 16453:16453 --namen stoneAPI stone-challenge