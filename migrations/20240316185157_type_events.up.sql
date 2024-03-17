/* Таблица типов мероприятий */
BEGIN;

create table type_events (
    id SERIAL PRIMARY KEY,      -- Уникальный ID

    named VARCHAR(255),         -- Название типа мероприятия
    property VARCHAR(255),      -- Описание

    publish BOOLEAN DEFAULT 0,  -- Виден ли тип всем польз.

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW()
);

COMMIT;
