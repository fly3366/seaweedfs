package operation

import (
	"context"
	"github.com/chrislusf/seaweedfs/weed/pb/volume_server_pb"
	"google.golang.org/grpc"
)

func GetVolumeSyncStatus(server string, grpcDialOption grpc.DialOption, vid uint32) (resp *volume_server_pb.VolumeSyncStatusResponse, err error) {

	WithVolumeServerClient(server, grpcDialOption, func(ctx context.Context, client volume_server_pb.VolumeServerClient) error {

		resp, err = client.VolumeSyncStatus(ctx, &volume_server_pb.VolumeSyncStatusRequest{
			VolumeId: vid,
		})
		return nil
	})

	return
}
