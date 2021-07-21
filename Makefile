docker-build:
	docker build -t stone-challenge .

docker-run:
	docker run -d -p 8000:8000 --namen stoneAPI stone-challenge