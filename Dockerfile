# 第一阶段：构建阶段
FROM golang:1.20 AS builder  

# 设置工作目录
WORKDIR /root/project/EugeneGoBlog

# 将项目文件复制到容器内
COPY . .

# 编译 Go 程序，生成静态链接的二进制文件
RUN go build -o main .

# 第二阶段：运行阶段
FROM scratch

# 将构建好的二进制文件从 builder 阶段复制到最终镜像中
COPY --from=builder /root/project/EugeneGoBlog/main /main

# 暴露端口
EXPOSE 80

# 设置容器启动时执行的命令
CMD ["/main"]
