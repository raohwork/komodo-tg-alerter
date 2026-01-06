/*
Copyright © 2026 Ronmi Ren

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package tmpl

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"time"

	"github.com/raohwork/komodo-tg-alerter/komodo"
)

// helper function to create PayloadItem from any value
func payloadItem(v any) komodo.PayloadItem {
	b, _ := json.Marshal(v)
	return komodo.PayloadItem(b)
}

// helper function to create Map
func payloadMap(data map[string]any) komodo.Map {
	m := make(komodo.Map)
	for k, v := range data {
		m[k] = payloadItem(v)
	}
	return m
}

// sampleAlerts provides example AlertInfo for each alert type
var sampleAlerts = map[string]*komodo.AlertInfo{
	"Test": {
		Timestamp: time.Now().UnixMilli(),
		Level:     "info",
		Resolved:  false,
		Target: komodo.AlertTarget{
			ID:   "test-target-1",
			Type: "test",
		},
		Data: komodo.AlertData{
			Type: "Test",
			Payload: payloadMap(map[string]any{
				"id":   "test-123",
				"name": "Test Alert Example",
			}),
		},
	},
	"ServerVersionMismatch": {
		Timestamp: time.Now().UnixMilli(),
		Level:     "warning",
		Resolved:  false,
		Target: komodo.AlertTarget{
			ID:   "server-1",
			Type: "server",
		},
		Data: komodo.AlertData{
			Type: "ServerVersionMismatch",
			Payload: payloadMap(map[string]any{
				"id":             "server-1",
				"name":           "production-server",
				"region":         "us-west-2",
				"server_version": "v1.2.3",
				"core_version":   "v1.2.5",
			}),
		},
	},
	"ServerUnreachable": {
		Timestamp: time.Now().UnixMilli(),
		Level:     "critical",
		Resolved:  false,
		Target: komodo.AlertTarget{
			ID:   "server-2",
			Type: "server",
		},
		Data: komodo.AlertData{
			Type: "ServerUnreachable",
			Payload: payloadMap(map[string]any{
				"id":     "server-2",
				"name":   "backup-server",
				"region": "eu-central-1",
				"err":    "connection timeout after 30s",
			}),
		},
	},
	"ServerCpu": {
		Timestamp: time.Now().UnixMilli(),
		Level:     "warning",
		Resolved:  false,
		Target: komodo.AlertTarget{
			ID:   "server-3",
			Type: "server",
		},
		Data: komodo.AlertData{
			Type: "ServerCpu",
			Payload: payloadMap(map[string]any{
				"id":         "server-3",
				"name":       "api-server-1",
				"region":     "us-east-1",
				"percentage": 85.5,
			}),
		},
	},
	"ServerMem": {
		Timestamp: time.Now().UnixMilli(),
		Level:     "warning",
		Resolved:  false,
		Target: komodo.AlertTarget{
			ID:   "server-4",
			Type: "server",
		},
		Data: komodo.AlertData{
			Type: "ServerMem",
			Payload: payloadMap(map[string]any{
				"id":       "server-4",
				"name":     "db-server-1",
				"region":   "ap-southeast-1",
				"used_gb":  14.5,
				"total_gb": 16.0,
			}),
		},
	},
	"ServerDisk": {
		Timestamp: time.Now().UnixMilli(),
		Level:     "critical",
		Resolved:  false,
		Target: komodo.AlertTarget{
			ID:   "server-5",
			Type: "server",
		},
		Data: komodo.AlertData{
			Type: "ServerDisk",
			Payload: payloadMap(map[string]any{
				"id":       "server-5",
				"name":     "storage-server",
				"region":   "us-west-1",
				"path":     "/var/lib/docker",
				"used_gb":  95.2,
				"total_gb": 100.0,
			}),
		},
	},
	"ContainerStateChange": {
		Timestamp: time.Now().UnixMilli(),
		Level:     "info",
		Resolved:  false,
		Target: komodo.AlertTarget{
			ID:   "container-1",
			Type: "container",
		},
		Data: komodo.AlertData{
			Type: "ContainerStateChange",
			Payload: payloadMap(map[string]any{
				"id":          "container-1",
				"name":        "web-app",
				"server_id":   "server-1",
				"server_name": "production-server",
				"from":        "running",
				"to":          "stopped",
			}),
		},
	},
	"DeploymentImageUpdateAvailable": {
		Timestamp: time.Now().UnixMilli(),
		Level:     "info",
		Resolved:  false,
		Target: komodo.AlertTarget{
			ID:   "deployment-1",
			Type: "deployment",
		},
		Data: komodo.AlertData{
			Type: "DeploymentImageUpdateAvailable",
			Payload: payloadMap(map[string]any{
				"id":          "deployment-1",
				"name":        "api-deployment",
				"server_id":   "server-2",
				"server_name": "api-server",
				"image":       "myapp:v2.0.0",
			}),
		},
	},
	"DeploymentAutoUpdated": {
		Timestamp: time.Now().UnixMilli(),
		Level:     "info",
		Resolved:  false,
		Target: komodo.AlertTarget{
			ID:   "deployment-2",
			Type: "deployment",
		},
		Data: komodo.AlertData{
			Type: "DeploymentAutoUpdated",
			Payload: payloadMap(map[string]any{
				"id":          "deployment-2",
				"name":        "worker-deployment",
				"server_id":   "server-3",
				"server_name": "worker-server",
				"image":       "worker:v1.5.0",
			}),
		},
	},
	"StackStateChange": {
		Timestamp: time.Now().UnixMilli(),
		Level:     "warning",
		Resolved:  false,
		Target: komodo.AlertTarget{
			ID:   "stack-1",
			Type: "stack",
		},
		Data: komodo.AlertData{
			Type: "StackStateChange",
			Payload: payloadMap(map[string]any{
				"id":          "stack-1",
				"name":        "monitoring-stack",
				"server_id":   "server-4",
				"server_name": "monitoring-server",
				"from":        "running",
				"to":          "degraded",
			}),
		},
	},
	"StackImageUpdateAvailable": {
		Timestamp: time.Now().UnixMilli(),
		Level:     "info",
		Resolved:  false,
		Target: komodo.AlertTarget{
			ID:   "stack-2",
			Type: "stack",
		},
		Data: komodo.AlertData{
			Type: "StackImageUpdateAvailable",
			Payload: payloadMap(map[string]any{
				"id":          "stack-2",
				"name":        "web-stack",
				"server_id":   "server-5",
				"server_name": "web-server",
				"service":     "nginx",
				"image":       "nginx:1.25.0",
			}),
		},
	},
	"StackAutoUpdated": {
		Timestamp: time.Now().UnixMilli(),
		Level:     "info",
		Resolved:  false,
		Target: komodo.AlertTarget{
			ID:   "stack-3",
			Type: "stack",
		},
		Data: komodo.AlertData{
			Type: "StackAutoUpdated",
			Payload: payloadMap(map[string]any{
				"id":          "stack-3",
				"name":        "app-stack",
				"server_id":   "server-6",
				"server_name": "app-server",
				"images":      "frontend:v2.1.0, backend:v3.0.0, redis:7.0",
			}),
		},
	},
	"AwsBuilderTerminationFailed": {
		Timestamp: time.Now().UnixMilli(),
		Level:     "critical",
		Resolved:  false,
		Target: komodo.AlertTarget{
			ID:   "builder-1",
			Type: "builder",
		},
		Data: komodo.AlertData{
			Type: "AwsBuilderTerminationFailed",
			Payload: payloadMap(map[string]any{
				"instance_id": "i-1234567890abcdef0",
				"message":     "Unable to terminate instance: InvalidInstanceID.NotFound",
			}),
		},
	},
	"ResourceSyncPendingUpdates": {
		Timestamp: time.Now().UnixMilli(),
		Level:     "info",
		Resolved:  false,
		Target: komodo.AlertTarget{
			ID:   "sync-1",
			Type: "sync",
		},
		Data: komodo.AlertData{
			Type: "ResourceSyncPendingUpdates",
			Payload: payloadMap(map[string]any{
				"id":   "sync-1",
				"name": "config-sync",
			}),
		},
	},
	"BuildFailed": {
		Timestamp: time.Now().UnixMilli(),
		Level:     "critical",
		Resolved:  false,
		Target: komodo.AlertTarget{
			ID:   "build-1",
			Type: "build",
		},
		Data: komodo.AlertData{
			Type: "BuildFailed",
			Payload: payloadMap(map[string]any{
				"id":      "build-1",
				"name":    "frontend-build",
				"version": "v2.5.0",
			}),
		},
	},
	"RepoBuildFailed": {
		Timestamp: time.Now().UnixMilli(),
		Level:     "critical",
		Resolved:  false,
		Target: komodo.AlertTarget{
			ID:   "repo-1",
			Type: "repo",
		},
		Data: komodo.AlertData{
			Type: "RepoBuildFailed",
			Payload: payloadMap(map[string]any{
				"id":   "repo-1",
				"name": "backend-repo",
			}),
		},
	},
	"ProcedureFailed": {
		Timestamp: time.Now().UnixMilli(),
		Level:     "critical",
		Resolved:  false,
		Target: komodo.AlertTarget{
			ID:   "procedure-1",
			Type: "procedure",
		},
		Data: komodo.AlertData{
			Type: "ProcedureFailed",
			Payload: payloadMap(map[string]any{
				"id":   "procedure-1",
				"name": "database-backup",
			}),
		},
	},
	"ActionFailed": {
		Timestamp: time.Now().UnixMilli(),
		Level:     "critical",
		Resolved:  false,
		Target: komodo.AlertTarget{
			ID:   "action-1",
			Type: "action",
		},
		Data: komodo.AlertData{
			Type: "ActionFailed",
			Payload: payloadMap(map[string]any{
				"id":   "action-1",
				"name": "deploy-to-production",
			}),
		},
	},
	"ScheduleRun": {
		Timestamp: time.Now().UnixMilli(),
		Level:     "info",
		Resolved:  false,
		Target: komodo.AlertTarget{
			ID:   "schedule-1",
			Type: "schedule",
		},
		Data: komodo.AlertData{
			Type: "ScheduleRun",
			Payload: payloadMap(map[string]any{
				"resource_type": "backup",
				"id":            "schedule-1",
				"name":          "nightly-backup",
			}),
		},
	},
	"Custom": {
		Timestamp: time.Now().UnixMilli(),
		Level:     "info",
		Resolved:  false,
		Target: komodo.AlertTarget{
			ID:   "custom-1",
			Type: "custom",
		},
		Data: komodo.AlertData{
			Type: "Custom",
			Payload: payloadMap(map[string]any{
				"message": "Custom alert triggered",
				"details": "This is a custom alert with additional details",
			}),
		},
	},
	"None": {
		Timestamp: time.Now().UnixMilli(),
		Level:     "info",
		Resolved:  false,
		Target: komodo.AlertTarget{
			ID:   "none-1",
			Type: "none",
		},
		Data: komodo.AlertData{
			Type:    "None",
			Payload: komodo.Map{},
		},
	},
}

// Lint checks all templates for syntax errors and renders them with sample data
func Lint(fs fs.FS, tz *time.Location) error {
	if fs == nil {
		fs = Files
	}

	renderer := NewRenderer(fs, tz)
	var hasError bool

	// Try to render each sample alert
	for typeName, sampleData := range sampleAlerts {
		fmt.Printf("📝 Rendering %s...\n", typeName)
		result, err := renderer.Render(sampleData)
		if err != nil {
			fmt.Printf("❌ Error: %v\n\n", err)
			hasError = true
			continue
		}

		fmt.Println("✅ Success:")
		fmt.Println("---")
		fmt.Println(result)
		fmt.Println("---\n")
	}

	if hasError {
		return fmt.Errorf("some templates failed to render")
	}

	fmt.Println("✅ All templates validated successfully!")
	return nil
}
