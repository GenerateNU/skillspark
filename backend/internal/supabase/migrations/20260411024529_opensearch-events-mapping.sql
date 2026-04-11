CREATE OR REPLACE FUNCTION notify_opensearch()
RETURNS TRIGGER AS $$
BEGIN
  BEGIN
    PERFORM net.http_post(
      url     := 'http://supabase_edge_runtime_skillspark:8081/opensearch-access',
      headers := jsonb_build_object(
        'Content-Type',  'application/json',
        'Authorization', 'Bearer sb_publishable_ACJWlzQHlZjBrEguHvfOxg_3BJgxAaH'
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

