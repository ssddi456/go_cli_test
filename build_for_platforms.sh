#!/bin/sh
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
DIST=$DIR/dist

test=false

function doBuild {
    echo build start
    rm -rf $DIST
    go build -o $DIST/default/go_cli_test main.go
    CGO_ENABLED=0 GOOS=windows  GOARCH=amd64  go  build -o $DIST/windows/go_cli_test.exe main.go
    CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64  go  build -o $DIST/mac/go_cli_test main.go
    echo build finish
}

function doTest {
    echo test start
    $DIST/default/go_cli_test
    echo test finish
}

function usage {
    echo "build and test script for this little tool

    -b --build do build only
    -t --test do build and test
    -h --help print this help info
"
}

while [ "$1" != "" ]; do
    case $1 in
        -b | --build )          doBuild
                                exit
                                ;;
        -t | --test )           doBuild
                                doTest
                                exit
                                ;;
        -h | --help )           usage
                                exit
                                ;;
        * )                     usage
                                exit 1
    esac
    shift
done
usage
exit
