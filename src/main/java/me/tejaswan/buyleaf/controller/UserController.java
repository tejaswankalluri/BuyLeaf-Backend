package me.tejaswan.buyleaf.controller;

import lombok.RequiredArgsConstructor;
import me.tejaswan.buyleaf.entity.UserEntity;
import me.tejaswan.buyleaf.exception.ResourceNotFoundException;
import me.tejaswan.buyleaf.repository.UserRepository;

import java.util.Optional;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import io.swagger.v3.oas.annotations.security.SecurityRequirement;
import jakarta.servlet.http.HttpServletRequest;

@RestController
@SecurityRequirement(name = "jwtSchema")
@RequestMapping("/api/v1/user")
@RequiredArgsConstructor
public class UserController {
    private final UserRepository userRepository;

    @GetMapping
    public ResponseEntity<UserEntity> sayHello(HttpServletRequest request) throws ResourceNotFoundException {
         String userEmail = (String) request.getAttribute("userEmail");
         Optional<UserEntity> user = userRepository.findByEmail(userEmail);
         if(user.isEmpty()){
             throw new ResourceNotFoundException("User Not Found");
         }
         return ResponseEntity.ok(user.get());
    }
}
