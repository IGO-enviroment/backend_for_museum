/* Таблица статей с информацией */
BEGIN;

create table informations (
    id SERIAL PRIMARY KEY,      -- Уникальный ID

    title VARCHAR(255),
    descriptions TEXT,

    -- Опубликована или нет статья
    is_visible BOOLEAN NOT NULL DEFAULT false,

    published_at TIMESTAMP,     -- Время начало

    preview_url TEXT,           -- Ссылка на картинку превью

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

COMMIT;
