# Resources

In folder **aws** are located event pattern and matching rule for this event, tested in AWS eventbridge sandbox.

In folder **localstack** is located event pattern for testing eventbridge in localstack. Implementation on AWS EventBridge in localstack use different parameters. For testing EventBridge in localstack, first install [AWS Command Line Interface | Docs (localstack.cloud)](https://docs.localstack.cloud/user-guide/integrations/aws-cli/) and execute command :

> awslocal events put-events --entries file://file-name.json localstack-event.json

- use event pattern in **localstack** folder
