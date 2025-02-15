package enable

import (
	"context"
	"fmt"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl"
	"log"
	"time"

	ctlcommon "github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/ctl/rrsa/common"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/openapi"
	"github.com/AliyunContainerService/ack-ram-tool/pkg/types"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable RRSA feature",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		client := ctlcommon.GetClientOrDie()
		ctlcommon.YesOrExit("Are you sure you want to enable RRSA feature?")
		ctx := context.Background()

		clusterId := ctl.GlobalOption.ClusterId
		c := common.AllowRRSAFeatureOrDie(ctx, clusterId, client)
		if c.MetaData.RRSAConfig.Enabled {
			log.Println("RRSA feature is already enabled.")
			return
		}

		var task *types.ClusterTask
		var err error
		log.Println("Start to enable RRSA feature")
		spin := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		spin.Start()

		if task, err = enableRRSA(ctx, clusterId, client); err != nil {
			spin.Stop()
			ctlcommon.ExitByError(fmt.Sprintf("Failed to enable RRSA feature for cluster %s: %+v", clusterId, err))
		}
		ctx, cancel := context.WithTimeout(ctx, time.Minute*15)
		defer cancel()
		if err := common.WaitClusterUpdateFinished(ctx, clusterId, task.TaskId, client); err != nil {
			spin.Stop()
			ctlcommon.ExitByError(fmt.Sprintf("Failed to enable RRSA feature for cluster %s: %+v", clusterId, err))
		}

		spin.Stop()
		log.Printf("Enable RRSA feature for cluster %s successfully", clusterId)
	},
}

func enableRRSA(ctx context.Context, clusterId string, client openapi.CSClientInterface) (*types.ClusterTask, error) {
	boolValue := true
	return client.UpdateCluster(ctx, clusterId, openapi.UpdateClusterOption{EnableRRSA: &boolValue})
}

func SetupCmd(rootCmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
	ctlcommon.SetupClusterIdFlag(cmd)
}
