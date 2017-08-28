package main

var arr [][]uint64

func main() {
    arr = make([][]uint64, 21)
    for i := range arr {
        arr[i] = make([]uint64, 21)
        for j := range arr[i] {
            // sets every coordicnate to 0 - clearing the board
            arr[i][j] = uint64(0)
        }
    }
    // sets the bottom right coordinate, we have to visit it eventually!
    arr[20][20] = uint64(1)
    println(route(0, 0))
}

// i is left
// j is down
func route(i, j int) uint64 {
  // if we visited there, or if we are at the bottom right corner
    if arr[i][j] != uint64(0) {
        return arr[i][j]
    }
    // returns values based on how far down and right
    var result uint64
    // if neither at the right or bottom edge - move down and to the right
    if i < 20 && j < 20 {
        result = noOfRoutes(i+1, j) + noOfRoutes(i, j+1)
    // if at the bottom edge - move right
    } else if i < 20 && j == 20 {
        result = noOfRoutes(i+1, j)
    // if at the right edge - move down
    } else {
        result = noOfRoutes(i, j+1)
    }
    // removes the uint64 0 - marks we visited this intersection before
    arr[i][j] = result
    return result
}
