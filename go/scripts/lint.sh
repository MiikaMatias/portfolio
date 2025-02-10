#/bin/sh

for f in $(find ./ -name '*.go')
do
    gofmt -s -w $f
    go vet $f
done

