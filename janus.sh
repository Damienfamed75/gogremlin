#!/bin/sh

JANUS_VERSION=0.3.0

STOP='stop'
STATUS='status'
LIST='list'
CLEAN='clean'

function realpathMac() {
    [[ $1 = /* ]] && echo "$1" || echo "$PWD/${1#./}"
}

case "$OSTYPE" in
	darwin*) WORKDIR="$(realpathMac `dirname $0`)/assets" ;;
    *) WORKDIR="$(realpath `dirname $0`)/assets" ;;
esac

JDIR=janusgraph-${JANUS_VERSION}-hadoop2
PKGNAME=${JDIR}.zip

JPATH="${WORKDIR}/${JDIR}"

case $1 in
  $STOP)
    cd "$JPATH" && ./bin/janusgraph.sh stop
    exit 1
    ;;
  $STATUS)
    cd "$JPATH" && ./bin/janusgraph.sh status
    exit 1
    ;;
  $CLEAN)
    cd "$WORKDIR" && rm -rf $PKGNAME && rm -rf $JDIR
    echo 'Removed zip and directory for Janusgraph'
    exit 1
    ;;
  $LIST)
    echo $STOP
    echo $STATUS
    echo $CLEAN
    echo $LIST
    exit 1
    ;;
esac

function janus_exists(){
  test -d "${JPATH}"
}

function download_janus(){
  echo Downloading janus
  mkdir -p $WORKDIR
  cd "$WORKDIR"
  case "$OSTYPE" in
  darwin*)  curl -LO https://github.com/JanusGraph/janusgraph/releases/download/v${JANUS_VERSION}/${PKGNAME} ;; 
  *)        wget -c https://github.com/JanusGraph/janusgraph/releases/download/v${JANUS_VERSION}/${PKGNAME} ;;
  esac
  echo Extracting
  unzip -o -q $PKGNAME
}


janus_exists || download_janus

echo starting janus
# make sure it is not running first
cd "$JPATH" && ./bin/janusgraph.sh stop
cd "$JPATH" && ./bin/janusgraph.sh start

echo "JanusGraph Begun!"