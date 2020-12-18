
FROM ubuntu:18.04  

WORKDIR /

COPY todo .
RUN ls -l
RUN ls -l todo
EXPOSE 8080

ENTRYPOINT ["./todo" ]


