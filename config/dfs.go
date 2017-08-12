package config

var (
	EVILCORP         map[string]string
	MEDIACLUSTER     map[string]string
	EVILCORP_DIR     string
	MEDIACLUSTER_DIR string
	Architecture     *setUP
)

const (
	windows = "WINDOWS"
	linux   = "LINUX"

	Reports = "Reports"
	Images  = "Images"
	Audios  = "Audios"
	Docs    = "Docs"
)

type setUP struct {
	SELF_OS string
}
