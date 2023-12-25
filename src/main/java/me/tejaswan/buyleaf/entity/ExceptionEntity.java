package me.tejaswan.buyleaf.entity;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.springframework.http.HttpStatus;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class ExceptionEntity {
    private HttpStatus Status;
    private String Message;
    private boolean Success = false;
}
