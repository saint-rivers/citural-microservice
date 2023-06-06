package com.saintrivers.commandtower.module.project.web.response;

import java.time.LocalDateTime;
import java.util.List;

import lombok.Builder;

@Builder
public record ProjectResponse(
                long id, String name, List<EndpointResponse> endpoints,
                LocalDateTime createdAt, LocalDateTime updatedAt) {
}
