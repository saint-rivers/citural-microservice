package com.saintrivers.commandtower.shared.response;

import lombok.Builder;

@Builder
public record PagedResponse<T>(
        T payload, int pageCount) {

}
