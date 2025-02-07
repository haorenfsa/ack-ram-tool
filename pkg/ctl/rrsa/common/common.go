package common

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
)

func AllowRRSAFeatureOrDie(ctx context.Context, clusterId string, client *openapi.Client) *types.Cluster {
	c, err := GetRRSAStatus(ctx, clusterId, client)
	if err != nil {
		common.ExitByError(fmt.Sprintf("get status failed: %+v", err))
	}
	if c.State != types.ClusterStateRunning {
		common.ExitByError(fmt.Sprintf("cluster state is not running: %s", c.State))
	}
	if c.ClusterType != types.ClusterTypeManagedKubernetes {
		common.ExitByError("only support managed cluster")
	}
	return c
}

func GetRRSAStatus(ctx context.Context, clusterId string, client openapi.CSClientInterface) (*types.Cluster, error) {
	c, err := client.GetCluster(ctx, clusterId)
	return c, err
}

func getRRSAFailMessage(ctx context.Context, clusterId string, client openapi.CSClientInterface) string {
	logs, err := client.GetRecentClusterLogs(ctx, clusterId)
	if err != nil {
		// TODO: xxx
		return ""
	}
	max := 20
	n := 0
	for _, item := range logs {
		n++
		if n >= max {
			break
		}
		if !strings.Contains(item.Log, "Failed") {
			continue
		}
		if !strings.Contains(item.Log, "RRSA") {
			continue
		}
		return item.Log
	}
	return ""
}

func WaitAddonActionFinished(ctx context.Context, clusterId string, addon types.ClusterAddon, client openapi.CSClientInterface) error {
	n := int64(1)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		ret, err := client.GetAddonStatus(ctx, clusterId, addon.Name)
		if err == nil && ret != nil {
			if addon.NextVersion == "" {
				if ret.Installed() {
					return nil
				}
			} else {
				if ret.Version == addon.NextVersion {
					return nil
				}
			}
		}

		jitter := time.Duration(rand.Int63n(int64(time.Second) * n)) // #nosec G404
		if jitter > time.Second*15 {
			jitter = time.Second*15 + time.Duration(rand.Int63n(int64(time.Second)*10)) // #nosec G404
		}
		time.Sleep(time.Minute + jitter)
		n++
	}
}

func WaitClusterUpdateFinished(ctx context.Context, clusterId, taskId string, client openapi.CSClientInterface) error {
	n := int64(1)
	var taskSuccess bool
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		// check task status first
		if !taskSuccess {
			c, err := client.GetTask(ctx, taskId)
			if err != nil {
				return err
			}
			if c.State.IsNotSuccess() {
				return fmt.Errorf("task not success: %s, message: %s",
					c.State, getRRSAFailMessage(ctx, clusterId, client),
				)
			}
			if c.State == types.ClusterTaskStateSuccess {
				taskSuccess = true
				n = 1
			}
		}
		// if task successes, then check cluster status
		if taskSuccess {
			c, err := client.GetCluster(ctx, clusterId)
			if err != nil {
				return err
			}
			if c.State.IsRunning() {
				return nil
			}
		}
		jitter := time.Duration(rand.Int63n(int64(time.Second) * n)) // #nosec G404
		if jitter > time.Second*15 {
			jitter = time.Second*15 + time.Duration(rand.Int63n(int64(time.Second)*10)) // #nosec G404
		}
		time.Sleep(time.Second*20 + jitter)
		n++
	}
}
