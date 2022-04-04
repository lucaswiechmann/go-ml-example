package main

import (
	"flag"
	"fmt"
	"strings"

	"go-ml-example/cmd/anomaly_detection"
	"go-ml-example/cmd/classification"
	"go-ml-example/cmd/neural_networks"
	"go-ml-example/cmd/outlier_analysis"
	"go-ml-example/cmd/plot"
	"go-ml-example/cmd/regression"
	"go-ml-example/cmd/stats"
	"go-ml-example/cmd/tensorflow"
)

func main() {
	flag.Parse()
	if stop := checkArgs(); stop {
		return
	}

	if strings.ToUpper(flag.Args()[0]) == "STATS" {
		stats.CalculateStats()
	} else if strings.ToUpper(flag.Args()[0]) == "REGRESSION" {
		regression.Regression()
	} else if strings.ToUpper(flag.Args()[0]) == "PLOT" {
		plot.PlotLR()
	} else if strings.ToUpper(flag.Args()[0]) == "CLUSTER" {
		classification.ClassifyByKNNMethod()
	} else if strings.ToUpper(flag.Args()[0]) == "KMEANS" {
		classification.Kmeans()
	} else if strings.ToUpper(flag.Args()[0]) == "ANOMALY_DETECTION" {
		anomaly_detection.DetectAnomaly()
	} else if strings.ToUpper(flag.Args()[0]) == "NEURAL" {
		neural_networks.TrainNeural()
	} else if strings.ToUpper(flag.Args()[0]) == "OUTLIER" {
		outlier_analysis.CalcOutlier()
	} else if strings.ToUpper(flag.Args()[0]) == "TENSORFLOW" {
		tensorflow.RunTensorFlow()
	}
}

func checkArgs() bool {
	if len(flag.Args()) <= 1 {
		if strings.ToUpper(flag.Args()[0]) == "NEURAL" {
			return false
		}
		fmt.Println("")
		fmt.Println("Usage:")
		fmt.Println("make <OPTION>")
		fmt.Println("")
		fmt.Println("To check options available, run 'make help'")
		return true
	}
	return false
}
