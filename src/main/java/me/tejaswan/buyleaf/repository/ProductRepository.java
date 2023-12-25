package me.tejaswan.buyleaf.repository;

import me.tejaswan.buyleaf.entity.ProductEntity;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.UUID;

public interface ProductRepository extends JpaRepository<ProductEntity, UUID> {
}
