build: ## Build application
	go build

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

stats: ## Run stats
	./go-ml-example stats data.txt

regression: ## Run regression
	./go-ml-example regression cmd/regression/reg_data.txt

plot: ## Run plot
	./go-ml-example plot cmd/regression/reg_data.txt 1 2

classification: ## Run clustering classification
	./go-ml-example cluster cmd/classification/class_data.txt 1

kmeans: ## Run kmeans classification
	./go-ml-example kmeans 5

neural-networks: ## Run neural networks
	./go-ml-example neural

outlier-analysis: ## Run outlier_analysis
	./go-ml-example outlier cmd/outlier_analysis/data.txt 1000

tensorflow: ## Run tensor flow
	./go-ml-example tensorflow 299 47346

anomaly-detection: ## Run anomaly detection
	./go-ml-example anomaly_detection 10

.DEFAULT_GOAL := help
