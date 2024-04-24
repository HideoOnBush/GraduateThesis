package spark

import (
	"context"
	"testing"
)

func TestSparkClient_SparkCreateMission(t *testing.T) {
	c := New(Config{
		Host: "http://127.0.0.1:4040",
	})
	result := c.SparkCreateMission(context.TODO(), SubmitMissionReq{
		Action:               "CreateSubmissionRequest",
		AppArgs:              []string{"main.py"},
		AppResource:          "hdfs://namenode:9000/input/main.py",
		ClientSparkVersion:   "2.4.3",
		EnvironmentVariables: EnvironmentVariables{SparkEnvLoaded: "1"},
		MainClass:            "org.apache.spark.deploy.SparkSubmit",
		SparkProperties: SparkProperties{
			SparkDriverSupervise:  "false",
			SparkAppName:          "Simple App",
			SparkEventLogEnabled:  "false",
			SparkSubmitDeployMode: "client",
			SparkMaster:           "spark://master:4040",
			SparkJars:             "hdfs://namenode:9000/input/elasticsearch-hadoop-8.12.2.jar",
			SparkEsResource:       "test_lines",
		},
	})
	t.Log(result)
}

func TestSparkClient_SparkCheckMission(t *testing.T) {
	c := New(Config{
		Host: "http://127.0.0.1:4040",
	})
	result := c.SparkCheckMission(context.TODO(), CheckMissionReq{"driver-20240305182409-0001"})
	t.Log(result)
}
