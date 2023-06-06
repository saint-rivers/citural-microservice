package com.saintrivers.commandtower.module.project.mapper;

import java.util.List;

import org.springframework.stereotype.Component;

import com.saintrivers.commandtower.module.project.web.request.EndpointRequest;
import com.saintrivers.commandtower.module.project.web.response.EndpointResponse;
import com.saintrivers.commandtower.shared.entity.Endpoint;
import com.saintrivers.commandtower.shared.mapper.BodyMapper;

import lombok.RequiredArgsConstructor;

@Component
@RequiredArgsConstructor
public class EndpointMapper {

    public List<EndpointResponse> toResponse(List<Endpoint> d) {
        return d.stream().map(this::toResponse).toList();
    }

    public EndpointResponse toResponse(Endpoint d) {
        return EndpointResponse.builder()
                .method(d.getMethod())
                .name(d.getName())
                .path(d.getPath())
                .responseBody(BodyMapper.parameters(d.getResponseBody()))
                .build();
    }

    public Endpoint toEntity(EndpointRequest r) {
        return Endpoint.builder()
                .method(r.method())
                .name(r.name())
                .path(r.path())
                .responseBody(r.responseBody())
                .build();
    }
}
