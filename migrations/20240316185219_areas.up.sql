/* Таблица площадок */
BEGIN;

create table areas (
    id SERIAL PRIMARY KEY,      -- Уникальный ID

    named VARCHAR(255),         -- Название площадки
    property VARCHAR(255),      -- Описание

    publish BOOLEAN DEFAULT 0,  -- Видна ли площадка всем польз.

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW()
);

COMMIT;
