/* Таблица популярных фильтров */
BEGIN;

create table filters (
    id SERIAL PRIMARY KEY,          -- Уникальный ID

    named VARCHAR(255),             -- Название
    property VARCHAR(255),          -- Описание

    publish BOOLEAN DEFAULT false,      -- Видна ли площадка всем польз.    

    order_value INTEGER,            -- Порядок вывода

    user_id INTEGER,                -- ID админа, который создал фильтр

    options JSON NOT NULL DEFAULT '{}'::jsonb,   -- Настройки фильтра

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW()
);

COMMIT;
