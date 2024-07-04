#!/bin/bash
rm -rf /tmp/sad-temp

sad --simnet --appdir=/tmp/sad-temp --profile=6061 &
SAD_PID=$!

sleep 1

orphans --simnet -alocalhost:22511 -n20 --profile=7000
TEST_EXIT_CODE=$?

kill $SAD_PID

wait $SAD_PID
SAD_EXIT_CODE=$?

echo "Exit code: $TEST_EXIT_CODE"
echo "sad exit code: $SAD_EXIT_CODE"

if [ $TEST_EXIT_CODE -eq 0 ] && [ $SAD_EXIT_CODE -eq 0 ]; then
  echo "orphans test: PASSED"
  exit 0
fi
echo "orphans test: FAILED"
exit 1
