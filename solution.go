package main
import "fmt"

func main() {
  arr := []int{2,-25,6,8,20,1,-6}
  qs(arr, 0, len(arr)-1)
  fmt.Printf("%v",arr)
}

func qs(arr []int, l int, h int){
    if l < h{
        s := pivot(arr, l, h)
        qs(arr, l, s-1)
        qs(arr, s+1, h)
    }
}

func pivot(arr []int, l int, h int) int{
    j := l
    for i:=l+1 ; i<=h; i++{
        if(arr[i] < arr[l]){
            j++;
            arr[j], arr[i] = arr[i], arr[j]
        }
    }
    arr[l], arr[j] = arr[j], arr[l]
    return j
}
