.PHONY: deploy
deploy:
	git pull
	CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o main .
	docker build -t eugene-go-blog .
	rm main
	# 阿里云
	docker tag eugene-go-blog registry.cn-hangzhou.aliyuncs.com/eugene_images/blog:latest
	docker push registry.cn-hangzhou.aliyuncs.com/eugene_images/blog:latest
	# 腾讯云
	docker tag eugene-go-blog sgccr.ccs.tencentyun.com/eugene_images/blog:latest
	docker push sgccr.ccs.tencentyun.com/eugene_images/blog:latest
	docker image prune -f

.PHONY: run
run:
	# 默认使用腾讯云
	docker pull sgccr.ccs.tencentyun.com/eugene_images/blog:latest
	docker run -d -p 80:80 sgccr.ccs.tencentyun.com/eugene_images/blog:latest