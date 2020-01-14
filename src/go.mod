module aoko

go 1.12

// github link latest
// for example: github.com/pkg/sftp latest
// go clean -modcache 清除缓存
// go mod vendor 自动创建vendor 目录

require (
	github.com/axgle/mahonia v0.0.0-20180208002826-3358181d7394
	github.com/garyburd/redigo v1.6.0 // indirect
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/golang/protobuf v1.3.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/gonutz/binpacker v0.0.0-20170423181552-aa2d50f5ae3a // indirect
	github.com/gonutz/ide v0.0.0-20180502124734-e9fc8c14ed56
	github.com/gonutz/truetype v0.0.0-20170213141940-8a63cd5e8201 // indirect
	github.com/gorilla/websocket v1.4.1
	github.com/kr/pretty v0.1.0 // indirect
	github.com/mattn/go-sqlite3 v2.0.2+incompatible
	github.com/satori/go.uuid v1.2.0
	github.com/stvp/tempredis v0.0.0-20181119212430-b82af8480203 // indirect
	golang.org/x/sys v0.0.0-20191228213918-04cbcbbfeed8
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/redsync.v1 v1.0.1
	gopkg.in/tomb.v2 v2.0.0-20161208151619-d5d1b5820637 // indirect
)

replace (
	golang.org/x/arch => github.com/golang/arch v0.0.0-20190312162104-788fe5ffcd8c
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190701094942-4def268fd1a4
	golang.org/x/net => github.com/golang/net v0.0.0-20190724013045-ca1201d0de80
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190804053845-51ab0e2deafa
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190806215303-88ddfcebc769
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20190717185122-a985d3407aa7
)
