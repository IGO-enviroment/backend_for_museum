/* Таблица связи между мероприятиями и тегами */
BEGIN;

create table event_tags (
    id SERIAL PRIMARY KEY,      -- Уникальный ID

    event_id SERIAL REFERENCES events,    -- К какому мероприятию относится
    tag_id SERIAL REFERENCES tags,          -- С каким тегом есть связь

    -- Таймстамп
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

COMMIT;
