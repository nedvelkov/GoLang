resource "aws_dynamodb_table" "users_table" {
  name           = var.table_name
  billing_mode   = "PROVISIONED"
  read_capacity  = 5
  write_capacity = 5

  hash_key = "Id"
  attribute {
    name = "Id"
    type = "S"
  }
}
