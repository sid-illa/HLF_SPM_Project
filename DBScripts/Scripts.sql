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


-- SEQUENCE: public.user_id_seq

-- DROP SEQUENCE IF EXISTS public.user_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.user_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1
    OWNED BY "user".id;

ALTER SEQUENCE public.user_id_seq
    OWNER TO postgres;

-- Table: public.user

 DROP TABLE IF EXISTS public."user";

CREATE TABLE IF NOT EXISTS public."user"
(
    id integer NOT NULL DEFAULT nextval('user_id_seq'::regclass),
    username character varying COLLATE pg_catalog."default" NOT NULL,
    password character varying COLLATE pg_catalog."default" NOT NULL,
    organization character varying COLLATE pg_catalog."default",
    CONSTRAINT user_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."user"
    OWNER to postgres;


-- SEQUENCE: public.certificatesinfo_id_seq

-- DROP SEQUENCE IF EXISTS public.certificatesinfo_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.certificatesinfo_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 2147483647
    CACHE 1
    OWNED BY certificatesinfo.id;

ALTER SEQUENCE public.certificatesinfo_id_seq
    OWNER TO postgres;

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

