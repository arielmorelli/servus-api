#!/bin/sh

FILE_ARG=""
if [[ ! -z $ROUTE_FILE ]]
then
	FILE_ARG="--file $ROUTE_FILE"
fi

MODE_ARG=""
if [[ "$MODE" == "debug" ]]
then
	MODE_ARG="--debug"
fi

exec ./servus-api --port $PORT $MODE_ARG $FILE_ARG
