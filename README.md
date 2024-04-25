# Tubes2_BE_Penyelam_Handal

## How to Run / Build

To build the project, execute
```
cd/src
go mod tidy
go get -u github.com/gin-gonic/gin
```

To run the project, execute
```
go run .
```

## API format
```
localhost:9090/{algorithm}/{results}?source="{source_title}"&goal"{goal_title}"
```
1. algorithm = ```bfs / ids```
2. results = ```single/many```