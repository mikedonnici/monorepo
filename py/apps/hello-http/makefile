

pb:
	python -m grpc_tools.protoc \
	--proto_path=../../../proto \
	--python_out=. \
	--grpc_python_out=. \
	hello.proto

run:
	uvicorn main:app --reload
