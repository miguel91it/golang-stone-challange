docker-build:
	docker build -t stone-challenge .

docker-run:
	docker run -d -p 16453:16453 --name stoneAPI stone-challenge

docker-run-attached:
	docker run -p 16453:16453 --name stoneAPI stone-challenge

docker-remove:
	docker rm -fv stoneAPI