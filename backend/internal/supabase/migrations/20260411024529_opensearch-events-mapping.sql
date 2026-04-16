CREATE OR REPLACE FUNCTION notify_opensearch()
RETURNS TRIGGER AS $$
DECLARE
  _url  text;
  _key  text;
  _body jsonb;
BEGIN
  BEGIN
    SELECT decrypted_secret INTO _url FROM vault.decrypted_secrets WHERE name = 'edge_function_url' LIMIT 1;
    SELECT decrypted_secret INTO _key FROM vault.decrypted_secrets WHERE name = 'edge_function_key' LIMIT 1;
  EXCEPTION WHEN undefined_table THEN
    -- vault schema not available (e.g. test environment) — skip sync
    IF TG_OP = 'DELETE' THEN RETURN OLD; ELSE RETURN NEW; END IF;
  END;

  IF _url IS NULL OR _key IS NULL THEN
    IF TG_OP = 'DELETE' THEN RETURN OLD; ELSE RETURN NEW; END IF;
  END IF;

  IF TG_OP = 'DELETE' THEN
    _body := jsonb_build_object(
      'type',       TG_OP,
      'record',     row_to_json(OLD),
      'old_record', row_to_json(OLD)
    );
    PERFORM net.http_post(
      url     := _url,
      headers := jsonb_build_object('Content-Type', 'application/json', 'Authorization', _key),
      body    := _body
    );
    RETURN OLD;
  ELSE
    _body := jsonb_build_object(
      'type',       TG_OP,
      'record',     row_to_json(NEW),
      'old_record', row_to_json(OLD)
    );
    BEGIN
      PERFORM net.http_post(
        url     := _url,
        headers := jsonb_build_object('Content-Type', 'application/json', 'Authorization', _key),
        body    := _body
      );
    EXCEPTION WHEN OTHERS THEN
      RAISE WARNING 'notify_opensearch failed for operation %, SQLSTATE %, error: %', TG_OP, SQLSTATE, SQLERRM;
    END;
    RETURN NEW;
  END IF;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER sync_event_to_opensearch
AFTER INSERT OR UPDATE OR DELETE ON event
FOR EACH ROW 
EXECUTE FUNCTION notify_opensearch();

