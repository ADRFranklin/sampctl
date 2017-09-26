package main

import (
	"bufio"
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
	"github.com/google/uuid"
)

// RunContainer does what Run does but inside a Linux container
func RunContainer(endpoint, version, dir, appVersion string) (err error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	cnt, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image:        "southclaws/sampctl:" + appVersion,
			Entrypoint:   strslice.StrSlice{"sampctl", "run"},
			Tty:          true,
			AttachStdout: true,
			AttachStderr: true,
		},
		&container.HostConfig{
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: dir,
					Target: "/samp",
				},
			},
		},
		&network.NetworkingConfig{},
		"sampctl-"+uuid.New().String())
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting container...")
	err = cli.ContainerStart(context.Background(), cnt.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Warnings:", cnt.Warnings)
	go func() {
		reader, err := cli.ContainerLogs(context.Background(), cnt.ID, types.ContainerLogsOptions{
			ShowStdout: true,
			ShowStderr: true,
			Follow:     true,
			Timestamps: false,
		})
		if err != nil {
			panic(err)
		}
		defer reader.Close()

		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	// for now, we just wait for the app to exit and prints logs afterwards
	n, err := cli.ContainerWait(context.Background(), cnt.ID)
	fmt.Println("container exited:", n, err)

	return
}