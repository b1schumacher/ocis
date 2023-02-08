package svc

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"path"
	"path/filepath"
)

var (
	errInvalidThemeConfig       = errors.New("invalid themes config")
	_themesConfigPath           = filepath.FromSlash("themes/owncloud/theme.json")
	_allowedExtensionMediatypes = map[string]string{
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".png":  "image/png",
		".gif":  "image/gif",
	}
)

// UploadLogo implements the endpoint to upload a custom logo for the oCIS instance.
func (p Web) UploadLogo(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("logo")
	if err != nil {
		if errors.Is(err, http.ErrMissingFile) {
			w.WriteHeader(http.StatusBadRequest)
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	mediatype := fileHeader.Header.Get("Content-Type")
	if !allowedFiletype(fileHeader.Filename, mediatype) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fp := filepath.Join("branding", filepath.Join("/", fileHeader.Filename))
	err = p.storeAsset(fp, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = p.updateLogoThemeConfig(fp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p Web) storeAsset(name string, asset io.Reader) error {
	dst, err := p.fs.Create(name)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, asset)
	return err
}

func (p Web) updateLogoThemeConfig(logoPath string) error {
	f, err := p.fs.Open(_themesConfigPath)
	if err == nil {
		defer f.Close()
	}

	// This decoding of the themes.json file is not optimal. If we need to decode it for other
	// usecases as well we should consider decoding to a struct.
	var m map[string]interface{}
	_ = json.NewDecoder(f).Decode(&m)

	webCfg, ok := m["web"].(map[string]interface{})
	if !ok {
		return errInvalidThemeConfig
	}

	defaultCfg, ok := webCfg["default"].(map[string]interface{})
	if !ok {
		return errInvalidThemeConfig
	}

	logoCfg, ok := defaultCfg["logo"].(map[string]interface{})
	if !ok {
		return errInvalidThemeConfig
	}

	logoCfg["login"] = logoPath
	logoCfg["topbar"] = logoPath

	defaultDarkCfg, ok := webCfg["default-dark"].(map[string]interface{})
	if !ok {
		return errInvalidThemeConfig
	}

	logoDarkCfg, ok := defaultDarkCfg["logo"].(map[string]interface{})
	if !ok {
		return errInvalidThemeConfig
	}

	logoDarkCfg["login"] = logoPath
	logoDarkCfg["topbar"] = logoPath

	dst, err := p.fs.Create(_themesConfigPath)
	if err != nil {
		return err
	}

	return json.NewEncoder(dst).Encode(m)
}

func allowedFiletype(filename, mediatype string) bool {
	ext := path.Ext(filename)

	// Check if we allow that extension and if the mediatype matches the extension
	mt, ok := _allowedExtensionMediatypes[ext]
	return ok && mt == mediatype
}
