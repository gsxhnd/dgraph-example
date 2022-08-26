package example

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type dc struct {
	client *dgo.Dgraph
}

func NewClient() *dc {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial("127.0.0.1:9080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	c := dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)

	return &dc{
		client: c,
	}
}

func (d *dc) Setup() {
	if err := d.client.Alter(context.Background(), &api.Operation{
		Schema: `
		type Person {
			name
			age
			friend
			owns_pet
		}

		type Animal {
			name
			age
		}

		name: string @index(term) @lang .
		age: int @index(int) .
		friend: [uid] @count .
		owns_pet: [uid] .
		`,
	}); err != nil {
		log.Fatal(err)
	}
	mu := &api.Mutation{
		CommitNow: true,
	}

	if _, err := d.client.NewTxn().Mutate(context.Background(), mu); err != nil {
		log.Fatal(err)
	}

}

// DropAll 删除所有数据
func (d *dc) DropAll() {
	err := d.client.Alter(context.Background(), &api.Operation{
		DropAll: true,
	})
	log.Fatal(err)
}
