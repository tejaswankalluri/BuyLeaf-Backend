package me.tejaswan.buyleaf.service.impl;

import lombok.AllArgsConstructor;
import me.tejaswan.buyleaf.entity.ProductEntity;
import me.tejaswan.buyleaf.exception.ResourceNotFoundException;
import me.tejaswan.buyleaf.repository.ProductRepository;
import me.tejaswan.buyleaf.service.ProductService;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

@Service
@AllArgsConstructor
public class ProductServiceImpl implements ProductService {
    private final ProductRepository productRepository;
    public ProductEntity createProduct(ProductEntity product) {
        return productRepository.save(product);
    }

    public List<ProductEntity> getAllProducts() {
        return productRepository.findAll();
    }
    public ProductEntity getProductById(UUID id) throws ResourceNotFoundException {
        Optional<ProductEntity> user= productRepository.findById(id);
        if(user.isEmpty()) {
            throw new ResourceNotFoundException("Product Not found");
        }
        return user.get();
    }
}
