FROM alpine:latest

WORKDIR /app

COPY Deployment_task/main /app/main 

RUN chmod +x main

ENTRYPOINT ["./main"]
