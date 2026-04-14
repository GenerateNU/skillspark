-- NOTE - run each chunk of sql code depending on what lane is being deployed (uncomment/comment and then run supabase db reset)

-- dev
ALTER DATABASE postgres SET app.settings.edge_function_url = 'http://supabase_edge_runtime_skillspark:8081/opensearch-access';
ALTER DATABASE postgres SET app.settings.edge_function_key = 'Bearer sb_publishable_ACJWlzQHlZjBrEguHvfOxg_3BJgxAaH';

-- prod (NOTE** url may be incorrect)
-- ALTER DATABASE postgres SET app.settings.edge_function_url = 'https://jpxypbmcgjhirlaatufp.supabase.co/functions/v1/opensearch-access';
-- ALTER DATABASE postgres SET app.settings.edge_function_key = 'Bearer sb_publishable_wLzPTZVEXCDAz7HYJ9LQ5Q_AW65QGKO';