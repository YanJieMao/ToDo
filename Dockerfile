

FROM alpine:3.12

WORKDIR /
COPY todo .
RUN ls -l
RUN ls -l todo
EXPOSE 8080

ENTRYPOINT ["./todo" "config" " --serverPort "8080" " "--dbType "mysql" " "--dbHost "localhost" " "--dbPort "3306" " "--dbName "todo" " "--dbUser "root" " " --dbPasswd "1997""]




