package generator

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func DownloadSpringProject(params map[string]string, dest string) error {
	if _, ok := params["bootVersion"]; !ok {
		params["bootVersion"] = "3.4.0"
	}
	baseURL := "https://start.spring.io/starter.zip"
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to download zip: status %d, body: %s", resp.StatusCode, string(body))
	}
	defer resp.Body.Close()

	// Lire le zip en mémoire
	tmpZip, err := os.CreateTemp("", "spring-initializr-*.zip")
	if err != nil {
		return err
	}
	defer os.Remove(tmpZip.Name())
	defer tmpZip.Close()

	_, err = io.Copy(tmpZip, resp.Body)
	if err != nil {
		return err
	}

	if err := unzip(tmpZip.Name(), dest); err != nil {
		return err
	}

	return nil
}

// unzip extrait le zip src directement dans dest sans créer de sous-dossier supplémentaire.
func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		// On enlève le premier dossier du chemin (baseDir)
		parts := strings.SplitN(f.Name, string(os.PathSeparator), 2)
		var relPath string
		if len(parts) == 2 {
			relPath = parts[1]
		} else {
			relPath = parts[0]
		}
		if relPath == "" {
			continue
		}
		fpath := filepath.Join(dest, relPath)

		// ZipSlip check
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", fpath)
		}

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(fpath, f.Mode()); err != nil {
				return err
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), 0755); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

