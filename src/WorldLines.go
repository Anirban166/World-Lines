package main

import (
        "fmt"
        "log"
        "time"
        "runtime"
        "math/rand"
)

/* A range-based random number generator to generate a user-defined number of
   float64 divergence numbers associated with the same number of world lines: */
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
    
    var i, worldLines int
    attractorFields := 2 // Alpha, Beta

    // Take number of world lines as user-input, along with possible error:
    _, err := fmt.Scan(&worldLines)
    /* Feel free to take attractor fields (AF) into account as well, if required: 
       _, err := fmt.Scan(&worldLines, &attractorFields) */
    
    // Print error if applicable: (eg: negative units, resulting in panic for size out of range)
    if err != nil {
      log.Println(err)
    }
    
    // Using all available cores, although this program is not compute-intensive:
    runtime.GOMAXPROCS(runtime.NumCPU()) 
    // Feel free to change to a reasonable number within your cpu-core count.
    
    /* Creating a buffered channel (based on AF count) to recieve and 
       store returned values (float64 slice) inside the goroutine: */
    c := make(chan []float64, attractorFields)
    
    // Initialize divergence number limits for an attractor field, based on the standard range:
    divergenceBase, divergenceCap := 0.0, 0.99
    
    /* Initial intent on scripting: Inclusion of all attractor fields, which is what the code below is structured for.
       However after much thought, I realized there wasn't much to do with the other world lines (!alpha/beta)
       so the loop wasn't necessary. (two calls to the generator function or 2 goroutines would have sufficed)
    */
    for i < attractorFields { 
        go func() {
            c <- divergenceNumberGenerator(divergenceBase + float64(i), divergenceCap + float64(i), worldLines)
            i += 1
        }()
   }

   alpha := <-c
   beta := <-c

   fmt.Println("Alpha line:", alpha, "\nBeta line:", beta) 
}