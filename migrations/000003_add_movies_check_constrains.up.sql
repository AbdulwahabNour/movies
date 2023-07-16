Alter table movies Add constraint movies_runtime_check check (runtime > 0);
Alter table movies Add constraint movies_year_check check (year Between 1888 and date_part('year', now()));
Alter table movies Add constraint genres_length_check check (array_length(genres, 1)  Between 1 and 5);

