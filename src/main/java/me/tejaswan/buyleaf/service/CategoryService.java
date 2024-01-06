package me.tejaswan.buyleaf.service;

import me.tejaswan.buyleaf.entity.CategoryEntity;
import me.tejaswan.buyleaf.exception.ResourceNotFoundException;

import java.util.List;
import java.util.UUID;

public interface CategoryService {
    List<CategoryEntity> getAllCategories();

    CategoryEntity createCategory(String categoryName);

    void deleteCategory(UUID id) throws ResourceNotFoundException;
}
