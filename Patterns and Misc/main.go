package main

import (
	"github.com/RohBot/GOLANG/GoPhase3/painkiller"
	"time"
)

type pnl struct {
	FromDate time.Time `json:"fromDate"`
	ToDate   time.Time `json:"toDate"`
}

func main() {
	painkiller.Acetaminophen.String()
}
