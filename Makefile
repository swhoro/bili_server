debug_server:
	cd server && go build -gcflags"=all=-N -l" -ldflags="-extldflags --static -X main.listenAddr=127.0.0.1:3000" -o ../build/__debug_server

release_server:
	cd server && go build -trimpath -ldflags="-extldflags --static -s -w -X main.listenAddr=127.0.0.1:4433" -o ../build/server

auto_migrate:
	cd migrate && go build -trimpath -ldflags="-extldflags --static -s -w" -o ../build/auto_migrate

server_upx: release_server
	cd build && upx --best -o server_upx server

clean:
	cd build && rm *
