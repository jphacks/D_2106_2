package usecase

import (
	"mime/multipart"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/jphacks/D_2106_2/domain"
	"github.com/jphacks/D_2106_2/repository"
	"github.com/jphacks/D_2106_2/utils"
)

type ImageUsecase struct {
	ImageRepo      repository.ImageRepository
	CoordinateRepo repository.CoordinateRepository
	S3service      repository.S3service
}

type ImageProp struct {
	Name         string
	CreatedAt    time.Time
	CoordinateId int
}

func (uc *ImageUsecase) UploadImages(albumId int, images []multipart.File, names []string) error {
	var imageProps []*ImageProp

	coordinates, _ := uc.CoordinateRepo.GetCoordinatesByAlbumId(albumId)

	for _, name := range names {
		unixtimeStr := strings.Split(name, ".")[0]
		unixtime, _ := strconv.Atoi(unixtimeStr)

		imageProps = append(imageProps,
			&ImageProp{
				Name:      name,
				CreatedAt: time.Unix(int64(unixtime), 0),
			})
	}

	sort.Slice(imageProps, func(i, j int) bool {
		return imageProps[i].CreatedAt.Before(imageProps[j].CreatedAt)
	})

	sort.Slice(coordinates, func(i, j int) bool {
		return utils.StringToTime(coordinates[i].Timestamp).Before(utils.StringToTime(coordinates[j].Timestamp))
	})

	i := 0
	j := 0

	for i+1 < len(coordinates) && j < len(imageProps) {
		t1 := utils.StringToTime(coordinates[i].Timestamp)
		t2 := utils.StringToTime(coordinates[i+1].Timestamp)

		if (t1.Equal(imageProps[j].CreatedAt) || t1.Before(imageProps[j].CreatedAt)) &&
			t2.After(imageProps[j].CreatedAt) {
			if t2.Sub(imageProps[j].CreatedAt) >= imageProps[j].CreatedAt.Sub(t1) {
				imageProps[j].CoordinateId = coordinates[i].Id
			} else {
				imageProps[j].CoordinateId = coordinates[i+1].Id
			}
			j++
		} else {
			i++
		}

	}

	imageUrls, err := uc.S3service.S3Uploader(images, names)
	if err != nil {
		return err
	}

	var dbInputs []*domain.Image

	for _, url := range imageUrls {
		parsedUrl := strings.Split(url, "/")
		for _, image := range imageProps {
			if image.Name == parsedUrl[len(parsedUrl)-1] {
				dbInputs = append(dbInputs,
					&domain.Image{
						Url:          url,
						AlbumId:      albumId,
						CreatedAt:    image.CreatedAt,
						CoordinateId: image.CoordinateId,
					})
			}
		}
	}
	_, err = uc.ImageRepo.StoreImages(dbInputs)
	if err != nil {
		return err
	}

	return nil
}
