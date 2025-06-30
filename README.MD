# Image Server

Простой сервер-галерея изображений на Go с хранением путей в SQLite и отдачей статических файлов.

## Структура проекта

```
image-server/
├── cmd/
│   └── server/
│       └── main.go         # Точка входа
├── internal/
│   ├── db/                 # Работа с БД
│   ├── images/             # Сканирование изображений
│   └── web/                # HTTP-обработчики и шаблоны
│       └── templates/
│           └── index.html  # Главная страница
├── static/                 # Папка с изображениями
├── data/                   # Папка с базой данных
│   └── images.db           # SQLite база
├── go.mod
├── go.sum
└── README.md
```

## Требования
- Go 1.18+
- SQLite (через modernc.org/sqlite, устанавливается автоматически)

## Быстрый старт

### 1. Клонируйте репозиторий и перейдите в корень проекта

**PowerShell:**
```powershell
git clone github.com/112Alex/image-server
cd image-server
```

**Linux (Ubuntu):**
```bash
git clone github.com/112Alex/image-server
cd image-server
```

### 2. Установите зависимости

```bash
go mod download
```

### 3. Запустите сервер

**PowerShell:**
```powershell
go run .\cmd\server\main.go
```

**Linux (Ubuntu):**
```bash
go run ./cmd/server/main.go
```

Сервер будет доступен по адресу: http://localhost:8080

### 4. Добавьте изображения

Положите файлы изображений (`.jpg`, `.jpeg`, `.png`, `.webp`, `.gif`) в папку `static/` в корне проекта. После перезапуска сервера новые изображения автоматически попадут в базу и появятся в галерее.

## Сброс базы данных

Если нужно пересоздать базу (например, после смены структуры хранения):

**PowerShell:**
```powershell
Remove-Item .\data\images.db
```

**Linux (Ubuntu):**
```bash
rm ./data/images.db
```

После этого перезапустите сервер.

## Примечания
- Сервер ищет шаблон `internal/web/templates/index.html` относительно корня проекта. Запускайте сервер из корня!
- Все пути к изображениям в базе — только имена файлов. Для отдачи используется маршрут `/static/<имя_файла>`.
- Для production рекомендуется использовать отдельный reverse proxy (например, nginx) для отдачи статики.
