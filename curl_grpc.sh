echo 'Making calls to grpc server'

printf '\nCalling CreateAirportV1'
grpc_cli call localhost:8082 CreateAirportV1 "name: 'LAX', location: 'Los-Angeles'"

echo '\nCalling DescribeAirportV1'
grpc_cli call localhost:8082 DescribeAirportV1 "airport_id: 1"

echo '\nCalling ListAirports'
grpc_cli call localhost:8082 ListAirportsV1 ""
echo '\nCalling DeleteAirport'

grpc_cli call localhost:8082 DeleteAirportV1 "airport_id: 1"
