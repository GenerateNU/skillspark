alter table guardian
add column if not exists line_account_id text;

comment on column guardian.line_account_id is 'Line Account ID for receiving notifications';


create extension if not exists pg_net with schema extensions;

create or replace function notify_on_insert()
returns trigger as $$
begin
  perform net.http_post(
    url := 'http://host.docker.internal:54321/functions/v1/send-notification',
    headers := '{"Content-Type": "application/json"}'::jsonb,
    body := jsonb_build_object(
      'type',       TG_OP,
      'table',      TG_TABLE_NAME,
      'schema',     TG_TABLE_SCHEMA,
      'record',     row_to_json(NEW),
      'old_record', row_to_json(OLD)
    )
  );
  return NEW;
end;
$$ language plpgsql security definer;

create trigger on_notification_insert
  after insert on public.notification
  for each row execute function notify_on_insert();