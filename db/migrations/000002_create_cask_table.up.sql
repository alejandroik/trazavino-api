CREATE TABLE IF NOT EXISTS cask
(
    id         uuid PRIMARY KEY,
    created_at timestamp                   NOT NULL,
    updated_at timestamp,
    deleted_at timestamp,
    winery_id  uuid references winery (id) NOT NULL,
    name       varchar(20)                 NOT NULL,
    c_type     varchar(20)                 NOT NULL,
    is_empty   bool                        NOT NULL
);