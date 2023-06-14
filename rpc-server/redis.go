import (
    "context"
    "encoding/json"
    "github.com/redis/go-redis/v9"
    "time"
)

type RedisClient struct {
    cli *redis.Client
}

func (c *RedisClient) InitClient(ctx context.Context, address, password string) error {
    r := redis.NewClient(&redis.Options{
       Addr:     address,
       Password: password, // no password set
       DB:       0,        // use default DB
    })

    // test connection
    if err := r.Ping(ctx).Err(); err != nil {
       return err
    }

    c.cli = r
    return nil
}
var (
    rdb = &RedisClient{} // make the RedisClient with global visibility in the 'main' scope
)

func main() {
    ctx := context.Background() // https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go

    err := rdb.InitClient(ctx, "redis:6379", "")
    if err != nil {
       errMsg := fmt.Sprintf("failed to init Redis client, err: %v", err)
       log.Fatal(errMsg)
    }

    r, err := etcd.NewEtcdRegistry([]string{"etcd:2379"}) // r should not be reused.
    if err != nil {
       log.Fatal(err)
    }

    svr := rpc.NewServer(new(IMServiceImpl), server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
       ServiceName: "demo.rpc.server",
    }))

    err = svr.Run()
    if err != nil {
       log.Println(err.Error())
    }
}
