resource "aws_opensearch_domain" "skillspark_opensearch" {
  domain_name    = var.opensearch_domain_name
  engine_version = var.opensearch_engine_version

  cluster_config {
    instance_type = var.opensearch_instance
  }

  advanced_security_options {
    enabled                        = true
    anonymous_auth_enabled         = true
    internal_user_database_enabled = true
    master_user_options {
      master_user_name     = var.master_user
      master_user_password = var.master_pass
    }
  }

  encrypt_at_rest {
    enabled = true
  }

  domain_endpoint_options {
    enforce_https       = true
    tls_security_policy = "Policy-Min-TLS-1-2-2019-07"
  }

  node_to_node_encryption {
    enabled = true
  }

  ebs_options {
    ebs_enabled = true
    volume_size = 10
  }
}

resource "opensearch_index" "event_index" {
  name               = "events"
  number_of_shards   = var.event_shards
  number_of_replicas = var.event_replicas
  mappings           = <<EOF
{
  "properties": {
    "id":             { "type": "keyword" },
    "title_en":       { "type": "text" },
    "title_th":       { "type": "text" },
    "description_en": { "type": "text" },
    "description_th": { "type": "text" },
    "category":       { "type": "keyword" },
    "age_range_min":  { "type": "integer" },
    "age_range_max":  { "type": "integer" }
  }
}
EOF
}