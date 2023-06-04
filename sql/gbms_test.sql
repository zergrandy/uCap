--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2 (Debian 15.2-1.pgdg110+1)
-- Dumped by pg_dump version 15.2

-- Started on 2023-03-19 13:43:33 UTC

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
-- TOC entry 5 (class 2615 OID 16389)
-- Name: gbms; Type: SCHEMA; Schema: -; Owner: root
--

CREATE SCHEMA gbms;


ALTER SCHEMA gbms OWNER TO root;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 214 (class 1259 OID 16391)
-- Name: active; Type: TABLE; Schema: gbms; Owner: root
--

CREATE TABLE gbms.active (
    active smallint NOT NULL,
    description character varying(255)
);


ALTER TABLE gbms.active OWNER TO root;

--
-- TOC entry 216 (class 1259 OID 16397)
-- Name: group; Type: TABLE; Schema: gbms; Owner: root
--

CREATE TABLE gbms."group" (
    gid integer NOT NULL,
    name character varying(60),
    description text,
    updated timestamp with time zone DEFAULT now() NOT NULL,
    created timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE gbms."group" OWNER TO root;

--
-- TOC entry 215 (class 1259 OID 16396)
-- Name: group_gid_seq; Type: SEQUENCE; Schema: gbms; Owner: root
--

ALTER TABLE gbms."group" ALTER COLUMN gid ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME gbms.group_gid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 221 (class 1259 OID 16426)
-- Name: group_permission; Type: TABLE; Schema: gbms; Owner: root
--

CREATE TABLE gbms.group_permission (
    gid integer NOT NULL,
    pid integer NOT NULL
);


ALTER TABLE gbms.group_permission OWNER TO root;

--
-- TOC entry 223 (class 1259 OID 16434)
-- Name: grouping_group; Type: TABLE; Schema: gbms; Owner: root
--

CREATE TABLE gbms.grouping_group (
    gid integer NOT NULL,
    ggid integer NOT NULL
);


ALTER TABLE gbms.grouping_group OWNER TO root;

--
-- TOC entry 218 (class 1259 OID 16407)
-- Name: permission; Type: TABLE; Schema: gbms; Owner: root
--

CREATE TABLE gbms.permission (
    pid integer NOT NULL,
    title character varying(100) NOT NULL,
    slug character varying(100) NOT NULL,
    description text,
    active smallint NOT NULL,
    updated timestamp with time zone DEFAULT now() NOT NULL,
    created timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE gbms.permission OWNER TO root;

--
-- TOC entry 217 (class 1259 OID 16406)
-- Name: permission_pid_seq; Type: SEQUENCE; Schema: gbms; Owner: root
--

ALTER TABLE gbms.permission ALTER COLUMN pid ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME gbms.permission_pid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 225 (class 1259 OID 16551)
-- Name: uav_management; Type: TABLE; Schema: gbms; Owner: root
--

CREATE TABLE gbms.uav_management (
    uav_id integer NOT NULL,
    name character varying(20),
    brand character varying(20),
    model character varying(20),
    owner character varying(10) NOT NULL,
    location character varying(40) NOT NULL,
    active smallint DEFAULT 0,
    remark text,
    purchase date NOT NULL,
    maintenance timestamp with time zone,
    created timestamp with time zone DEFAULT now() NOT NULL,
    updated timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE gbms.uav_management OWNER TO root;

--
-- TOC entry 224 (class 1259 OID 16550)
-- Name: uav_managment_uav_id_seq; Type: SEQUENCE; Schema: gbms; Owner: root
--

CREATE SEQUENCE gbms.uav_managment_uav_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE gbms.uav_managment_uav_id_seq OWNER TO root;

--
-- TOC entry 3413 (class 0 OID 0)
-- Dependencies: 224
-- Name: uav_managment_uav_id_seq; Type: SEQUENCE OWNED BY; Schema: gbms; Owner: root
--

ALTER SEQUENCE gbms.uav_managment_uav_id_seq OWNED BY gbms.uav_management.uav_id;


--
-- TOC entry 220 (class 1259 OID 16417)
-- Name: user; Type: TABLE; Schema: gbms; Owner: root
--

CREATE TABLE gbms."user" (
    uid integer NOT NULL,
    login character varying(60) NOT NULL,
    username character varying(50) NOT NULL,
    password character varying(255) NOT NULL,
    email character varying(100) NOT NULL,
    active smallint,
    remark character varying(100),
    updated timestamp with time zone DEFAULT now() NOT NULL,
    created timestamp with time zone DEFAULT now() NOT NULL
);


ALTER TABLE gbms."user" OWNER TO root;

--
-- TOC entry 222 (class 1259 OID 16430)
-- Name: user_group; Type: TABLE; Schema: gbms; Owner: root
--

CREATE TABLE gbms.user_group (
    uid integer NOT NULL,
    gid integer NOT NULL
);


ALTER TABLE gbms.user_group OWNER TO root;

--
-- TOC entry 219 (class 1259 OID 16416)
-- Name: user_uid_seq; Type: SEQUENCE; Schema: gbms; Owner: root
--

ALTER TABLE gbms."user" ALTER COLUMN uid ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME gbms.user_uid_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 3213 (class 2604 OID 16554)
-- Name: uav_management uav_id; Type: DEFAULT; Schema: gbms; Owner: root
--

ALTER TABLE ONLY gbms.uav_management ALTER COLUMN uav_id SET DEFAULT nextval('gbms.uav_managment_uav_id_seq'::regclass);


--
-- TOC entry 3396 (class 0 OID 16391)
-- Dependencies: 214
-- Data for Name: active; Type: TABLE DATA; Schema: gbms; Owner: root
--

COPY gbms.active (active, description) FROM stdin;
0	disable
1	enable
\.


--
-- TOC entry 3398 (class 0 OID 16397)
-- Dependencies: 216
-- Data for Name: group; Type: TABLE DATA; Schema: gbms; Owner: root
--

COPY gbms."group" (gid, name, description, updated, created) FROM stdin;
1	群組A	群組第一層	2023-03-14 11:50:07.904223+00	2023-03-14 11:50:07.904223+00
2	群組B	群組第二層	2023-03-14 11:50:16.007396+00	2023-03-14 11:50:16.007396+00
\.


--
-- TOC entry 3403 (class 0 OID 16426)
-- Dependencies: 221
-- Data for Name: group_permission; Type: TABLE DATA; Schema: gbms; Owner: root
--

COPY gbms.group_permission (gid, pid) FROM stdin;
1	1
2	2
\.


--
-- TOC entry 3405 (class 0 OID 16434)
-- Dependencies: 223
-- Data for Name: grouping_group; Type: TABLE DATA; Schema: gbms; Owner: root
--

COPY gbms.grouping_group (gid, ggid) FROM stdin;
1	1
2	2
\.


--
-- TOC entry 3400 (class 0 OID 16407)
-- Dependencies: 218
-- Data for Name: permission; Type: TABLE DATA; Schema: gbms; Owner: root
--

COPY gbms.permission (pid, title, slug, description, active, updated, created) FROM stdin;
1	權限2	/group/permission/1	權限1	1	2023-03-14 11:46:44.617064+00	2023-03-14 11:46:44.617064+00
2	權限3	/group/permission/2	權限2	1	2023-03-14 11:46:54.734577+00	2023-03-14 11:46:54.734577+00
\.


--
-- TOC entry 3407 (class 0 OID 16551)
-- Dependencies: 225
-- Data for Name: uav_management; Type: TABLE DATA; Schema: gbms; Owner: root
--

COPY gbms.uav_management (uav_id, name, brand, model, owner, location, active, remark, purchase, maintenance, created, updated) FROM stdin;
1	無人機1	台灣希望創新	UAV-0057	程冠霖	新北市三重區中興北街	1		2023-03-15	\N	2023-03-15 14:29:56.934817+00	2023-03-15 14:50:24.945171+00
\.


--
-- TOC entry 3402 (class 0 OID 16417)
-- Dependencies: 220
-- Data for Name: user; Type: TABLE DATA; Schema: gbms; Owner: root
--

COPY gbms."user" (uid, login, username, password, email, active, remark, updated, created) FROM stdin;
1	T1@ttttttttttt1	T1@ttttttttttt12	$2a$10$gfaxpO6i8U.fJkGmFcsVC.NQceeThFqFEAOG.Da8BzElg2I3lKqmO	T1@ttttttttttt12	1	T1@ttttttttttt12	2023-03-16 04:19:27.677909+00	2023-03-16 03:52:57.641164+00
2	T2@tttttttttt1	T2@tttttttttt1	$2a$10$xCGw/f2SNyaRq3k5jaTkr.oEVN4yMBkKW.Q7cPqW6OFPGbTScU4xq	T2@tttttttttt1	1	T2@tttttttttt1	2023-03-16 05:57:13.082226+00	2023-03-16 05:57:13.082226+00
\.


--
-- TOC entry 3404 (class 0 OID 16430)
-- Dependencies: 222
-- Data for Name: user_group; Type: TABLE DATA; Schema: gbms; Owner: root
--

COPY gbms.user_group (uid, gid) FROM stdin;
1	1
2	2
\.


--
-- TOC entry 3414 (class 0 OID 0)
-- Dependencies: 215
-- Name: group_gid_seq; Type: SEQUENCE SET; Schema: gbms; Owner: root
--

SELECT pg_catalog.setval('gbms.group_gid_seq', 2, true);


--
-- TOC entry 3415 (class 0 OID 0)
-- Dependencies: 217
-- Name: permission_pid_seq; Type: SEQUENCE SET; Schema: gbms; Owner: root
--

SELECT pg_catalog.setval('gbms.permission_pid_seq', 2, true);


--
-- TOC entry 3416 (class 0 OID 0)
-- Dependencies: 224
-- Name: uav_managment_uav_id_seq; Type: SEQUENCE SET; Schema: gbms; Owner: root
--

SELECT pg_catalog.setval('gbms.uav_managment_uav_id_seq', 2, true);


--
-- TOC entry 3417 (class 0 OID 0)
-- Dependencies: 219
-- Name: user_uid_seq; Type: SEQUENCE SET; Schema: gbms; Owner: root
--

SELECT pg_catalog.setval('gbms.user_uid_seq', 2, true);


--
-- TOC entry 3219 (class 2606 OID 16395)
-- Name: active active_pkey; Type: CONSTRAINT; Schema: gbms; Owner: root
--

ALTER TABLE ONLY gbms.active
    ADD CONSTRAINT active_pkey PRIMARY KEY (active);


--
-- TOC entry 3221 (class 2606 OID 16826)
-- Name: group group_name_key; Type: CONSTRAINT; Schema: gbms; Owner: root
--

ALTER TABLE ONLY gbms."group"
    ADD CONSTRAINT group_name_key UNIQUE (name);


--
-- TOC entry 3236 (class 2606 OID 16612)
-- Name: group_permission group_permission_uniq_1; Type: CONSTRAINT; Schema: gbms; Owner: root
--

ALTER TABLE ONLY gbms.group_permission
    ADD CONSTRAINT group_permission_uniq_1 UNIQUE (gid, pid);


--
-- TOC entry 3223 (class 2606 OID 16571)
-- Name: group group_pkey; Type: CONSTRAINT; Schema: gbms; Owner: root
--

ALTER TABLE ONLY gbms."group"
    ADD CONSTRAINT group_pkey PRIMARY KEY (gid);


--
-- TOC entry 3217 (class 2606 OID 16646)
-- Name: grouping_group grouping_group_chk_1; Type: CHECK CONSTRAINT; Schema: gbms; Owner: root
--

ALTER TABLE gbms.grouping_group
    ADD CONSTRAINT grouping_group_chk_1 CHECK ((gid <> ggid)) NOT VALID;


--
-- TOC entry 3242 (class 2606 OID 16639)
-- Name: grouping_group grouping_group_uniq_1; Type: CONSTRAINT; Schema: gbms; Owner: root
--

ALTER TABLE ONLY gbms.grouping_group
    ADD CONSTRAINT grouping_group_uniq_1 UNIQUE (gid, ggid);


--
-- TOC entry 3225 (class 2606 OID 16654)
-- Name: permission permission_pkey; Type: CONSTRAINT; Schema: gbms; Owner: root
--

ALTER TABLE ONLY gbms.permission
    ADD CONSTRAINT permission_pkey PRIMARY KEY (pid);


--
-- TOC entry 3227 (class 2606 OID 16824)
-- Name: permission permission_uniq_1; Type: CONSTRAINT; Schema: gbms; Owner: root
--

ALTER TABLE ONLY gbms.permission
    ADD CONSTRAINT permission_uniq_1 UNIQUE (title, slug);


--
-- TOC entry 3245 (class 2606 OID 16561)
-- Name: uav_management uav_managment_pkey; Type: CONSTRAINT; Schema: gbms; Owner: root
--

ALTER TABLE ONLY gbms.uav_management
    ADD CONSTRAINT uav_managment_pkey PRIMARY KEY (uav_id);


--
-- TOC entry 3240 (class 2606 OID 16697)
-- Name: user_group user_group_uniq_1; Type: CONSTRAINT; Schema: gbms; Owner: root
--

ALTER TABLE ONLY gbms.user_group
    ADD CONSTRAINT user_group_uniq_1 UNIQUE (uid, gid);


--
-- TOC entry 3232 (class 2606 OID 16668)
-- Name: user user_pkey; Type: CONSTRAINT; Schema: gbms; Owner: root
--

ALTER TABLE ONLY gbms."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (uid);


--
-- TOC entry 3234 (class 2606 OID 16822)
-- Name: user user_uniq_1; Type: CONSTRAINT; Schema: gbms; Owner: root
--

ALTER TABLE ONLY gbms."user"
    ADD CONSTRAINT user_uniq_1 UNIQUE (login, email) INCLUDE (username);


--
-- TOC entry 3228 (class 1259 OID 16474)
-- Name: fki_f; Type: INDEX; Schema: gbms; Owner: root
--

CREATE INDEX fki_f ON gbms."user" USING btree (active);


--
-- TOC entry 3229 (class 1259 OID 16480)
-- Name: fki_grouping_group_fk_1; Type: INDEX; Schema: gbms; Owner: root
--

CREATE INDEX fki_grouping_group_fk_1 ON gbms."user" USING btree (active);


--
-- TOC entry 3237 (class 1259 OID 16613)
-- Name: index_on_group_permission; Type: INDEX; Schema: gbms; Owner: root
--

CREATE INDEX index_on_group_permission ON gbms.group_permission USING btree (gid, pid);


--
-- TOC entry 3243 (class 1259 OID 16640)
-- Name: index_on_grouping_group; Type: INDEX; Schema: gbms; Owner: root
--

CREATE INDEX index_on_grouping_group ON gbms.grouping_group USING btree (gid, ggid);


--
-- TOC entry 3230 (class 1259 OID 16468)
-- Name: index_on_user; Type: INDEX; Schema: gbms; Owner: root
--

CREATE INDEX index_on_user ON gbms."user" USING btree (active);


--
-- TOC entry 3238 (class 1259 OID 16698)
-- Name: index_on_user_group; Type: INDEX; Schema: gbms; Owner: root
--

CREATE INDEX index_on_user_group ON gbms.user_group USING btree (uid, gid);


--
-- TOC entry 3248 (class 2606 OID 16601)
-- Name: group_permission group_permission_fk_1; Type: FK CONSTRAINT; Schema: gbms; Owner: root
--

ALTER TABLE ONLY gbms.group_permission
    ADD CONSTRAINT group_permission_fk_1 FOREIGN KEY (gid) REFERENCES gbms."group"(gid) ON DELETE CASCADE NOT VALID;


--
-- TOC entry 3249 (class 2606 OID 16655)
-- Name: group_permission group_permission_fk_2; Type: FK CONSTRAINT; Schema: gbms; Owner: root
--

ALTER TABLE ONLY gbms.group_permission
    ADD CONSTRAINT group_permission_fk_2 FOREIGN KEY (pid) REFERENCES gbms.permission(pid) ON DELETE CASCADE NOT VALID;


--
-- TOC entry 3252 (class 2606 OID 16628)
-- Name: grouping_group grouping_group_fk_1; Type: FK CONSTRAINT; Schema: gbms; Owner: root
--

ALTER TABLE ONLY gbms.grouping_group
    ADD CONSTRAINT grouping_group_fk_1 FOREIGN KEY (gid) REFERENCES gbms."group"(gid) ON DELETE CASCADE NOT VALID;


--
-- TOC entry 3253 (class 2606 OID 16641)
-- Name: grouping_group grouping_group_ggid_fkey; Type: FK CONSTRAINT; Schema: gbms; Owner: root
--

ALTER TABLE ONLY gbms.grouping_group
    ADD CONSTRAINT grouping_group_ggid_fkey FOREIGN KEY (ggid) REFERENCES gbms."group"(gid) ON DELETE CASCADE NOT VALID;


--
-- TOC entry 3246 (class 2606 OID 16517)
-- Name: permission permission_fk_1; Type: FK CONSTRAINT; Schema: gbms; Owner: root
--

ALTER TABLE ONLY gbms.permission
    ADD CONSTRAINT permission_fk_1 FOREIGN KEY (active) REFERENCES gbms.active(active) ON UPDATE RESTRICT ON DELETE RESTRICT NOT VALID;


--
-- TOC entry 3247 (class 2606 OID 16475)
-- Name: user user_fk_1; Type: FK CONSTRAINT; Schema: gbms; Owner: root
--

ALTER TABLE ONLY gbms."user"
    ADD CONSTRAINT user_fk_1 FOREIGN KEY (active) REFERENCES gbms.active(active) ON DELETE CASCADE NOT VALID;


--
-- TOC entry 3250 (class 2606 OID 16686)
-- Name: user_group user_group_fk_1; Type: FK CONSTRAINT; Schema: gbms; Owner: root
--

ALTER TABLE ONLY gbms.user_group
    ADD CONSTRAINT user_group_fk_1 FOREIGN KEY (uid) REFERENCES gbms."user"(uid) ON DELETE CASCADE NOT VALID;


--
-- TOC entry 3251 (class 2606 OID 16699)
-- Name: user_group user_group_fk_2; Type: FK CONSTRAINT; Schema: gbms; Owner: root
--

ALTER TABLE ONLY gbms.user_group
    ADD CONSTRAINT user_group_fk_2 FOREIGN KEY (gid) REFERENCES gbms."group"(gid) ON DELETE CASCADE NOT VALID;


-- Completed on 2023-03-19 13:43:33 UTC

--
-- PostgreSQL database dump complete
--

