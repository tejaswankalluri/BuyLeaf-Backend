package me.tejaswan.buyleaf.entity;

import jakarta.persistence.*;
import lombok.Data;

import java.util.UUID;

@Data
@Entity
@Table(name = "product_image")
public class ProductImage{
    @Id
    @GeneratedValue(strategy = GenerationType.UUID)
    private UUID id;

    private String url;

    @ManyToOne
    @JoinColumn(name= "product_id")
    private ProductEntity product;
}
