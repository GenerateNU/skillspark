CREATE OR REPLACE FUNCTION notify_opensearch()
RETURNS TRIGGER AS $$
BEGIN
  BEGIN
    PERFORM net.http_post(
      url     := current_setting('app.settings.edge_function_url'),
      headers := jsonb_build_object(
        'Content-Type',  'application/json',
        'Authorization', current_setting('app.settings.edge_function_key')
      ),
      body    := jsonb_build_object(
        'type',       TG_OP,
        'record',     row_to_json(NEW),
        'old_record', row_to_json(OLD)
      )
    );
  EXCEPTION WHEN OTHERS THEN
    NULL;
  END;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER sync_event_to_opensearch
AFTER INSERT OR UPDATE OR DELETE ON event
FOR EACH ROW 
EXECUTE FUNCTION notify_opensearch();

