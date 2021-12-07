echo 'Making calls to List endpoint'


#hey -m GET 'http://localhost:8080/v1/airports?cursor=1'
#hey -m GET 'http://localhost:8080/v1/airports/1'
hey -m POST -d '{ "name": "LAX", "location": "string" }' 'http://localhost:8080/v1/airports/'
