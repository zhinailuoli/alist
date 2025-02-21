package virtual

import (
	"context"
	"github.com/alist-org/alist/v3/internal/driver"
	"github.com/alist-org/alist/v3/internal/errs"
	"github.com/alist-org/alist/v3/internal/model"
	"github.com/alist-org/alist/v3/pkg/utils"
	"github.com/alist-org/alist/v3/pkg/utils/random"
	"github.com/pkg/errors"
	"io"
	"time"
)

type Virtual struct {
	model.Storage
	Addition
}

func (d *Virtual) Config() driver.Config {
	return config
}

func (d *Virtual) Init(ctx context.Context, storage model.Storage) error {
	d.Storage = storage
	err := utils.Json.UnmarshalFromString(storage.Addition, &d.Addition)
	if err != nil {
		return errors.Wrap(err, "error while unmarshal addition")
	}
	return nil
}

func (d *Virtual) Drop(ctx context.Context) error {
	return nil
}

func (d *Virtual) GetStorage() model.Storage {
	return d.Storage
}

func (d *Virtual) GetAddition() driver.Additional {
	return d.Addition
}

func (d *Virtual) List(ctx context.Context, dir model.Obj) ([]model.Obj, error) {
	var res []model.Obj
	for i := 0; i < d.NumFile; i++ {
		res = append(res, model.Object{
			Name:     random.String(10),
			Size:     random.RangeInt64(d.MinFileSize, d.MaxFileSize),
			IsFolder: false,
			Modified: time.Now(),
		})
	}
	for i := 0; i < d.NumFolder; i++ {
		res = append(res, model.Object{
			Name:     random.String(10),
			Size:     0,
			IsFolder: true,
			Modified: time.Now(),
		})
	}
	return res, nil
}

func (d *Virtual) Link(ctx context.Context, file model.Obj, args model.LinkArgs) (*model.Link, error) {
	return &model.Link{
		Data: io.NopCloser(io.LimitReader(random.Rand, file.GetSize())),
	}, nil
}

func (d *Virtual) MakeDir(ctx context.Context, parentDir model.Obj, dirName string) error {
	return nil
}

func (d *Virtual) Move(ctx context.Context, srcObj, dstDir model.Obj) error {
	return nil
}

func (d *Virtual) Rename(ctx context.Context, srcObj model.Obj, newName string) error {
	return nil
}

func (d *Virtual) Copy(ctx context.Context, srcObj, dstDir model.Obj) error {
	return nil
}

func (d *Virtual) Remove(ctx context.Context, obj model.Obj) error {
	return nil
}

func (d *Virtual) Put(ctx context.Context, dstDir model.Obj, stream model.FileStreamer, up driver.UpdateProgress) error {
	return nil
}

func (d *Virtual) Other(ctx context.Context, data interface{}) (interface{}, error) {
	return nil, errs.NotSupport
}

var _ driver.Driver = (*Virtual)(nil)
