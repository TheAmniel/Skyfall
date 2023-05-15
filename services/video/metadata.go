package video

import (
	"bytes"
	"encoding/json"
	"os/exec"

	"skyfall/utils"
)

func Metadata(input []byte) (map[string]any, error) {
	toolsPath, err := utils.GetOrCreatePath("tools")
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(toolsPath+"ffprobe", "-hide_banner", "-of", "json", "-show_streams", "-show_format", "pipe:0")
	result := bytes.NewBuffer(make([]byte, 0))
	cmd.Stdout = result

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	go func() {
		defer stdin.Close()
		stdin.Write(input)
	}()

	if err = cmd.Run(); err != nil {
		return nil, err
	}

	data := make(map[string]any)
	if err := json.NewDecoder(result).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}
