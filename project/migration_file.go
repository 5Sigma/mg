package project

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

type (
	MigrationFile struct {
		FilePath string
	}
)

func (mf *MigrationFile) ReadContents() (string, error) {
	buf := bytes.NewBuffer(nil)
	f, err := os.Open(mf.FilePath)
	if err != nil {
		return "", err
	}
	io.Copy(buf, f)
	f.Close()
	return string(buf.Bytes()), nil
}

func (f *MigrationFile) Edit() error {
	cmd := exec.Command("vim", f.FilePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	return err
}

func (f *MigrationFile) GetBlock(block string) (string, error) {
	file, err := os.Open(f.FilePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lineBuffer []string
	var inBlock = false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		blockRxp, _ := regexp.Compile("[\\-]{4}\\s?(?P<name>[A-Za-z]+)")
		if blockRxp.MatchString(line) {
			matches := blockRxp.FindStringSubmatch(strings.ToUpper(line))
			blockName := matches[1]
			if blockName == block {
				inBlock = true
				continue
			}
			if blockName != block {
				inBlock = false
				continue
			}
		} else {
			if inBlock == true {
				lineBuffer = append(lineBuffer, line)
				continue
			}
		}
	}
	return strings.Join(lineBuffer, "\n"), nil
}

func (f *MigrationFile) GetUp() (string, error) {
	return f.GetBlock("UP")
}

func (f *MigrationFile) GetDown() (string, error) {
	return f.GetBlock("DOWN")
}

func (f *MigrationFile) Filename() string {
	return filepath.Base(f.FilePath)
}

func (f *MigrationFile) Version() string {
	tokens := strings.Split(f.Filename(), "_")
	return tokens[0]
}
