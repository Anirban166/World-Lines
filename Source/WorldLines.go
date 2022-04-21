package main

import (
        "fmt"
        "log"
        "time"
        "math"
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

    return result
}

func main() { 
    
    // Setting seed based on time to avoid storing the same results time and again:
    rand.Seed(time.Now().UnixNano())

    var (worldLines int
         alpha, beta []float64)
    attractorFields := 2 // Considering alpha & beta fields for the moment

    // Taking number of world lines as user-input, along with possible error:
    fmt.Print("Enter the number of world lines to create: ")
    _, err := fmt.Scan(&worldLines)
    /* Feel free to take attractor fields (AF) into account as well, if required: 
       _, err := fmt.Scan(&worldLines, &attractorFields) */
    
    // Printing error if applicable: (eg: negative units, resulting in panic for size out of range)
    if err != nil {
      log.Println(err)
    }
    
    // Using all available cores, although this program is not compute-intensive:
    runtime.GOMAXPROCS(runtime.NumCPU()) 
    // Feel free to change to a reasonable number within your cpu-core count.
    
    /* Creating a buffered channel (based on AF count) to recieve and 
       store returned values (float64 slice) inside the goroutine: */
    c := make(chan []float64, attractorFields)
    
    // Initializing limits for the divergence numbers based on the number of attractor fields alotted:
    divergenceBase, divergenceCap := 0.0, (0.99 + float64(attractorFields) - 1.0)
    // Initially while scripting this, the inclusion of all attractor fields seemed suitable, which is what the code above is structured for.
    
    // Collect/Recieve the divergence numbers of the world lines in a bidirectional channel:
    go func() {
        c <- divergenceNumberGenerator(divergenceBase, divergenceCap, worldLines)
    }()
    
    // Send it from the channel to a variable: ([]float64)
    divergenceNumbers := <-c
    total := 0.0

    // Collect the numbers on new slices based on corresponding attractor field ranges:
    for _, divergenceNumber := range divergenceNumbers {
        total += divergenceNumber  // try using defer
        if divergenceNumber <= 0.99 && divergenceNumber >= 0.0 {
            alpha = append(alpha, divergenceNumber)  
        } else {
          beta = append(beta, divergenceNumber)
        }
    }

    average := total / float64(len(divergenceNumbers))
    relativeDivergence := math.Abs(average - 1.048596)
    fmt.Println("Alpha lines:", alpha, "\nBeta lines:", beta) 
    fmt.Print("Divergence(%): ")
    fmt.Printf("%.2f\n", relativeDivergence)
}