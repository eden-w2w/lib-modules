package gaode

import (
	"github.com/eden-framework/courier"
	"github.com/eden-framework/courier/client"
	"github.com/eden-w2w/lib-modules/constants/general_errors"
	"github.com/sirupsen/logrus"
)

type GaodeClient struct {
	Key    string
	Client *client.Client
}

func (c *GaodeClient) District(req DistrictRequest, metas ...courier.Metadata) (resp *DistrictResponse, err error) {
	req.Key = c.Key
	request := c.Client.Request("", "GET", "/v3/config/district", req, metas...)
	resp = &DistrictResponse{}
	err = request.Do().Into(resp)
	if err != nil {
		logrus.Errorf("[GaodeClient] District err: %v, request: %+v", err, req)
		return nil, general_errors.BadGateway
	}
	return
}
