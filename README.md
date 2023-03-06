# Go Language Take Home Challange

Created by Christopher Tok for the use of Bibit / Stockbit interview test

## Dev environment
- MacOS 11.2 Big Sur
- Apple Silicon (arm64)

## Pre-requisites
- Internet Connection
- Docker

## Some useful command
```
//see inside redis
docker run --network=container:redis -it redis redis-cli -h redis -a testing123

//generate .pb.go from .proto
make pb

//produce manual via command line
docker exec -it redpanda-0 rpk topic produce process_ohlc
```

## Experiences with tech stacks
- Go: ~3 years (used professionally)
- Redis: ~3 years (used professionally and on personal projects)
- Docker: ~3 years (used on personal projects)
- Redpanda: heard about this mq, usually use NSQ (professionally). Just tried it now.

## Architecture of part 1 challange
flow get summary API:
read data from cache by stock name on request
-> if exist, return data
-> if not exist, go to next flow
get data from .ndjson (source of truth)
calculate ohlc
publish message by stockname to process_ohlc topic
return ohlc by stock name on request

flow process_ohlc consumer:
read stock_name
get data from .ndjson (source of truth)
calculate ohlc
store to redis