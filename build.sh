#!/bin/bash
source ./env
arch="`uname -s`"
moduleMode="mod"
vendorMode="ved"
set -e

buildBin() {
    echo "begin build resource-access-control...."
    funcPrefix=""
    if [[ $1 == ${moduleMode} ]]; then
        make build
    else
        make prod
    fi
    echo "end build resource-access-control....."
}

buildImage() {
    echo "begin build image $version....."
    if [ ! -f "resource-access-control" ]; then
        buildBin $1
    fi
    docker build -t harbor.tusdao.com/switch/dev/resource-access-control:${version} .
    echo "end build image....."
}

printHelp() {
    echo "Usage: "
    echo "  build.sh <mode> [-v <image version>]"
    echo "    <mode> - one of 'all-${moduleMode}', 'all-${vendorMode}', 'bin-${moduleMode}', 'bin-${vendorMode}', 'image'"
    echo "      - 'all-${moduleMode}' - build with mod 、image "
    echo "      - 'all-${vendorMode}' - build with vendor 、image "
    echo "      - 'bin-${moduleMode}' - build with mod"
    echo "      - 'bin-${vendorMode}' - build with vendor"
    echo "      - 'image' - create docker image"
    echo
    echo "    -v - set docker image verison"
}

MODE=$1
shift
while getopts "h?v:" opt; do
    case "$opt" in
        h | \?)
            printHelp
            exit 0
        ;;
        v)
            version=$OPTARG
    esac
done

if [ "${MODE}" == "all-${moduleMode}" ]; then
    buildBin ${moduleMode}
# buildImage
elif [ "${MODE}" == "all-${vendorMode}" ]; then
    ## Clear the network
    buildBin ${vendorMode}
    buildImage
elif [ "${MODE}" == "bin-${moduleMode}" ]; then
    ## Clear the network
    buildBin ${moduleMode}
elif [ "${MODE}" == "bin-${vendorMode}" ]; then
    ## Clear the network
    buildBin ${vendorMode}
elif [ "${MODE}" == "image-${moduleMode}" ]; then
    ## Generate Artifacts
    buildImage ${moduleMode}
elif [ "${MODE}" == "image-${vendorMode}" ]; then
    ## Generate Artifacts
    buildImage ${vendorMode}
else
    printHelp
    exit 1
fi

