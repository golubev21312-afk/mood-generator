# Mood Generator

Ты пишешь как себя чувствуешь — а приложение превращает это в цветовую палитру, цитату и плейлист на YouTube.

Не очень практично? Наверное. Но именно это и делает его интересным.

---

## Как это работает

Пишешь что-то вроде *«устал, хочется тишины»* — и за пару секунд получаешь:

- **Цветовую палитру** — цвета, которые соответствуют твоему настроению
- **Цитату** — иногда точную, иногда неожиданную
- **6 треков** — подобранных по настроению, с прямой ссылкой на YouTube

Фоновый градиент страницы тоже меняется под твои цвета. Мелочь, но приятно.

---

## Стек

**Backend** — Go + Gin + PostgreSQL
**Frontend** — React + Vite + Tailwind CSS
**AI** — Groq API (llama-3.3-70b) — анализирует текст и генерирует палитру с цитатой
**Музыка** — Last.fm API — подбирает треки по тегу настроения

---

## Запуск

**Что нужно:** Go 1.21+, Node.js 18+, PostgreSQL

```bash
# 1. Клонируй репо
git clone <repo-url>
cd mood-generator

# 2. Настрой переменные окружения
cp .env.example backend/.env
# Заполни backend/.env своими ключами

# 3. Создай базу данных
psql -U postgres -c "CREATE DATABASE mood_generator;"
psql -U postgres -d mood_generator -c "
  CREATE TABLE mood_requests (
    id SERIAL PRIMARY KEY,
    user_input TEXT NOT NULL,
    mood_label VARCHAR(50),
    energy INTEGER,
    created_at TIMESTAMP DEFAULT NOW()
  );
  CREATE TABLE mood_results (
    id SERIAL PRIMARY KEY,
    request_id INTEGER REFERENCES mood_requests(id),
    palette JSONB,
    quote TEXT,
    quote_author VARCHAR(100),
    tracks JSONB,
    created_at TIMESTAMP DEFAULT NOW()
  );"

# 4. Запусти бэкенд
cd backend && go run ./cmd/main.go

# 5. В другом терминале — фронтенд
cd frontend && npm install && npm run dev
```

Открывай `http://localhost:5173` и пробуй.

---

## Ключи API

| Сервис | Где взять | Цена |
|--------|-----------|------|
| Groq | [console.groq.com](https://console.groq.com) | Бесплатно |
| Last.fm | [last.fm/api](https://www.last.fm/api/account/create) | Бесплатно |

---

## Тесты

```bash
cd backend
go test ./internal/... -v
```

Тесты не требуют реальных API ключей — всё замокировано.

---

Проект сделан для изучения связки Go + React и работы с AI API. Если хочешь что-то улучшить — PR приветствуется.
