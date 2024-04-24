package spark

import (
	"GraduateThesis/pkg/utils"
	"context"
	"github.com/go-resty/resty/v2"
	"log"
	"net/http"
)

type SparkClient struct {
	client *resty.Client
}

type Config struct {
	Host string
}

type RespType interface {
	getUri() string
	getData() interface{}
	getName() string
}

const (
	SparkCreateMissionUri = "v1/submissions/create"
	SparkCheckMissionUri  = "v1/submissions/status/{submissionId}"
	DriverSuccessStatus   = "FINISHED"
	DriverRunningStatus   = "RUNNING"
	DriverFailedStatus    = "FAILED"
	SparkIndex            = "spark_result"
)

func New(c Config) (m *SparkClient) {
	h := map[string]string{
		"Content-Type": "application/json",
		"charset":      "UTF-8",
	}
	client := resty.New().SetBaseURL(c.Host).SetHeaders(h)
	m = &SparkClient{
		client: client,
	}
	return m
}

func (c *SparkClient) SparkCreateMission(ctx context.Context, req SubmitMissionReq) *SubmitMissionResp {
	data := post[SubmitMissionResp](c, req)
	result, ok := data.(SubmitMissionResp)
	if ok {
		return &result
	}
	return nil
}

func (c *SparkClient) SparkCheckMission(ctx context.Context, req CheckMissionReq) *CheckMissionResp {
	data := get[CheckMissionResp](c, req)
	result, ok := data.(CheckMissionResp)
	if ok {
		return &result
	}
	return nil
}

func post[T RespType](c *SparkClient, req interface{}) interface{} {
	var data T
	// 构建请求
	uri := data.getUri()
	request := c.client.NewRequest().SetContext(context.WithValue(context.Background(), "Method-Name", uri))
	response, err := request.SetBody(req).SetResult(&data).Post(uri)
	if err != nil {
		log.Printf("[ERROR] Spark.%s url(%s) err(%v) \n", data.getName(), uri, err)
		return nil
	}

	defer func() {
		if response != nil {
			_ = response.RawResponse.Body.Close()
		}
	}()

	if response.StatusCode() != http.StatusOK {
		log.Printf("[ERROR] Spark.%s status code: %d, message: %s \n", data.getName(), response.StatusCode(), response.String())
		return nil
	}

	if data.getData() == nil {
		log.Printf("[ERROR] Spark.%s status code: %d, message: %s \n", data.getName(), response.StatusCode(), response.String())
	}

	return data.getData()
}

func get[T RespType](c *SparkClient, req interface{}) interface{} {
	var data T
	uri := data.getUri()
	request := c.client.NewRequest().SetContext(context.WithValue(context.Background(), "Method-Name", uri))
	parameters, err := utils.StructToMap(req)
	if err != nil {
		log.Printf("[ERROR] Rcm.%s get parameters err(%v) \n", data.getName(), err)
		return nil
	}
	request = request.SetPathParams(parameters)
	response, err := request.SetResult(&data).Get(uri)
	if err != nil {
		log.Printf("[ERROR] Rcm.%s url(%s) err(%v) \n", data.getName(), uri, err)
		return nil
	}
	defer func() {
		if response != nil {
			_ = response.RawResponse.Body.Close()
		}
	}()
	if response.StatusCode() != http.StatusOK {
		log.Printf("[ERROR] Rcm,.%s status code: %d, message: %s \n", data.getName(), response.StatusCode(), response.String())
		return nil
	}

	if data.getData() == nil {
		log.Printf("[ERROR] Rcm.%s status code: %d, message: %s \n", data.getName(), response.StatusCode(), response.String())
	}
	return data.getData()
}
