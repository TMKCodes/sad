#!/bin/bash
rm -rf /tmp/sad-temp

sad --devnet --appdir=/tmp/sad-temp --profile=6061 --loglevel=debug &
SAD_PID=$!

sleep 1

rpc-stability --devnet -p commands.json --profile=7000
TEST_EXIT_CODE=$?

kill $SAD_PID

wait $SAD_PID
SAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "sad exit code: $SAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SAD_EXIT_CODE -eq 0 ]; then
  echo "rpc-stability test: PASSED"
  exit 0
fi
echo "rpc-stability test: FAILED"
exit 1
