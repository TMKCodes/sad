#!/bin/bash

APPDIR=/tmp/sad-temp
SAD_RPC_PORT=29587

rm -rf "${APPDIR}"

sad --simnet --appdir="${APPDIR}" --rpclisten=0.0.0.0:"${SAD_RPC_PORT}" --profile=6061 &
SAD_PID=$!

sleep 1

RUN_STABILITY_TESTS=true go test ../ -v -timeout 86400s -- --rpc-address=127.0.0.1:"${SAD_RPC_PORT}" --profile=7000
TEST_EXIT_CODE=$?

kill $SAD_PID

wait $SAD_PID
SAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "sad exit code: $SAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SAD_EXIT_CODE -eq 0 ]; then
  echo "mempool-limits test: PASSED"
  exit 0
fi
echo "mempool-limits test: FAILED"
exit 1
