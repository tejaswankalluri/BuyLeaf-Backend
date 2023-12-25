package me.tejaswan.buyleaf.exception;

import me.tejaswan.buyleaf.entity.ExceptionEntity;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestControllerAdvice;
import org.springframework.web.context.request.WebRequest;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

@RestControllerAdvice
public class ResponseEntityErrorHandler extends ResponseEntityExceptionHandler {
    @ExceptionHandler(ResourceNotFoundException.class)
    @ResponseStatus(HttpStatus.NOT_FOUND)
    public ResponseEntity<ExceptionEntity> ResourceNotFoundException(ResourceNotFoundException exception, WebRequest request) {
        ExceptionEntity message = new ExceptionEntity(HttpStatus.NOT_FOUND, exception.getMessage(), false);
        return ResponseEntity.status(message.getStatus()).body(message);
    }
    @ExceptionHandler(Exception.class)
    @ResponseStatus(HttpStatus.INTERNAL_SERVER_ERROR)
    public ResponseEntity<ExceptionEntity> handleGlobalException(Exception exception, WebRequest request) {
        ExceptionEntity message = new ExceptionEntity(HttpStatus.INTERNAL_SERVER_ERROR, exception.getMessage(), false);
        return ResponseEntity.status(message.getStatus()).body(message);
    }
}
