package com.saintrivers.commandtower.module.project.web.response;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Map;

import com.saintrivers.commandtower.shared.constant.HttpMethod;

import lombok.Builder;

@Builder
public record EndpointResponse(
        String name,
        String path,
        Map<String, Object> responseBody,
        HttpMethod method,
        List<String> tag,
        int projectId,
        LocalDateTime createdAt,
        LocalDateTime updatedAt,
        String createdBy) {
}
