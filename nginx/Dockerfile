FROM ubuntu:22.04
WORKDIR /ngx
COPY ./nginx-1.24.0/ .
COPY ./build-ngx/add_configure.sh .
COPY ./build-ngx/nginx.service /etc/systemd/system/
RUN sh ./add_configure.sh

EXPOSE 80
STOPSIGNAL SIGQUIT

CMD ["nginx", "-g", "daemon off;"]