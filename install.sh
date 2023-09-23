#!/bin/bash
CMD_NAME="s2s"
echo "building cli binaries" && go build -o $CMD_NAME
chmod 755 $CMD_NAME
echo "moving binaries to /usr/local/bin" && sudo mv $CMD_NAME /usr/local/bin && echo "done"
