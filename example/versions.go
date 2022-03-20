package main

import (
	"github.com/rgooding/goav/avformat"
	"log"

	"github.com/rgooding/goav/avcodec"
	"github.com/rgooding/goav/avdevice"
	"github.com/rgooding/goav/avfilter"
	"github.com/rgooding/goav/avutil"
	"github.com/rgooding/goav/swresample"
	"github.com/rgooding/goav/swscale"
)

func main() {
	log.Printf("AvFilter Version:\t%v", avfilter.AvfilterVersion())
	log.Printf("AvDevice Version:\t%v", avdevice.AvdeviceVersion())
	log.Printf("SWScale Version:\t%v", swscale.SwscaleVersion())
	log.Printf("AvUtil Version:\t%v", avutil.AvutilVersion())
	log.Printf("AvCodec Version:\t%v", avcodec.AvcodecVersion())
	log.Printf("AvFormat Version:\t%v", avformat.AvformatVersion())
	log.Printf("Resample Version:\t%v", swresample.SwresampleLicense())
}
