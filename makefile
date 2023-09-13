docker-central: 
	docker build --tag docker-central .
	docker run -it --rm --name docker-central docker-central

docker-regional: docker-asia docker-europa

docker-rabbit:
	docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.12-management

docker-asia:
	docker build --tag docker-asia ./asia
	docker run -it --rm --name docker-asia docker-asia

docker-europa:
	docker build --tag docker-europa ./europa
	docker run -it --rm --name docker-europa docker-europa