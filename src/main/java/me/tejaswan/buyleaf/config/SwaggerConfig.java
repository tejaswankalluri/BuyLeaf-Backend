package me.tejaswan.buyleaf.config;

import org.springframework.context.annotation.Configuration;

import io.swagger.v3.oas.annotations.OpenAPIDefinition;
import io.swagger.v3.oas.annotations.enums.SecuritySchemeType;
import io.swagger.v3.oas.annotations.info.Contact;
import io.swagger.v3.oas.annotations.info.Info;
import io.swagger.v3.oas.annotations.security.SecurityScheme;

@Configuration
@OpenAPIDefinition(info = @Info(title = "BuyLeaf Backend", description = "", contact = @Contact(name = "Tejaswan Kalluri", email = "tejaswan@proton.me", url = "https://tejaswan.me"), version = "0.0.1"))
@SecurityScheme(name = "jwtSchema", type = SecuritySchemeType.HTTP, bearerFormat = "JWT", scheme = "bearer")
public class SwaggerConfig {

}
