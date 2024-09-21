### Flow Summary

1. User uploads an image → API Gateway → Image Upload Service → Cloud Storage.
2. Image Processing Service retrieves image → Processes it → Sends emotion data to Notification Service via Message Broker.
3. Notification Service receives data → Sends email to user via Email Service.
