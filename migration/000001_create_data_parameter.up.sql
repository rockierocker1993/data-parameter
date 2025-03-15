CREATE SEQUENCE system_values_id_seq
    AS BIGINT
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE SEQUENCE lookup_values_id_seq
    AS BIGINT
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

CREATE TABLE system_values (
    id BIGINT PRIMARY KEY DEFAULT nextval('system_values_id_seq'),
    module TEXT NOT NULL,
    key TEXT NOT NULL,
    value TEXT NOT NULL,
    is_encrypt BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

CREATE TABLE lookup_values (
    id BIGINT PRIMARY KEY DEFAULT nextval('lookup_values_id_seq'),
    key TEXT NOT NULL,
    value TEXT NOT NULL,
    text_id TEXT NOT NULL,
    text_en TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- Menambahkan index untuk kolom deleted_at guna mendukung soft delete
CREATE INDEX idx_system_values_deleted_at ON system_values (deleted_at);
CREATE INDEX idx_lookup_values_deleted_at ON lookup_values (deleted_at);