resource "aws_dynamodb_table" "users_table"{
    name                =   "users"
    billing_mode        =   "PROVISIONED"
    read_capacity       =   5
    write_capacity      =   5

    hash_key            =   "Email"
    attribute {
        name    =   "Email"
        type    =   "S"
    }
}