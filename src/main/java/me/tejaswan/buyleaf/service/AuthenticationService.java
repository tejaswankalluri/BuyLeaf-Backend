package me.tejaswan.buyleaf.service;

import me.tejaswan.buyleaf.dto.JwtAuthenticationResponse;
import me.tejaswan.buyleaf.dto.RefreshTokenRequest;
import me.tejaswan.buyleaf.dto.SignUpRequest;
import me.tejaswan.buyleaf.dto.SigninRequest;
import me.tejaswan.buyleaf.entity.UserEntity;

public interface AuthenticationService {
    UserEntity signup(SignUpRequest signUpRequest);
    JwtAuthenticationResponse signin(SigninRequest signinRequest);
    JwtAuthenticationResponse refreshToken(RefreshTokenRequest refreshTokenRequest);

}
