/* Таблица площадок */
BEGIN;

create table areas (
    id SERIAL PRIMARY KEY,      -- Уникальный ID

    name VARCHAR(255),         -- Название площадки
    description VARCHAR(255),      -- Описание

    publish BOOLEAN DEFAULT false,  -- Видна ли площадка всем польз.

    address_value VARCHAR(255), -- Адрес площадки

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

COMMIT;
