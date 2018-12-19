CREATE TABLE public.wm_test_ss
(
    force text,
    month_year text,
    time timestamp(6)
    without time zone,
    age_range text,
    ethnicity text,
    search_happened boolean,
    outcome text,
    gender text,
    outcome_linked_to_object boolean,
    object_of_search text
)
    WITH
    (
    OIDS = FALSE
);

    ALTER TABLE public.wm_test_ss
    OWNER to postgres;