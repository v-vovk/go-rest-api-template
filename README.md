#REST-API

##user-service

`GET /users` - list of users - 200, 404, 500 <br>
`GET /users/:id` - user by id - 200, 404, 500 <br>
`POST /users` - create user - 204, 4xx, Header Location: url <br>
`PUT /users/:id` - full update user - 204/200, 400, 404, 500 <br>
`PATCH /users/:id` - update user - 204/200, 400, 404, 500 <br>
`DELETE /users/:id` - delete user by id - 204, 404, 400 <br>