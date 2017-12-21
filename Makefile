# grpc, protobuf, consul example



gen: gen-proto

gen-proto:
	protoc -I ./pb --go_out=plugins=grpc:./pb ./pb/pb.proto