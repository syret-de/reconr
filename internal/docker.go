package internal

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/pkg/archive"
	"github.com/mitchellh/go-homedir"
	"log"
	"path/filepath"
	"time"

	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type Docker struct {
	ctx     context.Context
	cli     *client.Client
	config  Config
	logfile *os.File
}

func NewDocker(config Config) (Docker, error) {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return Docker{}, err
	}
	defer cli.Close()

	return Docker{ctx: ctx, cli: cli, config: config}, nil
}

func (d *Docker) setLogFile(logfile *os.File) {
	d.logfile = logfile
}

func (d *Docker) getLogFile() *os.File {
	return d.logfile
}

func (d *Docker) Build(path string, logger Logger) error {
	log.Printf("[*] Docker build started\n")
	d.setLogFile(logger.GetLogfile())

	filePath, _ := homedir.Expand(path)
	buildCtx, _ := archive.TarWithOptions(filePath, &archive.TarOptions{})

	opt := types.ImageBuildOptions{
		Context:    buildCtx,
		Dockerfile: "Dockerfile",
		Tags:       []string{"reconr"},
	}

	logs, err := d.cli.ImageBuild(context.Background(), buildCtx, opt)
	if err != nil {
		return err
	}

	_, err = io.Copy(d.getLogFile(), logs.Body)
	if err != nil {
		return err
	}

	defer logs.Body.Close()
	return nil
}

func (d *Docker) Run(commands string, id int) error {
	log.Printf("[*] Task %d Started\n", id)
	log.Printf("[*] Docker run %s\n", commands)

	fullPath := fmt.Sprintf("%s/%s", d.config.GetWorkPath(), d.config.GetTarget())
	outPath, err := filepath.Abs(fullPath)
	if err != nil {
		return err
	}

	configPath, err := filepath.Abs(d.config.GetConfigPath())
	if err != nil {
		return err
	}

	config := container.Config{
		Image: "reconr",
		Cmd:   []string{commands},
		Tty:   true,
	}

	hostConfig := container.HostConfig{
		AutoRemove: true,
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: outPath,
				Target: d.config.GetMountWork(),
			},
			{
				Type:   mount.TypeBind,
				Source: configPath,
				Target: d.config.GetMountConfig(),
			},
		},
		NetworkMode: "host",
	}

	resp, err := d.cli.ContainerCreate(d.ctx, &config, &hostConfig, nil, nil, fmt.Sprintf("reconr%d", id))
	log.Printf("[*] Created container %s\n", resp.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = d.cli.ContainerStart(d.ctx, resp.ID, container.StartOptions{})
	log.Printf("[*] Started container %s\n", resp.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	statusCh, errCh := d.cli.ContainerWait(d.ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return err
		}
	case <-statusCh:
	}

	out, err := d.cli.ContainerLogs(d.ctx, resp.ID, container.LogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		// This happens when the task is too short
		if err.Error() != "Error response from daemon: can not get logs from container which is dead or marked for removal" {
			return err
		}
		log.Printf("[*] Task to short, no stdout\n")
	} else {
		log.Printf("[*] Task stdout: \n")
		_, err = io.Copy(d.getLogFile(), out)
		if err != nil {
			return err
		}
	}

	log.Printf("[*] Stopped container %s\n", resp.ID)
	time.Sleep(100 * time.Millisecond)
	return nil
}
