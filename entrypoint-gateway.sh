#!/bin/bash

POD_ID=${HOSTNAME##*-}
START=$((POD_ID * SHARDS_PER_POD))
STOP=$(((POD_ID+1) * SHARDS_PER_POD))

exec /gateway \
	--redis="$REDIS" \
	--etcd="$ETCD" \
	--pprof="$PPROF" \
	--addr="$ADDR" \
	--played="$PLAYED" \
	--shards="$SHARDS" \
	--start="$START" \
	--stop="$STOP"