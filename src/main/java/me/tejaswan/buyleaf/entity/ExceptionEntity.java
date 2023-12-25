package me.tejaswan.buyleaf.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import lombok.RequiredArgsConstructor;
import org.springframework.http.HttpStatus;

import java.time.LocalDateTime;
import java.util.Optional;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class ExceptionEntity {
    private HttpStatus Status;
    private String Message;
    private boolean Success = false;
    private final String Timestamp = LocalDateTime.now().toString();
}
