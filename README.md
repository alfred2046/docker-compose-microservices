# Instruction
This repo is to demonstrate to using docker compose spin up the dockerize golang microservices.

## Setup Workspace
This workspace is setup using VSCode with folders structure depicted below:
- microservices-docker-compose
- microservices-product-api
- microservices-user-api

## Microservices
To emulate, 2 microservices service is created with Dockerfile respectively.

## Docker Build Images
VSCode > Right click Dockerfile > Build Image...

<img width="491" height="551" alt="image" src="https://github.com/user-attachments/assets/87842921-08f1-4c2c-bebb-90d51e161cae" />

Enter prompt to tag image:

<img width="413" height="68" alt="image" src="https://github.com/user-attachments/assets/e8baa58c-b080-4b84-8d27-d178bc2eb82d" />

## docker-compose.yml
```
version: '3.8'
services:
  ms-user-services:
    image: microservicesuserapi
    ports:
      - "8081:8081"
    restart: unless-stopped
  ms-product-services:
    image: microservicesproductapi
    ports:
      - "8082:8082"
    restart: unless-stopped
```

## Spin Up Microservices with Docker Compose
VSCode > Right click docker-compose.yml > Compose Up

<img width="580" height="567" alt="image" src="https://github.com/user-attachments/assets/bdad8b16-63a0-43ea-9eb2-ab2579bbaee7" />

Congratulation! It's up and running!

<img width="322" height="92" alt="image" src="https://github.com/user-attachments/assets/ce6ad0c2-9c80-4107-b2d0-c20f6b3a6c35" />

## Emulate Auto-Restart
```
docker exec <container id> kill 1
```
