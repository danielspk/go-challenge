FROM golang:1.15

RUN mkdir /api
ADD . /api
WORKDIR /api
RUN make go-build

EXPOSE 8000

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.7.3/wait /wait
RUN chmod +x /wait

CMD ["sh", "-c", "/wait && ./build/api"]
