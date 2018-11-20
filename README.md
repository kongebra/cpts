# Computer Party Ticket System

## Deployment

###Automatic build and deploy
```
sudo run ./start.sh
```

###Manual build and deploy
```
docker-compose rm -f
docker-compose pull
docker-compose up --build -d
```


## Configure
###Docker-Compose
Configure docker-compose.yaml if another version of Golang, MongoDB or Nginx is wanted. 

Also the Dockerfile's in /api, /web folders.

## Testing



## Running

# Spec

## Routes


## API

### GET /api

### GET /api/event

### POST /api/event

### DELETE /api/event