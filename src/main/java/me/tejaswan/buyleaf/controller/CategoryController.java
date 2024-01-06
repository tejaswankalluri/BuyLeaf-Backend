package me.tejaswan.buyleaf.controller;

import io.swagger.v3.oas.annotations.security.SecurityRequirement;
import io.swagger.v3.oas.annotations.tags.Tag;
import lombok.RequiredArgsConstructor;
import me.tejaswan.buyleaf.dto.category.CreateCategoryDao;
import me.tejaswan.buyleaf.entity.CategoryEntity;
import me.tejaswan.buyleaf.exception.ResourceNotFoundException;
import me.tejaswan.buyleaf.service.CategoryService;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.UUID;

@RestController
@RequestMapping("/api/v1/category")
@RequiredArgsConstructor
@Tag(name = "Category", description = "Category Apis")
@SecurityRequirement(name = "jwtSchema")
public class CategoryController {
    private final CategoryService categoryService;
    @GetMapping
    @PreAuthorize("hasAnyAuthority('ADMIN', 'USER')")
    public ResponseEntity<List<CategoryEntity>> getAllCategories(){
        return ResponseEntity.ok(categoryService.getAllCategories());
    }
    @PostMapping
    @PreAuthorize("hasAuthority('ADMIN')")
    public ResponseEntity<CategoryEntity> createCategory(@RequestBody CreateCategoryDao categoryDao){
        return ResponseEntity.ok(categoryService.createCategory(categoryDao.getName()));
    }
    @DeleteMapping("/{id}")
    @PreAuthorize("hasAuthority('ADMIN')")
    public ResponseEntity<String> deleteCategory(@PathVariable("id") UUID id) throws ResourceNotFoundException {
        categoryService.deleteCategory(id);
        return ResponseEntity.ok("ok");
    }

}
