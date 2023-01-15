package main

import (
	"flag"
	"io"
	"os"
	"strings"

	"code.cyb3r.social/skat/goDataCollector/app"
	log "github.com/sirupsen/logrus"
)

func setupLog(logFilePath string) {

	logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.WithError(err).Errorf("Failed to create logfile %s\n", logFilePath)
	} else {
		multi := io.MultiWriter(logFile, os.Stdout)
		log.SetOutput(multi)
	}

	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		PadLevelText:    true,
	})

}

func exit() {
	log.Infof("exit\n")
}

func main() {
	defer exit()
	var lLevel string
	var daemon bool
	var help bool
	var logToFile bool
	var devLocal bool

	flag.StringVar(&lLevel, "l", "Info", "default (Info) possibel values are Info, Debug, Trace")
	flag.BoolVar(&daemon, "d", false, "run as daemon")
	flag.BoolVar(&help, "h", false, "print help")
	flag.BoolVar(&logToFile, "f", false, "log to file")
	flag.BoolVar(&devLocal, "dl", false, "load local.config.json")

	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		PadLevelText:    true,
	})

	switch strings.ToLower(lLevel) {
	case "info":
		log.SetLevel(log.InfoLevel)
		log.Infof("Loglevel set to Info.")
	case "trace":
		log.SetLevel(log.TraceLevel)
		log.SetReportCaller(true)
		log.Infof("Loglevel set to Trace.")
	case "debug":
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
		log.Infof("Loglevel set to Debug.")
	default:
		log.SetLevel(log.InfoLevel)
		log.Warningf("Provided loglevel %s not valid using Info.\n", lLevel)
	}

	if logToFile {
		setupLog("./appData/Logs/today.log")
	}

	if daemon {
		app.Daemonize(devLocal)
	} else {
		app.Run(devLocal)
	}

}
