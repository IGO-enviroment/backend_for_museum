/* Таблица тегов */
BEGIN;

create table tags (
    id SERIAL PRIMARY KEY,      -- Уникальный ID

    named VARCHAR(255),         -- Название тега
    property VARCHAR(255),      -- Описание тега

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW()
);

COMMIT;
