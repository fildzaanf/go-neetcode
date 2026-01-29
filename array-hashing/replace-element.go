package arrayhashing

func ReplaceElements(arr []int) []int {
    newArr := make([]int, len(arr))
    rightMax := -1
    for i := len(arr) - 1; i >= 0; i-- { 
        newArr[i] = rightMax 
        if arr[i] > rightMax { 
            rightMax = arr[i] 
        }
    }
    return newArr
}