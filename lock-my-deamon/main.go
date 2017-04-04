package main

import (
	"log"
	"os"
	"strconv"

	"time"

	"github.com/Sirupsen/logrus"
	"github.com/docker/go-plugins-helpers/volume"
)

const socketAddress = "/run/docker/plugins/lock-my-daemon.sock"

type lockMyDaemonDriver struct {
}

func newSshfsDriver(root string) (*lockMyDaemonDriver, error) {
	logrus.WithField("method", "new driver").Debug(root)

	d := &lockMyDaemonDriver{}

	return d, nil
}

func (d *lockMyDaemonDriver) Create(r volume.Request) volume.Response {
	logrus.WithField("method", "create").Debugf("%#v", r)

	return volume.Response{}
}

func (d *lockMyDaemonDriver) Remove(r volume.Request) volume.Response {
	logrus.WithField("method", "remove").Debugf("%#v", r)

	return volume.Response{}
}

func (d *lockMyDaemonDriver) Path(r volume.Request) volume.Response {
	logrus.WithField("method", "path").Debugf("%#v", r)

	return volume.Response{Mountpoint: "/tmp"}
}

func (d *lockMyDaemonDriver) Mount(r volume.MountRequest) volume.Response {
	logrus.WithField("method", "mount").Debugf("%#v", r)

	return volume.Response{Mountpoint: "/tmp"}
}

func (d *lockMyDaemonDriver) Unmount(r volume.UnmountRequest) volume.Response {
	logrus.WithField("method", "unmount").Debugf("%#v", r)

	return volume.Response{}
}

func (d *lockMyDaemonDriver) Get(r volume.Request) volume.Response {
	logrus.WithField("method", "get").Debugf("%#v", r)

	if r.Name == "lock-me" {
		time.Sleep(60 * 60 * time.Second)
	}

	return volume.Response{Volume: &volume.Volume{Name: r.Name, Mountpoint: "/tmp"}}
}

func (d *lockMyDaemonDriver) List(r volume.Request) volume.Response {
	logrus.WithField("method", "list").Debugf("%#v", r)

	return volume.Response{
		Volumes: []*volume.Volume{
			{
				Name:       "foo",
				Mountpoint: "/tmp",
			},
			{
				Name:       "lock-me",
				Mountpoint: "/tmp",
			},
		},
	}
}

func (d *lockMyDaemonDriver) Capabilities(r volume.Request) volume.Response {
	logrus.WithField("method", "capabilities").Debugf("%#v", r)

	return volume.Response{Capabilities: volume.Capability{Scope: "local"}}
}

func main() {
	debug := os.Getenv("DEBUG")
	if ok, _ := strconv.ParseBool(debug); ok {
		logrus.SetLevel(logrus.DebugLevel)
	}

	d, err := newSshfsDriver("/mnt")
	if err != nil {
		log.Fatal(err)
	}
	h := volume.NewHandler(d)
	logrus.Infof("listening on %s", socketAddress)
	logrus.Error(h.ServeUnix(socketAddress, 0))
}
