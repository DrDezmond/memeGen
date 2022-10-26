package apiserver

import (
	"bytes"
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

	return func(w http.ResponseWriter, r *http.Request) {
		var data GeneratorData

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		s.generator.InitGeneratorValues(data.Texts, data.Orientation, data.FontSize)
		img := s.generator.GenerateImages()
		w.Header().Set("Content-Type", "image/jpeg")

		buf := new(bytes.Buffer)
		err := jpeg.Encode(buf, img, nil)
		if err != nil {
			http.Error(w, "Could not encode image", http.StatusBadRequest)
		}

		s.generator.Images = nil

		w.Write(buf.Bytes())
	}
}
