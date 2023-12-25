package me.tejaswan.buyleaf.exception;

import me.tejaswan.buyleaf.entity.ExceptionEntity;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;
import org.springframework.web.context.request.WebRequest;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

@RestControllerAdvice
public class ResponseEntityErrorHandler extends ResponseEntityExceptionHandler {
    @ExceptionHandler(ResourceNotFoundException.class)
    public ResponseEntity<ExceptionEntity> ResourceNotFoundException(ResourceNotFoundException exception, WebRequest request) {
        ExceptionEntity message = new ExceptionEntity(HttpStatus.NOT_FOUND, exception.getMessage(), false);
        return ResponseEntity.status(message.getStatus()).body(message);
    }
}
