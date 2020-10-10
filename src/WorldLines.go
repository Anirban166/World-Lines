package main

import (
	  "fmt"
    "math/rand"
    "time"
    "log"
    //"sync"
    "runtime"
)

/* Basically a range-based random number generator to generate a user-defined number 
   of float64 divergence numbers associated with the same number of world lines:  */
func divergenceNumberGenerator(lowerLimit, upperLimit float64, num int) []float64 {

    result := make([]float64, num)

    for i := range result {
        result[i] = lowerLimit + rand.Float64() * (upperLimit - lowerLimit)
    }
    
    // Setting seed based on time to avoid storing the same results time and again:
    rand.Seed(time.Now().UnixNano())

    return result
}

func main() { 
    
    var (i, worldLines, attractorFields int
         //wg sync.WaitGroup
    )

    _, err := fmt.Scan(&worldLines, &attractorFields)
    
    // Print error if applicable: (eg: negative units, resulting in panic for size out of range)
    if err != nil {
      log.Println(err)
    }
    
    // Using all available cores, although this program is not compute-intensive:
    runtime.GOMAXPROCS(runtime.NumCPU()) 
    // Feel free to change to a reasonable number within your cpu-core count.

    //wg.Add(attractorFields) 
    
    /* Creating a buffered channel (based on AF count) to recieve and 
       store returned values (float64 array) inside the goroutine: */
    c := make(chan []float64, attractorFields)
    
    divergenceBase, divergenceCap := 0.0, 0.99

    for i < attractorFields { 
        go func() {
            //defer wg.Done()
            c <- divergenceNumberGenerator(divergenceBase + float64(i), divergenceCap + float64(i), worldLines)
            i += 1
        }()
   }

   alpha := <-c
   beta := <-c

   fmt.Println("Alpha line:", alpha, "\nBeta line:", beta)
   //wg.Wait()
   fmt.Println("End") 
}