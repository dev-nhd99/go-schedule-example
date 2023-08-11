FROM alpine:latest

ENV APP_NAME go-schedule-example
RUN apk --no-cache add ca-certificates
WORKDIR /onsky/apps/
COPY ${APP_NAME} . 
EXPOSE 1323

ENTRYPOINT [ "./go-schedule-example" ]