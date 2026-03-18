# Mood Generator

Пишешь как себя чувствуешь — получаешь цветовую палитру, цитату и плейлист.

Не очень практично? Наверное. Но именно это и делает его интересным.

---

## Что происходит

Пишешь что-то вроде *«устал, хочется тишины»* — за пару секунд получаешь:

- **Цветовую палитру** из 3 цветов под настроение (кликни на цвет — скопируется HEX)
- **Цитату** — иногда точную, иногда неожиданную
- **6 треков** с Last.fm, с прямой ссылкой на YouTube
- **Шкалу энергии** от апатии до эйфории

Фон страницы меняется под твои цвета. История запросов — в боковой панели. Каждый результат живёт по своей ссылке `/mood/:id` — можно шарить.

---

## Стек

**Backend** — Go + Gin + PostgreSQL
**Frontend** — React + Vite + Tailwind CSS
**AI** — Groq API (llama-3.3-70b) — анализирует текст, генерирует палитру и цитату
**Музыка** — Last.fm API — треки по тегу настроения

---

## Запуск

**Нужно:** Go 1.22+, Node.js 18+, PostgreSQL

```bash
# Клонируй
git clone https://github.com/golubev21312-afk/mood-generator.git
cd mood-generator

# Переменные окружения
cp .env.example .env
# заполни .env — там DB_* + GROQ_API_KEY + LASTFM_API_KEY

# База данных
psql -U postgres -c "CREATE DATABASE mood_generator;"
psql -U postgres -d mood_generator -f schema.sql
```

```bash
# Бэкенд (из корня проекта)
cd backend && go run ./cmd/main.go

# Фронтенд (в другом терминале)
cd frontend && npm install && npm run dev
```

Открывай `http://localhost:5173`.

---

## База данных

```sql
CREATE TABLE mood_requests (
  id         SERIAL PRIMARY KEY,
  user_input TEXT NOT NULL,
  mood_label VARCHAR(50),
  energy     INTEGER,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE mood_results (
  id           SERIAL PRIMARY KEY,
  request_id   INTEGER REFERENCES mood_requests(id),
  palette      JSONB,
  quote        TEXT,
  quote_author VARCHAR(100),
  tracks       JSONB,
  created_at   TIMESTAMP DEFAULT NOW()
);
```

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

Без реальных API и БД — всё замокировано.

---

## Деплой

Backend → Railway, Frontend → Vercel.

```bash
# Backend
cd backend
railway init && railway up

# Frontend
cd frontend
npm run build
vercel --prod
```

В Railway добавь переменную `FRONTEND_URL` с доменом Vercel.
В Vercel добавь `VITE_API_URL` с URL Railway-бэкенда.

---

Сделан для практики связки Go + React и работы с AI API.
