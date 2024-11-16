.PHONY: deploy

deploy:
	git pull
	sudo docker build -t eugene-go-blog .
	sudo docker tag eugene-go-blog sgccr.ccs.tencentyun.com/eugene_images/blog:latest
	sudo docker push sgccr.ccs.tencentyun.com/eugene_images/blog:latest
	sudo docker image prune -f
	sudo docker run -d -p 80:80 eugene-go-blog
