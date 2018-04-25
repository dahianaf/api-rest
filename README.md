# api-rest

Is a Restful Api coded in Golang and Dockerized. 

Run Api:
Golang: go build && ./api-rest
Docker: docker run --name api-rest --rm -p 8080:8080 dfalcon/api-rest

Test Api:
On browser: http://localhost:8080/test
On prompt: curl -i http://localhost:8080/test

Show data:
On browser: http://localhost:8080/servers
On prompt: curl -i http://localhost:8080/servers

For one server:
On browser: http://localhost:8080/servers/1
On prompt: curl -i http://localhost:8080/servers/1


To try POST, PUT, PATCH and DELETE methods you can use Postman or with the following commands (recommended):

GET test response
curl -i http://localhost:8080/test

GET all servers
curl -i http://localhost:8080/servers

GET the server with ID=1
curl -i http://localhost:8080/servers/1

POST a new server with ID=4
curl -d '{"ID":"4", "Name":"fnadg-29", "Cores":"4", "Memory":"8GB", "Disk":"50GB"}' http://localhost:8080/servers/4

PUT method (in a worng way to try the PATCH method)
curl -X PUT -d '{"ID":"2", "Name":"vxadl-20", "cores":"4", "Disk":"50GB"}' http://localhost:8080/servers/2
curl -X PATCH -d '{"Memory":"8GB"}' http://localhost:8080/servers/2

curl -i -X DELETE http://localhost:8080/servers/4
