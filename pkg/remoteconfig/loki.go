// Copyright (C) 2025, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package remoteconfig

import (
	"github.com/paulgnz/metal-cli/pkg/utils"
)

func LokiFoldersToCreate() []string {
	return []string{utils.GetRemoteComposeServicePath("loki", "data")}
}
