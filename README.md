# 🏊🏻‍♀️ Tubes2_BE_Penyelam_Handal
<iframe src="https://giphy.com/embed/l41lFw057lAJQMwg0" width="480" height="270" frameBorder="0" class="giphy-embed" allowFullScreen></iframe><p><a href="https://giphy.com/gifs/rickandmorty-rick-and-morty-210-l41lFw057lAJQMwg0">via GIPHY</a></p>
This repository provides an API for solving [**WikiRace**](https://wiki-race.com/) in the shortest order as possible. Inspired by [**Six Degrees of Wikipedia**](https://www.sixdegreesofwikipedia.com/), we implemented ***Breadth First Search*** and ***Iterative Depth Search*** using GO Language.

## 🔨 How to Run / Build
To build the project, execute
```
cd/src
go mod tidy
go get -u github.com/gin-gonic/gin
go get github.com/gin-contrib/cors
go get github.com/PuerkitoBio/goquery
```

To run the project, execute
```
cd/src
go run .
```

## 📋 API format
```
localhost:9090/{bfs | ids}/{single | many}?source={source_title}&goal={goal_title}
```
1. ```bfs | ids ``` sets the mode for the searching algorithm
2. ```single | many``` sets the amount of results
3. ``` source_title && goal_title``` is the ```https://en.wikipedia.org/wiki/{HERE!}``` part of a valid wikipedia URL.

### Usage Example
Search from Spain to Classical Guitar :
```
localhost:9090/ids/single?source=Spain&goal=Classical_guitar
```
Search from Chicken to Football : 
```
localhost:9090/ids/single?source=CHicken&goal=Football
```
