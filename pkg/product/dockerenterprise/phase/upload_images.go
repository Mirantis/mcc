package phase

import (
	"io/ioutil"
	"path"
	"path/filepath"

	"github.com/Mirantis/mcc/pkg/api"
	"github.com/Mirantis/mcc/pkg/phase"
	"github.com/Mirantis/mcc/pkg/util"

	"github.com/alessio/shellescape"
	log "github.com/sirupsen/logrus"
)

// LoadImages phase uploads + docker loads images from host's imageDir to hosts
type LoadImages struct {
	phase.Analytics
	phase.HostSelectPhase
	totalBytes uint64
}

// Title is the title for the phase
func (p *LoadImages) Title() string {
	return "Upload images"
}

// HostFilterFunc returns true for hosts that have non-empty list of hooks returned by the StepListFunc
func (p *LoadImages) HostFilterFunc(h *api.Host) bool {
	if h.ImageDir == "" {
		return false
	}
	log.Debugf("%s: listing images in imageDir '%s'", h.Address, h.ImageDir)

	files, err := ioutil.ReadDir(h.ImageDir)
	if err != nil {
		log.Errorf("%s: failed to list images in imageDir '%s': %s", h.Address, h.ImageDir, err.Error())
		return false
	}

	for _, info := range files {
		if info.IsDir() {
			continue
		}

		ext := filepath.Ext(info.Name())
		if ext != ".tar" && ext != ".gz" {
			continue
		}

		imagePath := filepath.Join(h.ImageDir, info.Name())
		h.Metadata.ImagesToUpload = append(h.Metadata.ImagesToUpload, imagePath)
		h.Metadata.TotalImageBytes += uint64(info.Size())
	}

	return h.Metadata.TotalImageBytes > 0
}

// Prepare collects the hosts
func (p *LoadImages) Prepare(config *api.ClusterConfig) error {
	p.Config = config
	log.Debugf("collecting hosts for phase %s", p.Title())
	hosts := config.Spec.Hosts.Filter(p.HostFilterFunc)
	log.Debugf("found %d hosts for phase %s", len(hosts), p.Title())
	p.Hosts = hosts
	return nil
}

// Run does all the work
func (p *LoadImages) Run() error {
	var totalBytes uint64
	p.Hosts.Each(func(h *api.Host) error {
		totalBytes += h.Metadata.TotalImageBytes
		return nil
	})

	log.Infof("total %s of images to upload", util.FormatBytes(totalBytes))

	return p.Hosts.Each(func(h *api.Host) error {
		for idx, f := range h.Metadata.ImagesToUpload {
			log.Debugf("%s: uploading image %d/%d", h.Address, idx+1, len(h.Metadata.ImagesToUpload))

			base := path.Base(f)
			df := h.Configurer.JoinPath(h.Configurer.Pwd(), base)
			err := h.WriteFileLarge(f, df)
			if err != nil {
				return err
			}

			log.Infof("%s: loading image %d/%d : %s", h.Address, idx+1, len(h.Metadata.ImagesToUpload), base)
			err = h.Exec(h.Configurer.DockerCommandf("load -i %s", shellescape.Quote(base)))
			if err != nil {
				return err
			}
		}
		return nil
	})
}