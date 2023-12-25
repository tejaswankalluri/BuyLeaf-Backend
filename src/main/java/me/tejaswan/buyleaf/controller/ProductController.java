package me.tejaswan.buyleaf.controller;

import io.swagger.v3.oas.annotations.OpenAPIDefinition;
import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.security.SecurityRequirement;
import io.swagger.v3.oas.annotations.tags.Tag;
import io.swagger.v3.oas.models.annotations.OpenAPI30;
import lombok.AllArgsConstructor;
import me.tejaswan.buyleaf.aws.s3.S3Service;
import me.tejaswan.buyleaf.entity.ProductEntity;
import me.tejaswan.buyleaf.exception.ResourceNotFoundException;
import me.tejaswan.buyleaf.service.ProductService;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.multipart.MultipartFile;

import java.io.IOException;
import java.util.List;
import java.util.UUID;

@RestController
@RequestMapping("/api/v1/product")
@SecurityRequirement(name = "jwtSchema")
@AllArgsConstructor
@Tag(name = "Products", description = "Product Apis")
public class ProductController {
    private final ProductService productService;

    @GetMapping("/all")
    @PreAuthorize("hasAuthority('ADMIN')")
    public ResponseEntity<List<ProductEntity>> getAllProducts() {
        return ResponseEntity.ok(productService.getAllProducts());
    }

    @PostMapping("/")
    @PreAuthorize("hasAuthority('ADMIN')")
    public ResponseEntity<ProductEntity> createProduct(@RequestBody ProductEntity product) {
        return ResponseEntity.status(HttpStatus.CREATED).body(productService.createProduct(product));
    }

    @GetMapping("/{id}")
    @PreAuthorize("hasAnyAuthority('ADMIN', 'USER')")
    public ResponseEntity<ProductEntity> getProductById(@PathVariable UUID id) throws ResourceNotFoundException {
        return ResponseEntity.ok(productService.getProductById(id));
    }

    @PostMapping("{id}/upload/images")
    public ResponseEntity<?> addProductImages(@RequestParam("file") MultipartFile[] files, @PathVariable("id") UUID id) {
        productService.addImagesToProduct(files,id);
        return ResponseEntity.ok("OK");
    }
}
