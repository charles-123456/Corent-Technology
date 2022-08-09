// util
package util

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var HTTPClient = &http.Client{Transport: &http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}).DialContext,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
	TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
},
}

func GetWord(data []byte, count int) ([]byte, error) {
	var result []byte
	var err error
	var offset int
	for i := 0; i < count; i++ {
		data = data[offset:]
		offset, result, err = scanWord(data)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func GetPreviousPath(currentFolder string) string {
	dir, _ := filepath.Split(currentFolder)
	return dir[:len(dir)-1]
}

func scanWord(data []byte) (int, []byte, error) {
	return bufio.ScanWords(data, true)
}

func RemoveNewLine(input string) string {
	var result string
	writer := bytes.NewBufferString(result)
	scanner := bufio.NewScanner(bytes.NewBufferString(input))
	for scanner.Scan() {
		fmt.Fprint(writer, scanner.Text())
	}
	return writer.String()
}

func GetFileList(srcPath string) ([]string, error) {
	var streamList []string
	var DirWalker filepath.WalkFunc = func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		streamList = append(streamList, path)
		return nil
	}
	err := filepath.Walk(srcPath, DirWalker)
	if err != nil {
		return nil, err
	}
	if len(streamList) < 2 {
		return nil, errors.New("folder is empty")
	}
	streamList = streamList[1:]
	return streamList, nil
}

func StringToMap(input, seperator string) map[string]string {
	resultMap := make(map[string]string)
	buff := bytes.NewBuffer([]byte(input))
	scanner := bufio.NewScanner(buff)
	for scanner.Scan() {
		currentLine := scanner.Text()
		resultArray := strings.SplitN(currentLine, seperator, 2)
		if len(resultArray) < 2 {
			continue
		}
		key := strings.TrimSpace(resultArray[0])
		value := strings.TrimSpace(resultArray[1])
		if key != "" {
			resultMap[key] = value
		}
	}
	return resultMap
}
