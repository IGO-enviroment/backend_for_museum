/* Таблица мероприятий */
BEGIN;

create table events (
    id SERIAL PRIMARY KEY,                  -- Уникальный ID

    title VARCHAR(255),                     -- Название

    publish BOOLEAN NOT NULL DEFAULT false,      -- Опубликовано или нет событие

    ticket_count INTEGER DEFAULT 0,          -- Общее количество доступных билетов

    start_at TIMESTAMP NOT NULL,             -- Время начало

    preview_url TEXT,                        -- Ссылка на картинку превью

    price DECIMAL(10,2),                     -- Цена билета

    duration INTEGER,                        -- Длительность в секундах

    area_id SERIAL REFERENCES areas, -- На какой площадке проходит
    type_id SERIAL REFERENCES type_events, -- На какой площадке проходит

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW()
);

COMMIT;
