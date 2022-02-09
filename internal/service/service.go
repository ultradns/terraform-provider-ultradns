package service

import (
	"github.com/ultradns/ultradns-go-sdk/pkg/client"
	"github.com/ultradns/ultradns-go-sdk/pkg/rdpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/record"
	"github.com/ultradns/ultradns-go-sdk/pkg/sfpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/slbpool"
	"github.com/ultradns/ultradns-go-sdk/pkg/zone"
)

type Service struct {
	ZoneService    *zone.Service
	RecordService  *record.Service
	RDPoolService  *rdpool.Service
	SFPoolService  *sfpool.Service
	SLBPoolService *slbpool.Service
}

func NewService(c *client.Client) (*Service, error) {
	service := &Service{}

	var err error

	if service.ZoneService, err = zone.Get(c); err != nil {
		return nil, err
	}

	if service.RecordService, err = record.Get(c); err != nil {
		return nil, err
	}

	if service.RDPoolService, err = rdpool.Get(c); err != nil {
		return nil, err
	}

	if service.SFPoolService, err = sfpool.Get(c); err != nil {
		return nil, err
	}

	if service.SLBPoolService, err = slbpool.Get(c); err != nil {
		return nil, err
	}

	return service, nil
}
