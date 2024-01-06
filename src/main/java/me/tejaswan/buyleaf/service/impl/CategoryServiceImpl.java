package me.tejaswan.buyleaf.service.impl;

import lombok.AllArgsConstructor;
import me.tejaswan.buyleaf.entity.CategoryEntity;
import me.tejaswan.buyleaf.exception.ResourceNotFoundException;
import me.tejaswan.buyleaf.repository.CategoryRepository;
import me.tejaswan.buyleaf.service.CategoryService;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.UUID;

@Service
@AllArgsConstructor
public class CategoryServiceImpl implements CategoryService {
    private final CategoryRepository categoryRepository;

    public List<CategoryEntity> getAllCategories() {
        return categoryRepository.findAll();
    }

    public CategoryEntity createCategory(String categoryName) {
        CategoryEntity category = new CategoryEntity();
        category.setName(categoryName);
        return categoryRepository.save(category);
    }

    public void deleteCategory(UUID id) throws ResourceNotFoundException {
        CategoryEntity category = categoryRepository.findById(id).orElseThrow(() -> new ResourceNotFoundException("category Not Found"));
        categoryRepository.delete(category);
    }
}
