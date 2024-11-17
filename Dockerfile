# 使用最小化的 scratch 镜像作为基础镜像
FROM scratch
# 将本地编译好的二进制文件复制到镜像中
COPY ./main /main
# 如果有静态资源或配置文件需要复制，也一并添加
COPY ./etc /etc
# 暴露服务端口
EXPOSE 80
# 设置容器启动时执行的命令
CMD ["/main"]
