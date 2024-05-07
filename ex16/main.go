package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())

    N := 15
    array := make([]int, N)

    for i := 0; i < N; i++ {
        array[i] = rand.Intn(33) - 16 // от -16 до 16
    }

    fmt.Println(array)

    quickSort(array)

    fmt.Println(array)
}

func quickSort(array []int) {
    quickSortMain(array, 0, len(array)-1)
}

func quickSortMain(array []int, left int, right int) {
    if left > right {
        return
    }
    aver := array[(left+right)/2]

    i, j := left, right

    for i <= j {
        for array[i] < aver {
            i++
        }
        for array[j] > aver {
            j--
        }
        if i <= j {
            array[i], array[j] = array[j], array[i]
            i++
            j--
        }
    }

    quickSortMain(array, left, j)
    quickSortMain(array, i, right)
}
