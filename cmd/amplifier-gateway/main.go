package main
import (
 "flag"
 "net/http"

 "github.com/golang/glog"
 "golang.org/x/net/context"
 "github.com/grpc-ecosystem/grpc-gateway/runtime"
 "google.golang.org/grpc"

 "github.com/appcelerator/amp/api/rpc/logs"
)

var (
 amplifierEndpoint = flag.String("amplifier_endpoint", "localhost:8080", "endpoint of amplifier")
)

func run() error {
 ctx := context.Background()
 ctx, cancel := context.WithCancel(ctx)
 defer cancel()

 mux := runtime.NewServeMux()
 opts := []grpc.DialOption{grpc.WithInsecure()}
 err := logs.RegisterLogsHandlerFromEndpoint(ctx, mux, *amplifierEndpoint, opts)
 if err != nil {
   return err
 }

 http.ListenAndServe(":3000", mux)
 return nil
}

func main() {
 flag.Parse()
 defer glog.Flush()

 if err := run(); err != nil {
   glog.Fatal(err)
 }
}