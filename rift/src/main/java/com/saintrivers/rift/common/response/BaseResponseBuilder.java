package com.saintrivers.rift.common.response;

import java.time.LocalDateTime;

public class BaseResponseBuilder<T> {
    public static BaseResponse<?> of(String message, Object payload) {
        return new BaseResponse<>(message, LocalDateTime.now(), payload);
    }

    public static BaseResponse<?> of(String message) {
        return new BaseResponse<>(message, LocalDateTime.now(), null);
    }
}