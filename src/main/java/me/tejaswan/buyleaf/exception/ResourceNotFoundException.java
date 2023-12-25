package me.tejaswan.buyleaf.exception;

public class ResourceNotFoundException extends Exception{
    public ResourceNotFoundException(){
        super();
    }
    public ResourceNotFoundException(String message){
        super(message);
    }
    public ResourceNotFoundException(Throwable cause){
        super(cause);
    }
    public ResourceNotFoundException(String message, Throwable cause){
        super(message, cause);
    }
}
