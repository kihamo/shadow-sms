#!/usr/bin/env bash

BASE_PATH=$PWD

# versions
CMP_SMS="2.0."

for CMP in `find components -maxdepth 1 -type d ! -path components`
do
    cd $BASE_PATH/$CMP

    CMP_NAME=`basename $PWD | tr "[:lower:]" "[:upper:]"`
    CMP_VERSION="1.0."

    CMP_VAR="CMP_${CMP_NAME}"
    if [ -n "${!CMP_VAR}" ]; then
        CMP_VERSION=${!CMP_VAR}
    fi

    CMP_BUILD_NUMBER=`git log component.go | wc -l`
    CMP_VERSION="${CMP_VERSION}${CMP_BUILD_NUMBER}"
    CMP_PACKAGE=`go list -e -f '{{.Name}}' ./`

cat << EOF > version.go
package ${CMP_PACKAGE}

const (
	ComponentVersion = "${CMP_VERSION}"
)
EOF
done

# formatting
cd $BASE_PATH
goimports -w ./

cd $BASE_PATH/components/sms && go-bindata-assetfs -pkg=sms templates/...