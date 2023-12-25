package me.tejaswan.buyleaf.repository;

import me.tejaswan.buyleaf.entity.ProductImage;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.UUID;

public interface ProductImageRepository extends JpaRepository<ProductImage, UUID> {
}
