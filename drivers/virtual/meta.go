package virtual

import (
	"github.com/alist-org/alist/v3/internal/driver"
	"github.com/alist-org/alist/v3/internal/operations"
)

type Addition struct {
	driver.RootFolderPath
	NumFile     int   `json:"num_file" type:"number" default:"30" required:"true"`
	NumFolder   int   `json:"num_folder" type:"number" default:"30" required:"true"`
	MaxFileSize int64 `json:"max_file_size" type:"number" default:"1073741824" required:"true"`
	MinFileSize int64 `json:"min_file_size"  type:"number" default:"1048576" required:"true"`
}

var config = driver.Config{
	Name:      "Virtual",
	OnlyLocal: true,
	LocalSort: true,
	//NoCache:   true,
}

func New() driver.Driver {
	return &Virtual{}
}

func init() {
	operations.RegisterDriver(config, New)
}
