package utils

import (
	"github.com/saichler/l8collector/go/collector/common"
	"github.com/saichler/l8pollaris/go/types"
)

const (
	InvServiceName = "NetworkBox"
	K8sServiceName = "Cluster"
)

func CreateDevice(ip string, serviceArea uint16) *types.Device {
	device := &types.Device{}
	device.DeviceId = ip
	device.InventoryService = &types.DeviceServiceInfo{ServiceName: InvServiceName, ServiceArea: int32(serviceArea)}
	device.ParsingService = &types.DeviceServiceInfo{ServiceName: common.ParserServicePrefix + InvServiceName, ServiceArea: int32(serviceArea)}
	device.Hosts = make(map[string]*types.Host)
	host := &types.Host{}
	host.DeviceId = device.DeviceId

	host.Configs = make(map[int32]*types.Connection)
	device.Hosts[device.DeviceId] = host

	sshConfig := &types.Connection{}
	sshConfig.Protocol = types.Protocol_SSH
	sshConfig.Port = 22
	sshConfig.Addr = ip
	sshConfig.Username = "admin"
	sshConfig.Password = "admin"
	sshConfig.Terminal = "vt100"
	sshConfig.Timeout = 15

	host.Configs[int32(sshConfig.Protocol)] = sshConfig

	snmpConfig := &types.Connection{}
	snmpConfig.Protocol = types.Protocol_SNMPV2
	snmpConfig.Addr = ip
	snmpConfig.Port = 161
	snmpConfig.Timeout = 15
	snmpConfig.ReadCommunity = "public"

	host.Configs[int32(snmpConfig.Protocol)] = snmpConfig

	return device
}
