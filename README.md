#REST-API

##user-service

GET /users - list of users - 200, 404, 500
GET /users/:id - user by id - 200, 404, 500
POST /users - create user - 204, 4xx, Header Location: url
PUT /users/:id - full update user - 204/200, 400, 404, 500
PATCH /users/:id - update user - 204/200, 400, 404, 500
DELETE /users/:id - delete user by id - 204, 404, 400