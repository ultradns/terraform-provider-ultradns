package probeping

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/probe"
	sdkprobe "github.com/ultradns/ultradns-go-sdk/pkg/probe"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/ping"
)

func flattenProbePING(probeData *sdkprobe.Probe, rd *schema.ResourceData) error {
	if err := probe.FlattenProbe(probeData, rd); err != nil {
		return err
	}

	details := probeData.Details.(*ping.Details)

	if err := rd.Set("packets", details.Packets); err != nil {
		return err
	}

	if err := rd.Set("packet_size", details.PacketSize); err != nil {
		return err
	}

	if details.Limits != nil {
		limits := details.Limits

		if err := rd.Set("loss_percent_limit", probe.GetLimitList(limits.LossPercent)); err != nil {
			return err
		}

		if err := rd.Set("total_limit", probe.GetLimitList(limits.Total)); err != nil {
			return err
		}

		if err := rd.Set("average_limit", probe.GetLimitList(limits.Average)); err != nil {
			return err
		}

		if err := rd.Set("run_limit", probe.GetLimitList(limits.Run)); err != nil {
			return err
		}

		if err := rd.Set("avg_run_limit", probe.GetLimitList(limits.AvgRun)); err != nil {
			return err
		}
	}

	return nil
}
