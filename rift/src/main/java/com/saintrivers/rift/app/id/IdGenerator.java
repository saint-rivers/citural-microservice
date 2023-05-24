package com.saintrivers.rift.app.id;

import java.util.UUID;

import org.springframework.stereotype.Component;

@Component
public class IdGenerator {

    public static String generateAlphaNumericId() {
        UUID uuid = UUID.randomUUID();
        return uuid.toString().replace("-", "");
    }
}
