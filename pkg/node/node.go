// Copyright (C) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.
package node

import (
	"github.com/paulgnz/metal-cli/pkg/ansible"
	"github.com/paulgnz/metal-cli/pkg/application"
	"github.com/paulgnz/metal-cli/pkg/models"
	"github.com/paulgnz/metal-cli/pkg/utils"
)

func GetHostWithCloudID(app *application.Avalanche, clusterName string, cloudID string) (*models.Host, error) {
	hosts, err := ansible.GetInventoryFromAnsibleInventoryFile(app.GetAnsibleInventoryDirPath(clusterName))
	if err != nil {
		return nil, err
	}
	monitoringInventoryFile := app.GetMonitoringInventoryDir(clusterName)
	if utils.FileExists(monitoringInventoryFile) {
		monitoringHosts, err := ansible.GetInventoryFromAnsibleInventoryFile(monitoringInventoryFile)
		if err != nil {
			return nil, err
		}
		hosts = append(hosts, monitoringHosts...)
	}
	for _, host := range hosts {
		if host.GetCloudID() == cloudID {
			return host, nil
		}
	}
	return nil, nil
}

func GetICMRelayerHost(app *application.Avalanche, clusterName string) (*models.Host, error) {
	clusterConfig, err := app.GetClusterConfig(clusterName)
	if err != nil {
		return nil, err
	}
	relayerCloudID := ""
	for _, cloudID := range clusterConfig.GetCloudIDs() {
		if nodeConfig, err := app.LoadClusterNodeConfig(cloudID); err != nil {
			return nil, err
		} else if nodeConfig.IsICMRelayer {
			relayerCloudID = nodeConfig.NodeID
		}
	}
	return GetHostWithCloudID(app, clusterName, relayerCloudID)
}
