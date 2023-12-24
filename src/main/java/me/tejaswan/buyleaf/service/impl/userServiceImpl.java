package me.tejaswan.buyleaf.service.impl;

import lombok.RequiredArgsConstructor;
import me.tejaswan.buyleaf.repository.UserRepository;
import me.tejaswan.buyleaf.service.UserService;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class userServiceImpl implements UserService {
    private final UserRepository userReposityory;

    public UserDetailsService userDetailsService(){
        return new UserDetailsService() {
            @Override
            public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
                return userReposityory.findByEmail(username).orElseThrow(() -> new UsernameNotFoundException("User not found"));
            }
        };
    }
}
