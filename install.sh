#!/usr/bin/env bash
  # build bin
    if [[ ! -f "resource-access-control" ]]; then
        ./build.sh "bin-ved"
    fi
    destPath=${1}
    if [[ "${destPath}" == "" ]]; then
       echo "${destPath} cannot be empty"
       exit
    else
       if [[ ! -d "${destPath}" ]]; then
            echo "copy dest path not found..."
            exit
       fi
    fi
  # copy bin、config
    \cp -r resource-access-control swagger  ${destPath}/
