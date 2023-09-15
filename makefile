docker-central: 
	sudo docker build --tag docker-central .
	sudo docker run -it --rm --name docker-central docker-central

docker-regional: docker-asia docker-europa docker-america docker-oceania

docker-rabbit:
	sudo docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.12-management

docker-asia:
	sudo docker build --tag docker-asia ./asia
	sudo docker run -it --rm --name docker-asia docker-asia

docker-europa:
	sudo docker build --tag docker-europa ./europa
	sudo docker run -it --rm --name docker-europa docker-europa

docker-america:
	sudo docker build --tag docker-america ./america
	sudo docker run -it --rm --name docker-america docker-america

docker-oceania:
	sudo docker build --tag docker-oceania ./oceania
	sudo docker run -it --rm --name docker-oceania docker-oceania
