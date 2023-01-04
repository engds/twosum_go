package main

import "fmt"

func main() {
    fmt.Println("Two Sum!")

    myArray := []int{4,3,9,2,11,8,17}
    myTargetValue := 17


    //Get the indices of the first two values in the array that add up to the given target value
    result := GetTwoSum(myArray, myTargetValue)


    for _, value := range result{
        fmt.Printf(" %d", value)
    }

    if (len(result) == 0){
        fmt.Printf("No matches")
    }

}

func GetTwoSum(arr []int, target int) []int{
    var m map[int]int
    //need to initialize map

    m=make(map[int]int)

    for i:=0; i<len(arr); i++ {
        y := target-arr[i]
        z, ok := m[y]

        if(ok){
            return []int{z,i}
        }

        m[arr[i]] = i
    }

   return nil
}