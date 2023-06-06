package com.saintrivers.commandtower.module.project.web.request;

import com.saintrivers.commandtower.shared.constant.HttpMethod;

public record EndpointRequest(
        String name, String path, HttpMethod method, Object responseBody, long projectId) {

}
