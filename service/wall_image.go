package service

import (
	"bytes"
	"context"
	"image"
	"image/color"
	"image/jpeg"
	"math"

	"github.com/xescugc/chaoswall/hold"
	"gocv.io/x/gocv"
	"golang.org/x/xerrors"
)

const MinimumArea float64 = 100

// PreviewWallImage will read the img and use CV to detect the holds and return a new image with the
// holds marked so we know which holds would be used
func (s *service) PreviewWallImage(ctx context.Context, gCan string, img []byte) ([]byte, error) {
	originIM, err := newImage(img)
	if err != nil {
		return nil, xerrors.Errorf("could not initialize image: %w", err)
	}

	holds := getHolds(originIM)

	for _, h := range holds {
		rec := image.Rect(h.X-h.Size, h.Y-h.Size, h.X+h.Size, h.Y+h.Size)
		gocv.Rectangle(&originIM, rec, color.RGBA{240, 52, 52, 0}, 2)
	}

	auxIMG, err := originIM.ToImage()
	if err != nil {
		return nil, xerrors.Errorf("could not convert to image: %w", err)
	}

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, auxIMG, nil)
	if err != nil {
		return nil, xerrors.Errorf("could not encode image: %w", err)
	}

	return buf.Bytes(), nil
}

func newImage(img []byte) (gocv.Mat, error) {
	originIM, err := gocv.IMDecode(img, gocv.IMReadColor)
	if err != nil {
		return gocv.Mat{}, xerrors.Errorf("could not decode image: %w", err)
	}

	if originIM.Empty() {
		return gocv.Mat{}, xerrors.Errorf("empty image")
	}

	return originIM, nil
}

func getHolds(originIM gocv.Mat) []hold.Hold {
	gauBlurIM := gocv.NewMat()
	defer gauBlurIM.Close()

	gocv.GaussianBlur(originIM, &gauBlurIM, image.Point{X: 5, Y: 5}, 0, 0, gocv.BorderConstant)

	ctvColorIM := gocv.NewMat()
	defer ctvColorIM.Close()

	gocv.CvtColor(gauBlurIM, &ctvColorIM, gocv.ColorBGRToGray)

	thresholdIM := gocv.NewMat()
	defer thresholdIM.Close()

	otsu := gocv.Threshold(ctvColorIM, &thresholdIM, 0.0, 255.0, gocv.ThresholdBinary+gocv.ThresholdOtsu)

	cannyIM := gocv.NewMat()
	defer cannyIM.Close()

	gocv.Canny(gauBlurIM, &cannyIM, otsu, otsu*2)

	contours := gocv.FindContours(cannyIM, gocv.RetrievalList, gocv.ChainApproxSimple)

	newContours := make([][]image.Point, 0)
	for _, c := range contours {
		area := gocv.ContourArea(c)
		if area < MinimumArea {
			continue
		}

		hull := gocv.NewMat()
		gocv.ConvexHull(c, &hull, true, false)
		points := make([]image.Point, 0, hull.Rows())
		for i := 0; i < hull.Rows(); i++ {
			points = append(points, c[hull.GetIntAt(i, 0)])
		}
		newContours = append(newContours, points)
	}

	maskIM := gocv.NewMatWithSize(originIM.Rows(), originIM.Cols(), originIM.Type())
	defer maskIM.Close()

	gocv.DrawContours(&maskIM, newContours, -1, color.RGBA{255, 255, 255, 1}, -1)

	sbdp := gocv.NewSimpleBlobDetectorParams()

	sbd := gocv.NewSimpleBlobDetectorWithParams(sbdp)
	defer sbd.Close()

	kps := sbd.Detect(maskIM)
	holds := make([]hold.Hold, len(kps))
	for i, kp := range kps {
		x := int(kp.X)
		y := int(kp.Y)

		size := int(math.Ceil(kp.Size))

		holds[i] = hold.Hold{
			X:    x,
			Y:    y,
			Size: size,
		}
	}

	return holds
}
