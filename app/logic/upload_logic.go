package logic

import (
	"github.com/kataras/iris/v12"
	"github.com/nfnt/resize"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

type UploadLogic interface {
	GetUploads() []UploadedFile
	UpFile(ctx iris.Context) error
}

type uploadLogic struct{}

func NewUploadLogic() UploadLogic {
	return &uploadLogic{}
}

func (u uploadLogic) GetUploads() []UploadedFile {
	files := scanUploads(uploadsDir)

	return files.items
}

func (u uploadLogic) UpFile(ctx iris.Context) error {
	file, info, err := ctx.FormFile("file")
	if err != nil {
		ctx.Application().Logger().Warnf("Error while uploading: %v", err.Error())
		return err
	}
	defer file.Close()
	uploader := new(uploadedFiles)

	fname := info.Filename

	// Create a file with the same name
	// assuming that you have a folder named 'uploads'
	out, err := os.OpenFile(uploadsDir+fname, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		ctx.Application().Logger().Warnf("Error while preparing the new file: %v", err.Error())
		return err
	}
	defer out.Close()

	_, _ = io.Copy(out, file)
	// optionally, add that file to the list in order to be visible when refresh.
	uploadedFile := uploader.add(UploadedFile{
		Name: fname,
		Size: info.Size,
	})

	go uploader.createThumbnail(uploadedFile)

	return nil
}

const uploadsDir = "./public/uploads/"

type UploadedFile struct {
	// {name: "", size: } are the dropzone's only requirements.
	Name string `json:"name"`
	Size int64  `json:"size"`
}

type uploadedFiles struct {
	dir   string
	items []UploadedFile
	mu    sync.RWMutex // slices are safe but RWMutex is a good practise for you.
}

func scanUploads(dir string) *uploadedFiles {
	f := new(uploadedFiles)

	lindex := dir[len(dir)-1]
	if lindex != os.PathSeparator && lindex != '/' {
		dir += string(os.PathSeparator)
	}

	// create directories if necessary
	// and if, then return empty uploaded files; skipping the scan.
	if err := os.MkdirAll(dir, os.FileMode(0666)); err != nil {
		return f
	}

	// otherwise scan the given "dir" for files.
	f.scan(dir)
	return f
}

func (f *uploadedFiles) scan(dir string) {
	f.dir = dir
	_ = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// if it's directory or a thumbnail we saved earlier, skip it.
		if info.IsDir() || strings.HasPrefix(info.Name(), "thumbnail_") {
			return nil
		}

		f.add(UploadedFile{
			Name: info.Name(),
			Size: info.Size(),
		})
		return nil
	})
}

func (f *uploadedFiles) add(uf UploadedFile) UploadedFile {
	f.mu.Lock()
	f.items = append(f.items, uf)
	f.mu.Unlock()

	return uf
}

func (f *uploadedFiles) createThumbnail(uf UploadedFile) {
	file, err := os.Open(path.Join(f.dir, uf.Name))
	if err != nil {
		return
	}
	defer file.Close()

	name := strings.ToLower(uf.Name)

	out, err := os.OpenFile(f.dir+"thumbnail_"+uf.Name,
		os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer out.Close()

	if strings.HasSuffix(name, ".jpg") {
		// decode jpeg into image.Image
		img, err := jpeg.Decode(file)
		if err != nil {
			return
		}
		// write new image to file
		resized := resize.Thumbnail(180, 180, img, resize.Lanczos3)
		_ = jpeg.Encode(out, resized,
			&jpeg.Options{Quality: jpeg.DefaultQuality})

	} else if strings.HasSuffix(name, ".png") {
		img, err := png.Decode(file)
		if err != nil {
			return
		}
		// write new image to file
		resized := resize.Thumbnail(180, 180, img, resize.Lanczos3) // slower but better res
		_ = png.Encode(out, resized)
	}
	// and so on... you got the point, this code can be simplify, as a practise.
}
