--
-- PostgreSQL database dump
--

-- Dumped from database version 10.12 (Ubuntu 10.12-0ubuntu0.18.04.1)
-- Dumped by pg_dump version 10.12 (Ubuntu 10.12-0ubuntu0.18.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: achievements; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.achievements (
    id integer NOT NULL,
    name text NOT NULL,
    description text,
    link text
);


ALTER TABLE public.achievements OWNER TO postgres;

--
-- Name: achievements_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.achievements_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.achievements_id_seq OWNER TO postgres;

--
-- Name: achievements_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.achievements_id_seq OWNED BY public.achievements.id;


--
-- Name: sections; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sections (
    username text,
    name text NOT NULL,
    id integer NOT NULL
);


ALTER TABLE public.sections OWNER TO postgres;

--
-- Name: sections_achievements; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sections_achievements (
    section_id integer,
    achievement_id integer
);


ALTER TABLE public.sections_achievements OWNER TO postgres;

--
-- Name: sections_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sections_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sections_id_seq OWNER TO postgres;

--
-- Name: sections_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sections_id_seq OWNED BY public.sections.id;


--
-- Name: uniquecode; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.uniquecode (
    code text NOT NULL,
    emailid text NOT NULL,
    passwordset boolean
);


ALTER TABLE public.uniquecode OWNER TO postgres;

--
-- Name: userdetails; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.userdetails (
    username text,
    firstname text,
    lastname text,
    bio text,
    designation_stand text,
    contactnumber text,
    emailid text,
    currentlocation text
);


ALTER TABLE public.userdetails OWNER TO postgres;

--
-- Name: usernames; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.usernames (
    username text NOT NULL,
    user_id integer
);


ALTER TABLE public.usernames OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id integer NOT NULL,
    fname text NOT NULL,
    lname text NOT NULL,
    email character varying(50),
    password text
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: achievements id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.achievements ALTER COLUMN id SET DEFAULT nextval('public.achievements_id_seq'::regclass);


--
-- Name: sections id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sections ALTER COLUMN id SET DEFAULT nextval('public.sections_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: achievements; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.achievements (id, name, description, link) FROM stdin;
1	linkedin	positive site	www.linkedin.com
2	hackerR	positive hsckrt sfgsukg dgfksdbfksdha dhsgfas	www.hacker.com
3	linuxworld	vial daga site	www.linusworld.com
4	testing portal	quant site	www.testter.com
5	testproject		www.with no bio.com
6	ninthB	the first letter	ningth.com
7	twelth	the first spelling	twelthscience.com
8	sjs	10th	www.google.com
9	sjs	12th	sjsaburoad.com
10	git	graduation	gitjaipur.com
11	nintha	the first letter fghj kvbsd vH jhch sc X sdhc  scuhdasub adudsahsa udsaduhsaud sadusa	ningth.com
12	twelth3	the first spelling call the name by hellow and how are you	twelthscience.com
13	sjs	10th	www.google.com
14	sjs	12th	sjsaburoad.com
15	git	graduation	gitjaipur.com
\.


--
-- Data for Name: sections; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sections (username, name, id) FROM stdin;
test	certificates	1
test	expirences	2
test	prjects	3
test	tester	4
test	class	5
test	educaion	6
test	class2	7
test	educaion syatem	8
\.


--
-- Data for Name: sections_achievements; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.sections_achievements (section_id, achievement_id) FROM stdin;
1	1
1	2
2	3
3	4
3	5
5	6
5	7
6	8
6	9
6	10
7	11
7	12
8	13
8	14
8	15
\.


--
-- Data for Name: uniquecode; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.uniquecode (code, emailid, passwordset) FROM stdin;
\.


--
-- Data for Name: userdetails; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.userdetails (username, firstname, lastname, bio, designation_stand, contactnumber, emailid, currentlocation) FROM stdin;
test	testFname	testLname	its a test account bio where it is represention what to be done and what not , or we can say a sort discription about self	test is Software Engineer	0999999999	test@test.com	india
\.


--
-- Data for Name: usernames; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.usernames (username, user_id) FROM stdin;
test	1
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, fname, lname, email, password) FROM stdin;
1	test1	lastetest	test@test.com	$2a$10$FHmUKhwWOpfS8hxwr.y05e79hQCBHlPldaaAwBiZshR7bNrJPfayi
\.


--
-- Name: achievements_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.achievements_id_seq', 15, true);


--
-- Name: sections_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.sections_id_seq', 8, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 1, true);


--
-- Name: achievements achievements_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.achievements
    ADD CONSTRAINT achievements_pkey PRIMARY KEY (id);


--
-- Name: sections sections_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sections
    ADD CONSTRAINT sections_pkey PRIMARY KEY (id);


--
-- Name: usernames usernames_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.usernames
    ADD CONSTRAINT usernames_pkey PRIMARY KEY (username);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: sections_achievements sections_achievements_achievement_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sections_achievements
    ADD CONSTRAINT sections_achievements_achievement_id_fkey FOREIGN KEY (achievement_id) REFERENCES public.achievements(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: sections_achievements sections_achievements_section_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sections_achievements
    ADD CONSTRAINT sections_achievements_section_id_fkey FOREIGN KEY (section_id) REFERENCES public.sections(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: sections sections_username_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sections
    ADD CONSTRAINT sections_username_fkey FOREIGN KEY (username) REFERENCES public.usernames(username) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: userdetails userdetails_username_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.userdetails
    ADD CONSTRAINT userdetails_username_fkey FOREIGN KEY (username) REFERENCES public.usernames(username) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: usernames usernames_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.usernames
    ADD CONSTRAINT usernames_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

