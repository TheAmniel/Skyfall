package video

import (
	"bytes"
	"os"
	"os/exec"
	"sync"

	"skyfall/utils"
)

// TODO
type Video struct {
	sync.RWMutex
	raw  []byte
	args map[string]string
}

func New(input []byte) *Video {
	return &Video{
		raw:  input,
		args: make(map[string]string),
	}
}

func (v *Video) FirstFrame() *Video {
	v.AddArg("-ss", "1").
		AddArg("-vframes", "1").
		AddArg("-f", "image2")
	return v
}

func (v *Video) Thumbnail() *Video {
	v.FirstFrame().
		AddArg("-s", "400x280")
	return v
}

func (v *Video) Process() ([]byte, error) {
	toolsPath, err := utils.GetOrCreatePath("tools")
	if err != nil {
		return nil, err
	}

	baseArgs := []string{"-hide_banner", "-v", "error", "-i", "pipe:0"}
	args := v.AllArgs()
	args = append(baseArgs, args...)
	args = append(args, "pipe:1")

	cmd := exec.Command(toolsPath+"ffmpeg", args...)
	stdout := bytes.NewBuffer(nil)
	cmd.Stderr = os.Stderr
	cmd.Stdout = stdout

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	go func() {
		defer stdin.Close()
		stdin.Write(v.raw)
	}()

	if err = cmd.Run(); err != nil {
		return nil, err
	}
	return stdout.Bytes(), nil
}

func (v *Video) AddArg(k, val string) *Video {
	v.Lock()
	v.args[k] = val
	v.Unlock()
	return v
}

func (v *Video) GetArg(k string) (string, bool) {
	v.RLock()
	c, b := v.args[k]
	v.RUnlock()
	return c, b
}

func (v *Video) AllArgs() []string {
	v.RLock()
	args := []string{}
	for k, v := range v.args {
		args = append(args, k, v)
	}
	v.RUnlock()
	return args
}
