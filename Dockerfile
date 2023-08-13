FROM alpine
ADD build/auto-migrate build/server-upx init.sh /app/
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && apk update && apk add tini && chmod +x /app/init.sh
ENTRYPOINT [ "tini", "--", "/app/init.sh"]