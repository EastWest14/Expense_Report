--
-- PostgreSQL database dump
--

-- Dumped from database version 9.5.3
-- Dumped by pg_dump version 9.5.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: ExpenseReport; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON DATABASE "ExpenseReport" IS 'Works with the Expense-Report Go application.';


--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: app_user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE app_user (
    user_id integer NOT NULL,
    user_name character varying(50) NOT NULL
);


ALTER TABLE app_user OWNER TO postgres;

--
-- Name: app_user_user_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE app_user_user_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE app_user_user_id_seq OWNER TO postgres;

--
-- Name: app_user_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE app_user_user_id_seq OWNED BY app_user.user_id;


--
-- Name: expenditure; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE expenditure (
    exp_id integer NOT NULL,
    user_id integer NOT NULL,
    exp_amount real NOT NULL,
    exp_description character varying(100) NOT NULL
);


ALTER TABLE expenditure OWNER TO postgres;

--
-- Name: expenditure_exp_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE expenditure_exp_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE expenditure_exp_id_seq OWNER TO postgres;

--
-- Name: expenditure_exp_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE expenditure_exp_id_seq OWNED BY expenditure.exp_id;


--
-- Name: user_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY app_user ALTER COLUMN user_id SET DEFAULT nextval('app_user_user_id_seq'::regclass);


--
-- Name: exp_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY expenditure ALTER COLUMN exp_id SET DEFAULT nextval('expenditure_exp_id_seq'::regclass);


--
-- Name: app_user_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY app_user
    ADD CONSTRAINT app_user_pkey PRIMARY KEY (user_id);


--
-- Name: app_user_user_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY app_user
    ADD CONSTRAINT app_user_user_name_key UNIQUE (user_name);


--
-- Name: expenditure_exp_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY expenditure
    ADD CONSTRAINT expenditure_exp_id_key UNIQUE (exp_id);


--
-- Name: expenditure_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY expenditure
    ADD CONSTRAINT expenditure_pkey PRIMARY KEY (user_id, exp_id);


--
-- Name: expenditure_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY expenditure
    ADD CONSTRAINT expenditure_user_id_fkey FOREIGN KEY (user_id) REFERENCES app_user(user_id);


--
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

