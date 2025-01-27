package send

import (
	"errors"
	"fmt"

	"harmonycloud.cn/stellaris/config"
	proxy_stream "harmonycloud.cn/stellaris/pkg/proxy/stream"
	"harmonycloud.cn/stellaris/pkg/model"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var resourceLog = logf.Log.WithName("proxy_send_resource")

func SendSyncResourceRequest(request *config.Request) error {
	resourceLog.Info(fmt.Sprintf("start send resource request to core, request Data:%s", request.Body))
	stream := proxy_stream.GetConnection()
	var err error
	if stream == nil {
		err = errors.New("new stream failed")
		resourceLog.Error(err, "send resource request")
		return err
	}
	err = stream.Send(request)
	if err != nil {
		resourceLog.Error(err, "send resource request failed")
		return err
	}
	return nil
}

func NewResourceRequest(resType model.ServiceRequestType, clusterName string, body string) (*config.Request, error) {
	if len(clusterName) == 0 || len(body) == 0 {
		return nil, errors.New("clusterName or body is empty")
	}
	return &config.Request{
		Type:        resType.String(),
		ClusterName: clusterName,
		Body:        body,
	}, nil
}
