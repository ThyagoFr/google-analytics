package main

import (
	"context"
	"fmt"

	"github.com/joho/godotenv"
	goga "google.golang.org/api/analyticsreporting/v4"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()
	analyticsreportingService, err := goga.NewService(ctx)
	if err != nil {
		panic(err)
	}
	response, err := analyticsreportingService.Reports.BatchGet(&goga.GetReportsRequest{
		ReportRequests: []*goga.ReportRequest{
			{
				Metrics: []*goga.Metric{
					{
						Expression: "ga:sessions",
					},
				},
				Dimensions: []*goga.Dimension{
					{
						Name: "ga:browser",
					},
				},
			},
		},
	}).Do()
	if err != nil {
		panic(err)
	}
	for _, value := range response.Reports {
		fmt.Println(value)
	}
}
