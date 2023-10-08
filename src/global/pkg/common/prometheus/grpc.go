package prometheus

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/cwloo/gonet/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

func UnaryServerInterceptorProme(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	remote, _ := peer.FromContext(ctx)
	remoteAddr := remote.Addr.String()

	in, _ := json.Marshal(req)
	inStr := string(in)
	logs.Infof(strings.Join([]string{remoteAddr, "access_start", info.FullMethod, inStr}, " "))

	start := time.Now()
	defer func() {
		j, _ := json.Marshal(resp)
		duration := int64(time.Since(start) / time.Millisecond)
		if duration >= 500 {
			msg := strings.Join([]string{remoteAddr, "access_end", info.FullMethod, inStr, string(j), err.Error(), "elapsed:", fmt.Sprintf("%v", time.Since(start))}, " ")
			logs.Infof(msg)
		} else {
			msg := strings.Join([]string{remoteAddr, "access_end", info.FullMethod, inStr, string(j), err.Error(), "elapsed:", fmt.Sprintf("%v", time.Since(start))}, " ")
			logs.Infof(msg)
		}
	}()
	resp, err = handler(ctx, req)
	return
}
