


ImportPath=${HOME}/projects/zwczou/rubikpb/proto
OutputPath=${HOME}/projects/zwczou/rubikpb/proto/rbkpb

protobuf:
	@mkdir -p proto/rbkpb | true
	protoc -I/usr/local/include -I. \
		-I${ImportPath} \
		-I${ImportPath}/include \
		--go_out=paths=source_relative:${OutputPath} \
		--go-grpc_out=${OutputPath}  --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
		--grpc-gateway_out=logtostderr=true:${OutputPath} --grpc-gateway_opt=paths=source_relative \
		--validate_out=lang=go:${OutputPath} --validate_opt=paths=source_relative \
		--openapiv2_out=logtostderr=true:${OutputPath} \
		enum.proto \
		model.proto \
		tick.proto
