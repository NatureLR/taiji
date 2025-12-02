rm -rf test
mkdir test
cd test
taiji  init --mod=github.com/naturelr/test 
go mod tidy
go mod vendor
git add .
git commit -m "测试"
git tag -a "v1.0.0" -m "test"
make all