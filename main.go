package main

import (
	"archive/zip"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func copyFile(src string, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return destinationFile.Sync()
}

func zipFiles(zipFileName, filePath string, fileMap map[string][]string) error {
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, files := range fileMap {
		for _, fileName := range files {
			if fileName == "original_files.zip" {
				continue
			}

			filePath := filepath.Join(filePath, fileName)
			fileToZip, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer fileToZip.Close()

			zipFileWriter, err := zipWriter.Create(fileName)
			if err != nil {
				return err
			}

			if _, err = io.Copy(zipFileWriter, fileToZip); err != nil {
				return err
			}
		}
	}

	return nil
}

func getUserChoice(prompt string, reader *bufio.Reader) string {
	for {
		fmt.Print(prompt)
		choice, _ := reader.ReadString('\n')
		choice = strings.ToLower(strings.TrimSpace(choice))

		if choice == "y" || choice == "n" || choice == "d" {
			return choice
		}

		fmt.Println("Неверный ввод. Пожалуйста, введите 'y' (перезаписать), 'n' (отменить) или 'd' (создать дубликаты).")
	}
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите путь к директории:")

	filePath, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	filePath = strings.TrimSpace(filePath)
	filePath = strings.Trim(filePath, "\"'")

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("Указанная директория не существует:", filePath)
		return
	}

	sortFolder := filepath.Join(filePath, "sort")
	zipFileName := filepath.Join(filePath, "original_files.zip")

	sortFolderExists := false
	zipFileExists := false

	if _, err := os.Stat(sortFolder); err == nil {
		sortFolderExists = true
	}

	if _, err := os.Stat(zipFileName); err == nil {
		zipFileExists = true
	}

	if sortFolderExists || zipFileExists {
		choice := getUserChoice("Папка 'sort' или файл 'original_files.zip' уже существуют. Перезаписать (y), отменить (n) или создать дубликаты (d)? ", reader)

		switch choice {
		case "n":
			fmt.Println("Операция отменена.")
			return
		case "d":
			sortFolder = filepath.Join(filePath, "sort_copy")
			zipFileName = filepath.Join(filePath, "original_files_copy.zip")
		}
	}

	dir, err := os.ReadDir(filePath)
	if err != nil {
		log.Fatal(err)
	}

	fileMap := make(map[string][]string)

	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		if file.Name() == "original_files.zip" {
			continue
		}

		expansion := strings.ToLower(filepath.Ext(file.Name()))
		if expansion == "" {
			expansion = "no_expansion"
		}

		fileMap[expansion] = append(fileMap[expansion], file.Name())
	}

	fmt.Println("Предварительный просмотр сортировки файлов:")
	fmt.Println("sort")

	expansionCount := len(fileMap)
	currentExpansion := 0

	for exp, files := range fileMap {
		currentExpansion++

		if currentExpansion == expansionCount {
			fmt.Printf("└── Папка: %s\n", exp[1:])
		} else {
			fmt.Printf("├── Папка: %s\n", exp[1:])
		}

		fileCount := len(files)
		currentFile := 0

		for _, file := range files {
			currentFile++
			if currentFile == fileCount {
				if currentExpansion == expansionCount {
					fmt.Printf("    └── %s\n", file)
				} else {
					fmt.Printf("│   └── %s\n", file)
				}
			} else {
				if currentExpansion == expansionCount {
					fmt.Printf("    ├── %s\n", file)
				} else {
					fmt.Printf("│   ├── %s\n", file)
				}
			}
		}
	}

	fmt.Print("Вы хотите подтвердить новую структуру директории? (y/n): ")
	confirmation, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	confirmation = strings.ToLower(strings.TrimSpace(confirmation))

	if confirmation == "n" {
		fmt.Println("Сортировка отменена.")
		return
	} else {
		fmt.Println("Идёт сортировка и сжатие файлов в zip архив.")
	}

	if err := os.MkdirAll(sortFolder, os.ModePerm); err != nil {
		log.Fatal("Ошибка при создании папки 'sort':", err)
	}

	for exp, files := range fileMap {
		targetFolder := filepath.Join(sortFolder, exp[1:])
		if err := os.MkdirAll(targetFolder, os.ModePerm); err != nil {
			log.Fatal("Ошибка создания папки для расширения", err)
		}

		for _, file := range files {
			sourceFilePath := filepath.Join(filePath, file)
			targetFilePath := filepath.Join(targetFolder, file)
			if err := copyFile(sourceFilePath, targetFilePath); err != nil {
				log.Fatal("Ошибка при копировании файла:", err)
			}
		}
	}

	if err := zipFiles(zipFileName, filePath, fileMap); err != nil {
		log.Fatal(err)
	}

	for _, files := range fileMap {
		for _, file := range files {

			if file == "original_files.zip" {
				continue
			}

			sourceFilePath := filepath.Join(filePath, file)
			err := os.Remove(sourceFilePath)
			if err != nil {
				log.Printf("Ошибка при удалении файла %s: %v\n", file, err)
			}
		}
	}

	fmt.Println("Файлы успешно скопированы и оригинальные файлы сжаты в архив:", zipFileName)
}
