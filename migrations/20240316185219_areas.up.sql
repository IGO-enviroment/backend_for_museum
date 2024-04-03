/* Таблица площадок */
BEGIN;

create table areas (
    id SERIAL PRIMARY KEY,      -- Уникальный ID

    named VARCHAR(255),         -- Название площадки
    property VARCHAR(255),      -- Описание

    publish BOOLEAN DEFAULT false,  -- Видна ли площадка всем польз.

    address_value VARCHAR(255), -- Адрес площадки

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW()
);

COMMIT;
