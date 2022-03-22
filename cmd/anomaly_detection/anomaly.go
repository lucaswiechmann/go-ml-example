package anomaly_detection

import (
	"flag"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/lytics/anomalyzer"
)

func DetectAnomaly() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("usage: anomaly MAX")
		return
	}

	MAX, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		fmt.Println("MAX needs to be an integer")
		return
	}

	conf := &anomalyzer.AnomalyzerConf{
		Sensitivity: 0.1,
		UpperBound:  5,
		LowerBound:  anomalyzer.NA,
		ActiveSize:  1,
		NSeasons:    4,
		Methods:     []string{"diff", "fence", "magnitude", "ks"},
	}

	data := []float64{}
	SEED := time.Now().Unix()
	rand.Seed(SEED)

	for i := 0; i < MAX; i++ {
		data = append(data, float64(rand.Intn(MAX-0)+0))
	}

	fmt.Println("data: ", data)

	anom, _ := anomalyzer.NewAnomalyzer(conf, data)
	prob := anom.Push(8.0)
	fmt.Println("Anomalous Probability:", prob)
}
