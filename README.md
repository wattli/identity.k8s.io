# Secure Kubernetes Identity

```
$ gazelle -mode=fix -build_file_name=BUILD -external=vendored -repo_root=$(pwd) -proto=legacy
$ ./protoc --go_out=plugins=grpc:. pkg/apis/idmgr/idmgr.proto
```
