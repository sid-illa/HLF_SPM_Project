-- Database: capstone_blockchain

DROP DATABASE IF EXISTS capstone_blockchain;

CREATE DATABASE capstone_blockchain
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.UTF-8'
    LC_CTYPE = 'en_US.UTF-8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;


-- Table: public.certificatesinfo

DROP TABLE IF EXISTS public.certificatesinfo;

CREATE TABLE IF NOT EXISTS public.certificatesinfo
(
    id integer NOT NULL DEFAULT nextval('certificatesinfo_id_seq'::regclass),
    certificate character varying COLLATE pg_catalog."default" NOT NULL,
    privatekey character varying COLLATE pg_catalog."default",
    username character varying COLLATE pg_catalog."default",
    userid integer,
    CONSTRAINT certificatesinfo_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.certificatesinfo
    OWNER to postgres;


Table: public.certificatesinfo

-- DROP TABLE IF EXISTS public.certificatesinfo;

CREATE TABLE IF NOT EXISTS public.certificatesinfo
(
    id integer NOT NULL DEFAULT nextval('certificatesinfo_id_seq'::regclass),
    certificate character varying COLLATE pg_catalog."default" NOT NULL,
    privatekey character varying COLLATE pg_catalog."default",
    username character varying COLLATE pg_catalog."default",
    userid integer,
    CONSTRAINT certificatesinfo_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.certificatesinfo
    OWNER to postgres;