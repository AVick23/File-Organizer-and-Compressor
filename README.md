# File Organizer and Compressor

Этот проект на языке Go считывает директорию, сортирует файлы по их расширениям и создаёт zip-архив с оригинальными файлами.

## Установка

### Для Windows

1. Убедитесь, что у вас установлен [Go](https://golang.org/dl/).
2. Клонируйте репозиторий:
    ```sh
    git clone https://github.com/AVick23/File-Organizer-and-Compressor.git
    cd File-Organizer-and-Compressor
    ```
3. Соберите проект:
    ```sh
    go build
    ```
4. Запустите исполняемый файл:
    ```sh
    ./File-Organizer-and-Compressor.exe
    ```

### Для macOS

1. Убедитесь, что у вас установлен [Go](https://golang.org/dl/).
2. Клонируйте репозиторий:
    ```sh
    git clone https://github.com/AVick23/File-Organizer-and-Compressor.git
    cd File-Organizer-and-Compressor
    ```
3. Соберите проект:
    ```sh
    go build
    ```
4. Запустите исполняемый файл:
    ```sh
    ./File-Organizer-and-Compressor
    ```

### Для Linux

1. Убедитесь, что у вас установлен [Go](https://golang.org/dl/).
2. Клонируйте репозиторий:
    ```sh
    git clone https://github.com/AVick23/File-Organizer-and-Compressor.git
    cd File-Organizer-and-Compressor
    ```
3. Соберите проект:
    ```sh
    go build
    ```
4. Запустите исполняемый файл:
    ```sh
    ./File-Organizer-and-Compressor
    ```

## Использование

После запуска исполняемого файла следуйте инструкциям на экране:

1. Введите путь к директории, которую хотите обработать.
2. Подтвердите структуру сортировки.
3. Файлы будут скопированы в новую директорию `sort`, а оригинальные файлы сжаты в zip-архив.

### Пример

```sh
Введите путь к директории:
~/Documents/files

Предварительный просмотр сортировки файлов:
sort
├── txt
│   ├── file1.txt
│   ├── file2.txt
├── jpg
│   ├── image1.jpg
│   ├── image2.jpg
└── zip
    └── archive.zip

Вы хотите подтвердить новую структуру директории? (y/n): y
Идёт сортировка и сжатие файлов в zip архив.
Файлы успешно скопированы и оригинальные файлы сжаты в архив: ~/Documents/files/original_files.zip
