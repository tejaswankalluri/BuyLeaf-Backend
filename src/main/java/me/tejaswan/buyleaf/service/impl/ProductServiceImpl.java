package me.tejaswan.buyleaf.service.impl;

import lombok.AllArgsConstructor;
import me.tejaswan.buyleaf.aws.s3.S3Service;
import me.tejaswan.buyleaf.entity.ProductEntity;
import me.tejaswan.buyleaf.entity.ProductImage;
import me.tejaswan.buyleaf.exception.ResourceNotFoundException;
import me.tejaswan.buyleaf.repository.ProductImageRepository;
import me.tejaswan.buyleaf.repository.ProductRepository;
import me.tejaswan.buyleaf.service.ProductService;
import org.springframework.stereotype.Service;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.util.*;

@Service
@AllArgsConstructor
public class ProductServiceImpl implements ProductService {
    private final ProductRepository productRepository;
    private final ProductImageRepository productImageRepository;
    private final S3Service s3Service;

    public ProductEntity createProduct(ProductEntity product) {
        return productRepository.save(product);
    }

    public List<ProductEntity> getAllProducts() {
        return productRepository.findAll();
    }

    public ProductEntity getProductById(UUID id) throws ResourceNotFoundException {
        Optional<ProductEntity> user = productRepository.findById(id);
        if (user.isEmpty()) {
            throw new ResourceNotFoundException("Product Not found");
        }
        return user.get();
    }

    public List<ProductImage> addImagesToProduct(MultipartFile[] files, UUID id) {
        List<ProductImage> images = new ArrayList<>();
        Arrays.stream(files).forEach(file -> {
            try {
                byte[] byteFile = file.getBytes();
                String fileName = System.currentTimeMillis() + "_" + file.getOriginalFilename();
                String URL = s3Service.putObjectAndGetURL("buyleaf", fileName, byteFile);

                ProductImage image = new ProductImage();
                ProductEntity product = new ProductEntity();
                product.setId(id);
                image.setUrl(URL);
                image.setProduct(product);
                images.add(image);

            } catch (IOException e) {
                throw new RuntimeException(e);
            }
        });
        return productImageRepository.saveAll(images);
    }
}
