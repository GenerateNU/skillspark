create table if not exists notification (
  id uuid primary key default gen_random_uuid(),
  registration_id uuid references registration(id) on delete cascade not null,
  type text not null check (type in ('sms', 'email')),
  payload jsonb not null default '{}'::jsonb,
  created_at timestamptz not null default now()
);

-- Add comment
comment on table notification is 'Queue for sending notifications (SMS/Email) via Supabase Edge Functions';
