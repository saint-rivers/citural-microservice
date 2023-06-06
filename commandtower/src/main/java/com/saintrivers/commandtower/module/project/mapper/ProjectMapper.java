package com.saintrivers.commandtower.module.project.mapper;

import org.springframework.stereotype.Component;

import com.saintrivers.commandtower.module.project.web.request.ProjectRequest;
import com.saintrivers.commandtower.module.project.web.response.ProjectResponse;
import com.saintrivers.commandtower.shared.entity.Project;

import lombok.RequiredArgsConstructor;

@Component
@RequiredArgsConstructor
public class ProjectMapper {

    private final EndpointMapper endpointMapper;

    public Project toEntity(ProjectRequest r) {
        return Project.builder()
                .name(r.name())
                .build();
    }

    public ProjectResponse toResponse(Project d) {
        return ProjectResponse.builder()
                .id(d.getId())
                .name(d.getName())
                .endpoints(endpointMapper.toResponse(d.getEndpoints()))
                .build();
    }
}
