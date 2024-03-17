/* Таблица мероприятий */
BEGIN;

create table events (
    id SERIAL PRIMARY KEY,                  -- Уникальный ID

    title VARCHAR(255),                     -- Название

    publish BOOLEAN NOT NULL DEFAULT 0,      -- Опубликовано или нет событие

    ticket_count INTEGER DEFAULT 0,          -- Общее количество доступных билетов

    start_at TIMESTAMP NOT NULL,             -- Время начало 

    duration INTEGER,                        -- Длительность в секундах

    area_id SERIAL REFERENCES areas(area_id), -- На какой площадке проходит
    type_id SERIAL REFERENCES types(type_id), -- На какой площадке проходит

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW()
);

COMMIT;
