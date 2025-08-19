package vk

import (
	"context"
	"fmt"
	"os"

	"github.com/valkey-io/valkey-go"
	"github.com/will-x86/r-place/pkg/models"
)

var client valkey.Client

func InitialiseValkey() error {
	var err error
	client, err = valkey.NewClient(valkey.ClientOption{InitAddress: []string{os.Getenv("VALKEY_HOST")}, Password: os.Getenv("VALKEY_PASSWORD")})
	if err != nil {
		return err
	}

	/*
		ctx := context.Background()
		SET key val NX
		err = client.Do(ctx, client.B().Set().Key("key").Value("val").Nx().Build()).Error()
		// HGETALL hm
		hm, err := client.Do(ctx, client.B().Hgetall().Key("hm").Build()).AsStrMap()*/
	return nil
}
func CloseValkey() {
	client.Close()
}

func SetCanvasValue(ctx context.Context, c models.Canvas) error {
	// Format is KEY:x-y to VALUE #XXXXXX
	//cache.SetColor(c.X, c.Y, c.Hex)
	return client.Do(ctx, client.B().Set().Key(fmt.Sprintf("%d-%d", c.X, c.Y)).Value(c.Hex).Build()).Error()
}

// Format is KEY:x-y to VALUE #XXXXXX
func GetSingleCanvasValue(ctx context.Context, c models.Canvas) (string, error) {
	res, err := client.Do(ctx, client.B().Get().Key(fmt.Sprintf("%d-%d", c.X, c.Y)).Build()).ToString()
	if err != nil {
		return "", err
	}
	return res, nil
}
