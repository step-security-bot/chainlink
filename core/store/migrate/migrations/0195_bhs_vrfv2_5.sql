-- +goose Up
ALTER TABLE blockhash_store_specs 
    ADD COLUMN "coordinator_v2_5_address" bytea
    CHECK (octet_length(coordinator_v2_5_address) = 20);

UPDATE blockhash_store_specs SET coordinator_v2_5_address = coordinator_v2_plus_address
WHERE coordinator_v2_plus_address IS NOT NULL;

ALTER TABLE blockhash_store_specs DROP COLUMN "coordinator_v2_plus_address";

ALTER TABLE block_header_feeder_specs
    ADD COLUMN "coordinator_v2_5_address" bytea
    CHECK (octet_length(coordinator_v2_5_address) = 20);

UPDATE block_header_feeder_specs SET coordinator_v2_5_address = coordinator_v2_plus_address
WHERE coordinator_v2_plus_address IS NOT NULL;

ALTER TABLE block_header_feeder_specs DROP COLUMN "coordinator_v2_plus_address";

-- +goose Down
ALTER TABLE blockhash_store_specs 
    ADD COLUMN "coordinator_v2_plus_address" bytea
    CHECK (octet_length(coordinator_v2_plus_address) = 20);

UPDATE blockhash_store_specs SET coordinator_v2_plus_address = coordinator_v2_5_address
WHERE coordinator_v2_5_address IS NOT NULL;

ALTER TABLE blockhash_store_specs DROP COLUMN "coordinator_v2_5_address";

ALTER TABLE block_header_feeder_specs
    ADD COLUMN "coordinator_v2_plus_address" bytea
    CHECK (octet_length(coordinator_v2_plus_address) = 20);

UPDATE block_header_feeder_specs SET coordinator_v2_plus_address = coordinator_v2_5_address
WHERE coordinator_v2_5_address IS NOT NULL;

ALTER TABLE block_header_feeder_specs DROP COLUMN "coordinator_v2_5_address";