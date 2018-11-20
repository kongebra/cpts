### Live application can be found at http://10.212.138.218:3000 (API) & http://10.212.138.218


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

## Web
### URL
####Event
* http://10.212.138.218/event/index.html
* http://10.212.138.218/event/create.html
####User
 * http://10.212.138.218/user/index.html
 * http://10.212.138.218/user/create.html
####Ticket
* http://10.212.138.218/ticket/index.html
* http://10.212.138.218/ticket/create.html


## API
###Base URL
http://10.212.138.218:3000

### GET /api/event
Displays all current events

Response:
```json
[
  {
    "id": <value>,
    "name": <value>,
    "description": <value>,
    "date": {
      "start": <value>,
      "end": <value>
    },
    "participants": [],
    "img_url": <value>
  }
]
```
### POST /api/event
Creates new event

Body:
```json
{
  "name": <value>,
  "description": <value>,
  "date": {
    "start": <value>,
    "end": <value>
  },
  "image": <value>
}
```

### GET /api/user
Display all current users

Response:
```json
[
  {
    "id": <value>,
    "username": <value>,
    "email": <value>,
    "tickets": []
  }
]
```
### POST /api/user
Creates user

Body:
```json
{
  "username": <value>,
  "email": <value>,
  "password": <value>
}
```
### GET /api/ticket
Displays all tickets

Response:
```json
[
  {
    "id": <value>,
    "event": <value>,
    "scanned": <value>
  }
]
```
### POST /api/ticket
Creates a ticket for an user to an event

Body:
```json
{
  "event": <value>,
  "user": <value>
}
```

### POST /api/event/webhooks
Creates a Discord-webhook

Body:
```json
{
  "url": <discord-webhook-url>
}
```
