#!/bin/sh

/app/auto-migrate

if [ "$?" -ne 0 ]; then 
    echo "auto migrate failed"; 
    exit 1;
fi

/app/server-upx