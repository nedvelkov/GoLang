**** - must be replace with bucket name
#### - must be replace with policy json name
^^^^ - must be replace with absolute path for static website

YouTube course : https://www.youtube.com/watch?v=3_sqr0G9zb0&ab_channel=MostlyCode
GitHub repo for course : https://github.com/mimo84/yt-tutorials


restrict bucket access command:
aws s3api --endpoint-url="http://localhost:4566" put-public-access-block --bucket **** --public-access-block-configuration "BlockPublicAcls=true,IgnorePublicAcls=true,BlockPublicPolicy=true,RestrictPublicBuckets=true"

allow public access for url to read from bucket command:
aws s3api --endpoint-url="http://localhost:4566" put-bucket-policy --bucket **** --policy file://####.json

upload static index.html to s3 bucket command:
aws s3 --endpoint-url="http://localhost:4566" sync ^^^^ "s3://****"



commands:
 
aws s3api --endpoint-url="http://localhost:4566" put-public-access-block --bucket my-bucket --public-access-block-configuration "BlockPublicAcls=true,IgnorePublicAcls=true,BlockPublicPolicy=true,RestrictPublicBuckets=true"

aws s3api --endpoint-url="http://localhost:4566" put-bucket-policy --bucket my-bucket --policy file://bucket-policy.json

aws s3 --endpoint-url="http://localhost:4566" website "s3://my-bucket" --index-document index.html --error-document index.html

aws s3 --endpoint-url="http://localhost:4566" sync D:/localstack/website "s3://my-bucket"

aws --endpoint-url=http://localhost:4566 dynamodb create-table --cli-input-json file://table-definition.json --region us-west-2