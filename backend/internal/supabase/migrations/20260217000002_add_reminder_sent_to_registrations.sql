alter table registration 
add column if not exists reminder_sent boolean not null default false;

comment on column registration.reminder_sent is 'Flag to indicate if the 24h reminder has been sent';
