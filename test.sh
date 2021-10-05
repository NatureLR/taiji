make build
make install
mkdir test
cd test
taiji  init --mod=github.com/naturelr/test 
go mod tidy
go mod vendor
git add .
git commit -m "xxx"
make all