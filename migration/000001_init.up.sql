CREATE TABLE IF NOT EXISTS public.status
(
    status_id uuid NOT NULL DEFAULT gen_random_uuid(),
    name character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT status_pkey PRIMARY KEY (status_id)
);
CREATE TABLE IF NOT EXISTS public."user"
(
    user_id uuid NOT NULL DEFAULT gen_random_uuid(),
    name character varying COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT user_pk PRIMARY KEY (user_id)
);
CREATE TABLE IF NOT EXISTS public.task
(
    task_id uuid NOT NULL DEFAULT gen_random_uuid(),
    status_id uuid NOT NULL,
    title character varying COLLATE pg_catalog."default" NOT NULL DEFAULT ''::character varying,
    description character varying COLLATE pg_catalog."default" NOT NULL DEFAULT ''::character varying,
    CONSTRAINT task_pk PRIMARY KEY (task_id),
    CONSTRAINT task_status_id_fk FOREIGN KEY (status_id)
        REFERENCES public.status (status_id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE RESTRICT
        NOT VALID
);
CREATE TABLE IF NOT EXISTS public.message
(
    message_id uuid NOT NULL DEFAULT gen_random_uuid(),
    task_id uuid NOT NULL,
    user_id uuid DEFAULT '00000000-0000-0000-0000-000000000000'::uuid,
    text character varying COLLATE pg_catalog."default" NOT NULL DEFAULT ''::character varying,
    create_at timestamp with time zone NOT NULL DEFAULT now(),
    update_at timestamp with time zone,
    CONSTRAINT message_pk PRIMARY KEY (message_id),
    CONSTRAINT message_task_id_fk FOREIGN KEY (task_id)
        REFERENCES public.task (task_id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT message_user_id_fk FOREIGN KEY (user_id)
        REFERENCES public."user" (user_id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE SET NULL
        NOT VALID
);
CREATE TABLE IF NOT EXISTS public.task_user
(
    task_user_id uuid NOT NULL DEFAULT gen_random_uuid(),
    task_id uuid NOT NULL,
    user_id uuid NOT NULL,
    CONSTRAINT task_user_pk PRIMARY KEY (task_user_id),
    CONSTRAINT task_user_task_id_user_id_unique UNIQUE (task_id, user_id),
    CONSTRAINT task_user_task_id_fk FOREIGN KEY (task_id)
        REFERENCES public.task (task_id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID,
    CONSTRAINT task_user_user_id_fk FOREIGN KEY (user_id)
        REFERENCES public."user" (user_id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID
);