// A goroutine with job pool/workers pattern, finding prime numbers

package main

import (
	"fmt"
	"math"
)

type JobRange struct {
	From int
	To   int
}

type Job struct {
	JobID int
	JobRange
}

type Result struct {
	JobId int
	prime []int
}

// take a job from pool, calcaulte prime numbers,
// put them in to Result
func findPrime(w int, jobPool chan Job, resultPool chan Result) {
	for job := range jobPool {
		result := Result{job.JobID, make([]int, 0)}
		for i := job.From; i < job.To; i++ {
			isPrime := true
			for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
				if i%j == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				result.prime = append(result.prime, i)
			}
		}
		resultPool <- result
	}
}

func main() {
	// five jobs, two workes, five results
	jobPool := make(chan Job, 5)
	resultPool := make(chan Result, 5)

	jobPool <- Job{0, JobRange{2, 10}}
	jobPool <- Job{1, JobRange{10, 20}}
	jobPool <- Job{2, JobRange{20, 30}}

	for i := 0; i < 2; i++ {
		go findPrime(i, jobPool, resultPool)
	}

	jobPool <- Job{3, JobRange{30, 40}}
	jobPool <- Job{4, JobRange{40, 50}}
	close(jobPool)

	for i := 0; i < 5; i++ {
		result := <-resultPool
		fmt.Println(result)
	}
}
/*
$ go run worker.go
{0 [2 3 5 7]}
{1 [11 13 17 19]}
{2 [23 29]}
{3 [31 37]}
{4 [41 43 47]}

*/
