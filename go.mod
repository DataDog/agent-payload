module github.com/DataDog/agent-payload/v5

go 1.18

require (
	github.com/DataDog/mmh3 v0.0.0-20200805151601-30884ca2197a
	github.com/DataDog/zstd v1.4.8
	github.com/DataDog/zstd_0 v0.0.0-20210310093942-586c1286621f
	github.com/gogo/protobuf v1.0.0
	github.com/golang/protobuf v1.5.0
	github.com/stretchr/testify v1.3.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
)

retract v5.0.59 // Was accidentally deployed in place of v5.0.49
