package lib

import (
	"cloud.google.com/go/storage"
	"fmt"
	"github.com/satori/go.uuid"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"path"
	"reflect"
)

const publicURL = "https://storage.googleapis.com/%s/%s"

func UpdateBuilder(old, new interface{}) interface{} {
	oldVal := reflect.ValueOf(old).Elem()
	newVal := reflect.ValueOf(new).Elem()

	for i := 0; i < oldVal.NumField(); i++ {
		for j := 0; j < newVal.NumField(); j++ {
			if oldVal.Type().Field(i).Name == newVal.Type().Field(j).Name {
				if newVal.Field(j).Interface() != nil {
					oldVal.Field(i).Set(newVal.Field(j))
				}
			}
		}
	}
	return oldVal.Interface()
}

func UploadFileFromForm(r *http.Request, bucket *storage.BucketHandle, bucketName string) (url string, err error) {
	f, fh, err := r.FormFile("image")
	if err == http.ErrMissingFile {
		return "", nil
	}
	if err != nil {
		return "", err
	}

	// random filename, retaining existing extension.
	name := uuid.NewV4().String() + path.Ext(fh.Filename)

	ctx := context.Background()
	w := bucket.Object(name).NewWriter(ctx)

	// Warning: storage.AllUsers gives public read access to anyone.
	w.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
	w.ContentType = fh.Header.Get("Content-Type")

	// Entries are immutable, be aggressive about caching (1 day).
	w.CacheControl = "public, max-age=86400"

	if _, err := io.Copy(w, f); err != nil {
		return "", err
	}
	if err := w.Close(); err != nil {
		return "", err
	}

	return fmt.Sprintf(publicURL, bucketName, name), nil
}

func UploadMultipleFileFromForm(r *http.Request, bucket *storage.BucketHandle, bucketName string) (url []string, err error) {
	var urls []string

	m := r.MultipartForm

	files := m.File["files"]
	for i, _ := range files {

		name := uuid.NewV4().String() + path.Ext(files[i].Filename)

		ctx := context.Background()
		w := bucket.Object(name).NewWriter(ctx)

		// Warning: storage.AllUsers gives public read access to anyone.
		w.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}
		w.ContentType = files[i].Header.Get("Content-Type")

		// Entries are immutable, be aggressive about caching (1 day).
		w.CacheControl = "public, max-age=86400"

		file, err := files[i].Open()
		defer file.Close()
		if err != nil {
			return urls, err
		}

		if _, err := io.Copy(w, file); err != nil {
			return urls, err
		}
		if err := w.Close(); err != nil {
			return urls, err
		}
		urls = append(urls, fmt.Sprintf(publicURL, bucketName, name))
	}
	return urls, nil
}