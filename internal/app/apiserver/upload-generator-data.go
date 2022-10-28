package apiserver

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"image/jpeg"
	"net/http"
)

func (s *APIserver) HandleGeneratorDataUpload() http.HandlerFunc {
	type GeneratorData struct {
		Orientation string
		Texts       map[int][]string
		FontSize    float64
	}

	type Response struct {
		Image string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var data GeneratorData

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		s.generator.InitGeneratorValues(data.Texts, data.Orientation, data.FontSize)
		img := s.generator.GenerateImages()

		buf := new(bytes.Buffer)
		err := jpeg.Encode(buf, img, nil)
		if err != nil {
			http.Error(w, "Could not encode image", http.StatusBadRequest)
		}

		mimeType := http.DetectContentType(buf.Bytes())
		imgBase64Str := base64.StdEncoding.EncodeToString(buf.Bytes())

		switch mimeType {
		case "image/jpeg":
			imgBase64Str = "data:image/jpeg;base64," + imgBase64Str
		case "image/png":
			imgBase64Str = "data:image/png;base64," + imgBase64Str
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/text")
		w.Write([]byte(imgBase64Str))

		// res := Response{Image: imgBase64Str}

		// jData, err := json.Marshal(res)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		// w.Header().Set("Content-Type", "application/json")
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		// w.Write(jData)

		s.generator.Images = nil

		// w.Header().Set("Content-Type", "image/jpeg")
		// w.Write(buf.Bytes())
	}
}
