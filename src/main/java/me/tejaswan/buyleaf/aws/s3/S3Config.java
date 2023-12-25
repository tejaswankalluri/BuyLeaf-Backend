package me.tejaswan.buyleaf.aws.s3;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import software.amazon.awssdk.auth.credentials.AwsBasicCredentials;
import software.amazon.awssdk.auth.credentials.AwsCredentialsProvider;
import software.amazon.awssdk.regions.Region;
import software.amazon.awssdk.services.s3.S3Client;


@Configuration
public class S3Config {
    @Value("${aws.region}")
    private String awsRegion;
    @Value("${aws.credentials.access-key}")
    private String accessKey;
    @Value("${aws.credentials.secret-key}")
    private String secretKey;

    @Bean
    public S3Client s3Client() {
        AwsBasicCredentials awsCredentials = AwsBasicCredentials.create(accessKey, secretKey);
        S3Client client = S3Client.builder()
                .region(Region.of(awsRegion))
                .credentialsProvider(() -> awsCredentials)
                .build();
        return client;
    }
}
