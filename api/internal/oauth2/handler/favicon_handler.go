package handler

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"net/http"
	"strings"

	_ "image/jpeg"

	"golang.org/x/image/draw"
	_ "golang.org/x/image/webp"

	"github.com/locksmithhq/locksmith/api/internal/oauth2/contract"
)

type faviconHandler struct {
	getClientByClientIDRepository contract.GetClientByClientIDRepository
	getLoginByClientIDRepository  contract.GetLoginByClientIDRepository
}

func (h *faviconHandler) Execute(w http.ResponseWriter, r *http.Request) {
	clientID := r.URL.Query().Get("client_id")
	if clientID == "" {
		http.NotFound(w, r)
		return
	}

	client, err := h.getClientByClientIDRepository.Execute(r.Context(), clientID)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	login, err := h.getLoginByClientIDRepository.Execute(r.Context(), client.ID)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	faviconVal, _ := login.FaviconURL.Value()
	faviconURL, _ := faviconVal.(string)

	if faviconURL == "" || !strings.HasPrefix(faviconURL, "data:") {
		http.NotFound(w, r)
		return
	}

	// Parse data URL: data:<mime>;base64,<data>
	rest := strings.TrimPrefix(faviconURL, "data:")
	parts := strings.SplitN(rest, ",", 2)
	if len(parts) != 2 {
		http.NotFound(w, r)
		return
	}

	encoded := parts[1]

	data, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Decode image, resize to 512x512, re-encode as PNG
	src, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	dst := image.NewRGBA(image.Rect(0, 0, 512, 512))
	draw.BiLinear.Scale(dst, dst.Bounds(), src, src.Bounds(), draw.Over, nil)

	var buf bytes.Buffer
	if err := png.Encode(&buf, dst); err != nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "public, max-age=86400")
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())
}

func NewFaviconHandler(
	getClientByClientIDRepository contract.GetClientByClientIDRepository,
	getLoginByClientIDRepository contract.GetLoginByClientIDRepository,
) contract.FaviconHandler {
	return &faviconHandler{
		getClientByClientIDRepository: getClientByClientIDRepository,
		getLoginByClientIDRepository:  getLoginByClientIDRepository,
	}
}
