$bucket_name = terraform output s3_bucket_name

awslocal s3 rm s3://$bucket_name --recursive

terraform destroy -auto-approve

docker rm -f website
