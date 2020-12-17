

FROM alpine:3.12

WORKDIR /
COPY todo .
RUN ls -l
RUN ls -l todo
EXPOSE 8080

ENTRYPOINT ["./todo","config","--serverPort "8080"" ]




