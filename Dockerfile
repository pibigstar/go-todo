# 源镜像(一个极小的Linux)
FROM loads/alpine:3.8
# 作者
LABEL maintainer="pibigstar@sina.com"

# 设置固定的项目路径
ENV WORKDIR /var/www/go-todo

# 添加应用可执行文件，并设置执行权限
# 将编译后的可执行文件复制到docker中
COPY ./scritps/main $WORKDIR/main
RUN chmod +x  $WORKDIR/main

# 添加静态文件、配置文件、模板文件
COPY conf   $WORKDIR/conf
COPY https   $WORKDIR/https
#COPY public   $WORKDIR/public
#COPY template $WORKDIR/template

# 启动
WORKDIR $WORKDIR
ENTRYPOINT ["./main"]