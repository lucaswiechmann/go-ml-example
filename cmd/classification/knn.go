package classification

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/evaluation"
	"github.com/sjwhitworth/golearn/knn"
)

func ClassifyByKNNMethod() {
	flag.Parse()
	if len(flag.Args()) < 2 {
		fmt.Println("usage: classification filename k")
		return
	}

	dataset := flag.Args()[1]
	rawData, err := base.ParseCSVToInstances(dataset, false)
	if err != nil {
		fmt.Println(err)
		return
	}

	k, err := strconv.Atoi(flag.Args()[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	cls := knn.NewKnnClassifier("euclidean", "linear", k)

	train, test := base.InstancesTrainTestSplit(rawData, 0.50)
	cls.Fit(train)

	p, err := cls.Predict(test)
	if err != nil {
		fmt.Println(err)
		return
	}

	confusionMat, err := evaluation.GetConfusionMatrix(test, p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(evaluation.GetSummary(confusionMat))
}
