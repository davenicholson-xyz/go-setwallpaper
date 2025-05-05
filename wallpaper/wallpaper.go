package wallpaper

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var (
	ErrEmptyFilename      = errors.New("empty filename provided")
	ErrFileNotExist       = errors.New("file does not exist")
	ErrUnsupportedDesktop = errors.New("unsupported desktop environment")
	ErrSetCommandFailed   = errors.New("error setting wallpaper with command")
)

var commands = struct {
	Linux map[string]string
	MacOS string
}{
	Linux: map[string]string{
		"gnome": "gsettings set org.gnome.desktop.background picture-uri file://{IMG} && gsettings set org.gnome.desktop.background picture-uri-dark file://{IMG}",
	},
	MacOS: "osascript -e 'tell application \"System Events\" to set picture of every desktop to POSIX file \"{IMG}\"'",
}

func getCommand(desktop string) (string, error) {
	if runtime.GOOS == "darwin" {
		return commands.MacOS, nil
	}

	cmd, exists := commands.Linux["gnome"]
	if !exists {
		return "", fmt.Errorf("%w: %s", ErrUnsupportedDesktop, desktop)
	}

	return cmd, nil
}

func execCommand(cmd string, filepath string) error {

	cmdString := strings.ReplaceAll(cmd, "{IMG}", filepath)

	output, err := exec.Command("sh", "-c", cmdString).CombinedOutput()
	if err != nil {
		fmt.Println(output)
		return fmt.Errorf("%w: %s", ErrSetCommandFailed)
	}
	return nil
}

func Set(filepath string) error {
	if filepath == "" {
		return ErrEmptyFilename
	}

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return fmt.Errorf("%w: %s", ErrFileNotExist, filepath)
	}

	fmt.Println(filepath)

	cmd, err := getCommand("")
	if err != nil {
		return err
	}

	err = execCommand(cmd, filepath)
	if err != nil {
		return err
	}

	return nil
}
