resource "aws_glue_catalog_database" "object_db" {
  name = local.service_name
}