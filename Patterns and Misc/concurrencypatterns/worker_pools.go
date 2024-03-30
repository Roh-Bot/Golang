package concurrencypatterns

import "math/rand/v2"

type Calculator interface {
	Calculate() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

type Circle struct {
	Radius float64
}

func (r *Rectangle) Calculate() float64 {
	return r.Length * r.Width
}

func (c *Circle) Calculate() float64 {
	return 3.14 * c.Radius * c.Radius
}

func worker(jobs <-chan Calculator, results chan<- float64) {
	for j := range jobs {
		results <- j.Calculate()
	}
}

func WokerPools() {
	const numJobs = 100
	jobs := make(chan Calculator, numJobs)
	results := make(chan float64, numJobs)
	for i := 0; i <= 5; i++ {
		go worker(jobs, results)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- &Rectangle{Length: rand.Float64(), Width: rand.Float64()}
		jobs <- &Circle{Radius: rand.Float64()}
	}
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		println(<-results)
		println(<-results)
	}
	close(results)
}
