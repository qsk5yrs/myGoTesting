package util

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
)

func SaveStructToJsonFile(data interface{}, filePath string, indent bool) error {
	// 1. 序列化数据为JSON
	var jsonData []byte
	var err error

	if indent {
		jsonData, err = json.MarshalIndent(data, "", "    ")
	} else {
		jsonData, err = json.Marshal(data)
	}
	if err != nil {
		return fmt.Errorf("error marshalling to JSON: %v", err)
	}

	// 2. 确保文件夹存在
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("error creating directory: %v", err)
	}

	// 3. 打开文件（如果文件不存在则创建新文件，如果存在则覆盖）
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// 4. 写入JSON数据到文件
	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}

func JoinURLs(base string, paths []string) (string, error) {
	// 解析基础URL
	baseURL, err := url.Parse(base)
	if err != nil {
		return "", fmt.Errorf("error parsing base URL: %v", err)
	}

	// 创建一个新的相对URL对象，用于逐步拼接路径
	var relativeURL *url.URL

	for _, path := range paths {
		// 解析当前路径部分
		currentPath, err := url.Parse(path)
		if err != nil {
			return "", fmt.Errorf("error parsing path segment '%s': %v", path, err)
		}

		// 如果是第一个路径部分，则直接赋值给relativeURL
		if relativeURL == nil {
			relativeURL = currentPath
		} else {
			// 否则，将当前路径部分拼接到relativeURL上
			relativeURL = relativeURL.ResolveReference(currentPath)
		}
	}

	// 如果没有提供任何路径部分，直接返回基础URL
	if relativeURL == nil {
		return baseURL.String(), nil
	}

	// 拼接基础URL和所有路径部分
	completeURL := baseURL.ResolveReference(relativeURL)

	// 将拼接后的URL转换为字符串
	return completeURL.String(), nil
}
