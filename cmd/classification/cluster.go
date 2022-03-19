package classification

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/mash/gokmeans"
)

var observations []gokmeans.Node = []gokmeans.Node{
	gokmeans.Node{4},
	gokmeans.Node{5},
	gokmeans.Node{6},
	gokmeans.Node{8},
	gokmeans.Node{10},
	gokmeans.Node{12},
	gokmeans.Node{15},
	gokmeans.Node{0},
	gokmeans.Node{-1},
}

func Kmeans() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("usage: kmeans k")
		return
	}

	k, err := strconv.Atoi(flag.Args()[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	if success, centroids := gokmeans.Train(observations, k, 50); success {
		fmt.Println("The centroids are the following")
		for _, centroid := range centroids {
			fmt.Println(centroid)
		}

		fmt.Println("The cluster are the following")
		for _, observation := range observations {
			index := gokmeans.Nearest(observation, centroids)
			fmt.Println(observation, "belongs in cluster", index+1, ".")
		}
	}
}
