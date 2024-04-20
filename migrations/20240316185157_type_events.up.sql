/* Таблица типов мероприятий */
BEGIN;

create table type_events (
    id SERIAL PRIMARY KEY,      -- Уникальный ID

    name VARCHAR(255),         -- Название типа мероприятия
    description VARCHAR(255),      -- Описание

    publish BOOLEAN DEFAULT false,  -- Виден ли тип всем польз.

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

COMMIT;
