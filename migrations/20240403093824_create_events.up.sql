/* Таблица мероприятий */
BEGIN;

create table events (
    id SERIAL PRIMARY KEY,                  -- Уникальный ID

    title VARCHAR(255),                     -- Название

    description TEXT,                       -- Описание

    publish BOOLEAN NOT NULL DEFAULT false,  -- Опубликовано или нет событие

    ticket_count INTEGER DEFAULT 0,          -- Общее количество доступных билетов

    start_at TIMESTAMP,             -- Время начало

    preview_url TEXT,                        -- Ссылка на картинку превью

    price DECIMAL(10,2),                     -- Цена билета

    duration INTEGER,                        -- Длительность в секундах

    area_id int REFERENCES areas(id), -- На какой площадке проходит
    type_id int REFERENCES type_events(id), -- На какой площадке проходит

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

COMMIT;
