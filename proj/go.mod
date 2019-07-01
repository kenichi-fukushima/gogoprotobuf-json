module example.com/proj

require (
	example.com/api v0.0.0
	github.com/gogo/protobuf v1.2.1
	github.com/golang/protobuf v1.3.1
)

replace example.com/api => ../lib/example.com/api
