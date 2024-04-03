-- Откатывающая миграция для изменения имени таблицы, переименования колонок и удаления новой колонки
BEGIN TRANSACTION;

-- Удаляем добавленную колонку updated_at
ALTER TABLE event_types DROP COLUMN updated_at;

-- Возвращаем старое имя таблицы
ALTER TABLE event_types RENAME TO type_events;

-- Возвращаем старые имена колонок
ALTER TABLE type_events RENAME COLUMN name TO named;
ALTER TABLE type_events RENAME COLUMN description TO property;
ALTER TABLE type_events RENAME COLUMN is_visible TO publish;

COMMIT;