/**
* @Author: yibo_LastDay
* @Date: 2022/10/15 11:31
 */

package util

import (
	"bufio"
	"io"
	"os"
)

func IsDirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}

func MakeDir(path string) {
	os.Mkdir(path, os.ModePerm)
}

func CheckFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func CreateNewFile(fileName string, content []byte) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0600)
	if err != nil {
		return err
	} else {
		_, err = f.Write(content)
		if err != nil {
			return err
		}
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func DeleteFile(filename string) error {
	return os.Remove(filename)
}

func ReadFile(fileName string) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	str, err := r.ReadString('\n')
	if err != io.EOF {
		return "", err
	}
	return str, nil
}

//EOF
