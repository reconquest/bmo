#!/bin/bash

set -e -u

SRCURL="github.com/reconquest/bmo"
PKGDIR="bmo-deb"
SRCROOT="src"
SRCDIR=${SRCROOT}/${SRCURL}

mkdir -p $PKGDIR/etc/bmo
mkdir -p $PKGDIR/usr/bin
rm -rf $SRCROOT

export GOPATH=`pwd`
go get -v $SRCURL
pushd $SRCDIR

count=$(git rev-list HEAD| wc -l)
commit=$(git rev-parse --short HEAD)
VERSION="${count}.$commit"

go build -o bmo -ldflags "-X main.version=$VERSION" ./

popd

sed -i 's/\$VERSION\$/'$VERSION'/g' $PKGDIR/DEBIAN/control

cp -f bin/bmo $PKGDIR/usr/bin/bmo

dpkg -b $PKGDIR bmo-${VERSION}_amd64.deb

# restore version placeholder
git checkout $PKGDIR
