gitRev=`git rev-parse --short HEAD`
buildTime=`date +"%Y_%m_%d"`

fileName=gokeeper-$gitRev-$buildTime

go build -ldflags "-X main.gitRev=$gitRev -X main.buildTime=$buildTime" -o ./out/$fileName cmd/gokeeper/main.go

echo $fileName
