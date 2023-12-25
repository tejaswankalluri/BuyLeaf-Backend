package me.tejaswan.buyleaf.service;

import me.tejaswan.buyleaf.entity.ProductEntity;
import me.tejaswan.buyleaf.entity.ProductImage;
import me.tejaswan.buyleaf.exception.ResourceNotFoundException;
import org.springframework.web.multipart.MultipartFile;

import java.util.List;
import java.util.UUID;

public interface ProductService {
    ProductEntity createProduct(ProductEntity product);

    List<ProductEntity> getAllProducts();

    ProductEntity getProductById(UUID id) throws ResourceNotFoundException;

    List<ProductImage> addImagesToProduct(MultipartFile[] files, UUID id);
}
