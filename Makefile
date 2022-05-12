.PHONY: build stage prod

repo=vivaconagua/donation-feed-backend

build:
	docker build -t ${repo}:stage .

stage:
	docker push ${repo}:stage

prod:
	docker tag ${repo}:stage ${repo}:latest
	docker push ${repo}:latest
