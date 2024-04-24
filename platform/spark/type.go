package spark

type EnvironmentVariables struct {
	SparkEnvLoaded string `json:"SPARK_ENV_LOADED"`
}

type SparkProperties struct {
	SparkDriverSupervise  string `json:"spark.driver.supervise"`
	SparkAppName          string `json:"spark.app.name"`
	SparkEventLogEnabled  string `json:"spark.eventLog.enabled"`
	SparkSubmitDeployMode string `json:"spark.submit.deployMode"`
	SparkMaster           string `json:"spark.master"`
	SparkJars             string `json:"spark.jars"`
	SparkEsResource       string `json:"spark.es.resource"`
}

type SubmitMissionReq struct {
	Action               string               `json:"action"`
	AppArgs              []string             `json:"appArgs"`
	AppResource          string               `json:"appResource"`
	ClientSparkVersion   string               `json:"clientSparkVersion"`
	EnvironmentVariables EnvironmentVariables `json:"environmentVariables"`
	MainClass            string               `json:"mainClass"`
	SparkProperties      SparkProperties      `json:"sparkProperties"`
}

type CheckMissionReq struct {
	SubmissionId string `json:"submissionId"`
}

type SubmitMissionResp struct {
	Action             string `json:"action"`
	Message            string `json:"message"`
	ServerSparkVersion string `json:"serverSparkVersion"`
	SubmissionId       string `json:"submissionId"`
	Success            bool   `json:"success"`
}

type CheckMissionResp struct {
	Action             string `json:"action"`
	DriverState        string `json:"driverState"`
	ServerSparkVersion string `json:"serverSparkVersion"`
	SubmissionId       string `json:"submissionId"`
	Success            bool   `json:"success"`
	WorkerHostPort     string `json:"workerHostPort"`
	WorkerId           string `json:"workerId"`
}

func (q SubmitMissionResp) getUri() string {
	return SparkCreateMissionUri
}

func (q SubmitMissionResp) getData() interface{} {
	return q
}

func (q SubmitMissionResp) getName() string {
	return "SubmitMission"
}

func (q CheckMissionResp) getUri() string {
	return SparkCheckMissionUri
}

func (q CheckMissionResp) getData() interface{} {
	return q
}

func (q CheckMissionResp) getName() string {
	return "CheckMission"
}
