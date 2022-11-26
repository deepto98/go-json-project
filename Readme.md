### Setup Postgres:
-  docker run --name postgres -e POSTGRES_PASSWORD=abcd1234 -p 5432:5432 -d postgres
- To test : telnet localhost 5432
-  Go Postgres Library : https://pkg.go.dev/github.com/lib/pq. Go connection steps documented here

### Random learning : 
- %+v shows fields by name in struct output `fmt.Printf("%+v \n", store)`
- %#v formats the struct in Go source format `fmt.Printf("%#v \n", store)`
