package openstack

import (
	"github.com/emc-advanced-dev/pkg/errors"
	"github.com/cf-unik/unik/pkg/types"
)

func (p *OpenstackProvider) GetVolume(nameOrIdPrefix string) (*types.Volume, error) {
	return nil, errors.New("not yet supportded for openstack", nil)
}
