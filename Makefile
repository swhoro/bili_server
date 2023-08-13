debug_server:
	cd server && go build -gcflags"=all=-N -l" -ldflags="-extldflags --static -X main.listenAddr=127.0.0.1:3000" -o ../build/__debug_server.exe

release_server:
	cd server && go build -trimpath -ldflags="-extldflags --static -s -w -X main.listenAddr=127.0.0.1:4433" -o ../build/server.exe

auto_migrate:
	cd migrate && go build -trimpath -ldflags="-extldflags --static -s -w" -o ../build/auto_migrate.exe

clean:
	cd build && rm *
