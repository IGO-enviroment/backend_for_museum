/* Таблица тегов */
BEGIN;

create table tags (
    id SERIAL PRIMARY KEY,      -- Уникальный ID

    name VARCHAR(255),          -- Название тега
    description VARCHAR(255),   -- Описание тега

    group_name VARCHAR(255),         -- Группа тега

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

COMMIT;
