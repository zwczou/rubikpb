


ImportPath=${HOME}/projects/zwczou/rubikpb/proto
OutputPath=${HOME}/projects/zwczou/rubikpb/proto

protobuf:
	@mkdir -p proto/rbkpb | true
	protoc -I/usr/local/include -I. \
		-I${ImportPath} \
		-I${ImportPath}/include \
		--go_out=paths=source_relative:${OutputPath}/rbkpb \
		--go-grpc_out=${OutputPath}/rbkpb  --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
		--grpc-gateway_out=logtostderr=true:${OutputPath}/rbkpb --grpc-gateway_opt=paths=source_relative \
		--validate_out=lang=go:${OutputPath}/rbkpb --validate_opt=paths=source_relative \
		--openapiv2_out=logtostderr=true:${OutputPath}/rbkpb \
		enum.proto \
		model.proto

	@mkdir -p proto/tickpb | true
	protoc -I/usr/local/include -I. \
		-I${ImportPath} \
		-I${ImportPath}/include \
		--go_out=paths=source_relative:${OutputPath}/tickpb \
		--go-grpc_out=${OutputPath}/tickpb  --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false \
		--grpc-gateway_out=logtostderr=true:${OutputPath}/tickpb --grpc-gateway_opt=paths=source_relative \
		--validate_out=lang=go:${OutputPath}/tickpb --validate_opt=paths=source_relative \
		--openapiv2_out=logtostderr=true:${OutputPath}/tickpb \
		tick.proto
