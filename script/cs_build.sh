DIR=`dirname $0`

OUTDIR=$DIR/..
PBDIR=$DIR/../proto/msg #项目protobuf文件根目录
KIWIDIR=../../../.. #kiwi.proto存放根目录
PBOUTDIR=$OUTDIR/proto #输出根目录

echo complie cs

protoc \
  --proto_path=$GOPATH/pb/ \
  --proto_path=$KIWIDIR \
  --proto_path=$PBDIR \
  --csharp_out=$PBOUTDIR/cs \
  $PBDIR/model/*.proto $PBDIR/client/*.proto

echo cs finished