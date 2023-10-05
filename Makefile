debug-server:
	cd server && CC=musl-gcc go build -gcflags="all=-N -l" -ldflags='-linkmode "external" -extldflags "-static" -X main.listenAddr=127.0.0.1:3000' -o ../build/__debug_server

release-server-musl:
	cd server && CC=musl-gcc go build -trimpath -ldflags='-linkmode "external" -extldflags "-static" -s -w -X main.listenAddr=127.0.0.1:4433' -o ../build/server-musl

auto-migrate:
	cd migrate && CC=musl-gcc go build -trimpath -ldflags='-linkmode "external" -extldflags "-static" -s -w' -o ../build/auto-migrate

server-upx: release-server-musl
	cd build && rm server-upx && upx --best -o server-upx server-musl

server-image:
	docker build -t local/bili-server:latest .

clean:
	cd build && rm *
