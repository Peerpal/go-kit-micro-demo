#  ./autogen.sh ; ./configure ; make ; make install
#
# Update protoc Go bindings via
#  go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
#
# See also
#  https://github.com/grpc/grpc-go/tree/master/examples

REPO_ROOT="${REPO_ROOT:-$(cd "$(dirname "$0")/../.." && pwd)}"
PB_PATH="${REPO_ROOT}/proto/v1/watermark"
PROTO_FILE=${1:-"watermark.proto"}


echo "Generating pb files for ${PROTO_FILE} service"
protoc -I="${PB_PATH}"  "${PB_PATH}/${PROTO_FILE}" --go_out=plugins=grpc:"${PB_PATH}"