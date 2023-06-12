resource "aws_dynamodb_table_item" "item1"{
    depends_on= [
        aws_dynamodb_table.users_table
    ]
    table_name      =   aws_dynamodb_table.users_table.name
    hash_key        =   aws_dynamodb_table.users_table.hash_key

    item    =<<ITEM
    {
        "Email": { "S": "n.velkov@gmail.com" },
        "FirstName": { "S": "Ned" },
        "LastName": { "S": "Velkov" }
    }
    ITEM
}

resource "aws_dynamodb_table_item" "item2"{
    depends_on= [
        aws_dynamodb_table.users_table
    ]
    table_name      =   aws_dynamodb_table.users_table.name
    hash_key        =   aws_dynamodb_table.users_table.hash_key

    item    =<<ITEM
    {
          "Email": { "S": "test@test.com" },
          "FirstName": { "S": "Test" },
          "LastName": { "S": "Test" }
    }
    ITEM
}