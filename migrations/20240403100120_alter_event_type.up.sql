-- Миграция для изменения имени таблицы, переименования колонок и добавления новой колонки
BEGIN;

-- Изменяем имя таблицы с type_events на event_types
ALTER TABLE type_events RENAME TO event_types;

-- Переименовываем колонку Named на name
ALTER TABLE event_types RENAME COLUMN "named" TO name;

-- Переименовываем колонку Property на description
ALTER TABLE event_types RENAME COLUMN "property" TO description;

-- Переименовываем колонку Publish на is_visible
ALTER TABLE event_types RENAME COLUMN "publish" TO is_visible;

-- Добавляем новую колонку updated_at с типом timestamp
ALTER TABLE event_types ADD COLUMN updated_at TIMESTAMP;

COMMIT;