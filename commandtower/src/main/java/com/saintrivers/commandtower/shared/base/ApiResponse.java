package com.saintrivers.commandtower.shared.base;

import lombok.Builder;

@Builder
public record ApiResponse<T>(
        T payload) {
}
