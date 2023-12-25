package me.tejaswan.buyleaf.service;

import me.tejaswan.buyleaf.entity.ProductEntity;
import me.tejaswan.buyleaf.exception.ResourceNotFoundException;

import java.util.List;
import java.util.UUID;

public interface ProductService {
    ProductEntity createProduct(ProductEntity product);
    List<ProductEntity> getAllProducts();
    ProductEntity getProductById(UUID id) throws ResourceNotFoundException;

}
