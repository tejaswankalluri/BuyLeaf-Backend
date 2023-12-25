package me.tejaswan.buyleaf.aws.s3;

import lombok.AllArgsConstructor;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;
import software.amazon.awssdk.core.sync.RequestBody;
import software.amazon.awssdk.services.s3.S3Client;
import software.amazon.awssdk.services.s3.model.PutObjectRequest;

@Service
@AllArgsConstructor
public class S3Service {
    private final S3Client s3;

   public String putObjectAndGetURL(String bucketName, String key, byte[] file){
       PutObjectRequest objectRequest = PutObjectRequest.builder()
               .bucket(bucketName)
               .key(key)
               .build();
       s3.putObject(objectRequest, RequestBody.fromBytes(file));

       String objectUrl = "https://" + bucketName + ".s3.amazonaws.com/" + key;
       return objectUrl;
   }
}
