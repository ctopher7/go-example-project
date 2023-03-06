# Go Language Take Home Challange

Created by Christopher Tok for the use of Bibit / Stockbit interview test

## Dev environment
- MacOS 11.2 Big Sur
- Apple Silicon (arm64)

## Pre-requisites
- Internet Connection (for downloading docker images)
- Docker
- Make
- GRPC client (bloomRPC, grpcox, kreya, etc.)
- golangci-lint

## How to use
### part1
- run ``` docker-compose up -d ```
- connect to ``` localhost:9000 ``` using your GRPC client
- try make a request to getSummary endpoint
### part2
- run ```go run main.go part2``` to get result
- run ```golangci-lint run ./part2/``` to see golangci lint result

## Architecture of part 1
repo architecture:
- Config (to be injected to any layer / resource initialization. consist of configurations)
- Resources (to be injected to repository layer, usually client for other dependency like redis and kafka)
- Handler (layer for serialization and deserialization, call usecase)
- Usecase (layer for business logic, consists of datalogic and repository)
- Datalogic (layer for data logic, consists of reposistory)
- Repository (wrapper for other library, no unit test because dependencies not mockable)

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
calculate ohlc (only the data of the stock_name, not all data)
store to redis (expiry 10 seconds, can be extended)

## part 2 comments
- in ohlc function, there are 2 loops (looping records and indexMembers) inside a loop (stock codes). it is inefficient. refactored to only looping 2 loop (looping records and indexMembers, auto assign result to map) 
- use early return / continue instead of if / else if / else to prevent the code become very nested (more of my preference, but for me, less nested code block will have better readability)
- rename variable. x y z will confuse other people when seeing the code. better to use noun that describe the variable
- wrong summary.LowPrice calculation (always zero), stock price unlikely to be zero, so I added OR logic when summary.LowPrice is 0, auto assign low price
 
## Some useful command
```
//see inside redis
make redis-cli

//generate .pb.go from .proto
make pb

//produce manual via command line
make producer

//generate mock
make mockery
```

## Experiences with tech stacks
- Go: ~3 years (used professionally)
- Redis: ~3 years (used professionally and on personal projects)
- Docker: ~3 years (used on personal projects)
- Redpanda: heard about this mq, usually use NSQ (professionally). Just tried it now.